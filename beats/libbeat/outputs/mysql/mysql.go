/*
 * @Description: 
 * @Author: lihan
 * @Date: 2021-05-20 10:46:19
 */
 package mysql

 import (
	 "database/sql"
	 "fmt"
	 "os"
	 "path/filepath"
	 "time"
	 "strings"
	 "strconv"
	 //"errors"
	 "reflect"
	 _ "github.com/go-sql-driver/mysql"
	 //_ "github.com/lib/pq"
	 "github.com/bitly/go-simplejson" // for json get
 
	 "github.com/elastic/beats/libbeat/beat"
	 "github.com/elastic/beats/libbeat/common"
	 "github.com/elastic/beats/libbeat/common/file"
	 "github.com/elastic/beats/libbeat/logp"
	 "github.com/elastic/beats/libbeat/outputs"
	 "github.com/elastic/beats/libbeat/outputs/codec"
	 "github.com/elastic/beats/libbeat/publisher"
	 "github.com/robertkrimen/otto"
 )
 
 func init() {
	 outputs.RegisterType("mysql", makeMysqlout)
 }
 
 type mysqlOutput struct {
	 beat     beat.Info
	 observer outputs.Observer
	 db       *sql.DB
	 rotator  *file.Rotator
	 codec    codec.Codec
	 cfg      config
	 insert   string
	 parser   string
	 vm       *otto.Otto //otto.New()	
 }
 
 // makeFileout instantiates a new file output instance.
 func makeMysqlout(beat beat.Info, observer outputs.Observer, cfg *common.Config) (outputs.Group, error) {
	 config := defaultConfig
	 if err := cfg.Unpack(&config); err != nil {
		 return outputs.Fail(err)
	 }
 
	 // disable bulk support in publisher pipeline
	 cfg.SetInt("bulk_max_size", -1, -1)
 
	 mo := &mysqlOutput{
		 beat:     beat,
		 observer: observer,
	 }
 
	 if err := mo.init(beat, config); err != nil {
		 return outputs.Fail(err)
	 }
 
	 return outputs.Success(-1, 0, mo)
 }
 
 func (out *mysqlOutput) init(beat beat.Info, c config) error {
	 var path string
	 if c.Filename != "" {
		 path = filepath.Join(c.Path, c.Filename)
	 } else {
		 path = filepath.Join(c.Path, out.beat.Beat)
	 }
	 out.cfg = c
	 dsn := fmt.Sprintf("%s:%s@%s(%s)/%s", c.Username, c.Password, "tcp", c.Address, c.Database)
	 db, err := sql.Open("mysql", dsn)
	 if err != nil {
		 fmt.Printf("Open mysql failed,err:%v\n", err)
		 return err
	 }
	 db.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	 db.SetMaxOpenConns(100)                  // 设置最大连接数
	 db.SetMaxIdleConns(16)                   //设置闲置连接数
	 out.db = db                              // 创建mysql连接
	 out.insert = c.Insert
	 out.parser = c.Parser
	 out.vm = otto.New()

	 out.rotator, err = file.NewFileRotator(
		 path,
		 file.MaxSizeBytes(c.RotateEveryKb*1024),
		 file.MaxBackups(c.NumberOfFiles),
		 file.Permissions(os.FileMode(c.Permissions)),
	 )
	 if err != nil {
		 return err
	 }
 
	 out.codec, err = codec.CreateEncoder(beat, c.Codec)
	 if err != nil {
		 return err
	 }
 
	 logp.Info("Initialized file output. "+
		 "path=%v max_size_bytes=%v max_backups=%v permissions=%v",
		 path, c.RotateEveryKb*1024, c.NumberOfFiles, os.FileMode(c.Permissions))
 
	 return nil
 }
 
 // Implement Outputer
func (out *mysqlOutput) Close() error {
	return out.rotator.Close()
 }


func reflectUse(t interface{}) {
    switch reflect.TypeOf(t).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(t)
        for i := 0; i < s.Len(); i++ {
            fmt.Println(s.Index(i))
        }
    }
}

// func IsNum(s string) bool { _, err := strconv.ParseFloat(s, 64) return err == nil }

func (out *mysqlOutput) extractData(log string) []interface{} {
	js, _ := simplejson.NewJson([]byte(log))
	msgstr := js.Get("message").MustString() //mlogjson,_ := js.Get("message").MarshalJSON()
	out.vm.Set("$", string(msgstr))
	out.vm.Run(out.parser)

	value, _ := out.vm.Run("$$.join()")

	if str, err := value.ToString(); err == nil {
		fmt.Println("## ok")
		raw := strings.Split(str, ",")
		slice := make([] interface {}, len(raw), 2 * len(raw))
		for i, v := range raw {
			if strings.EqualFold("Null", v) {
				slice[i] = nil
			} else if intv, err := strconv.ParseInt(v, 10, 64); err == nil {
				slice[i] = intv
			} else if flotv, err := strconv.ParseFloat(v, 64); err == nil {
				slice[i] = flotv
			} else {
				slice[i] = v
			}
		}
		return slice
	}

	return nil
}

func (out *mysqlOutput) Publish(batch publisher.Batch) error {
	 defer batch.ACK()
 
	 st := out.observer
	 events := batch.Events()
	 st.NewBatch(len(events))
 
	 logp.Err("Initialized file output. #### %s", out.cfg.Address)
 
	 dropped := 0
	 writeBytesLen := 0;
	 db := out.db
	 tx, err := db.Begin()
	 if err != nil {
		 logp.Err("db error: %v", err)
	 }
 
	 sqlStr := out.insert
	 stmt, err := tx.Prepare(sqlStr)
	 if err != nil {
		 logp.Err("Prepare error:", err)
		 panic(err)
	 }
	 
	 for i := range events {
		 event := &events[i]
		 serializedEvent, err := out.codec.Encode(out.beat.Beat, &event.Content)
		 if err != nil {
			 if event.Guaranteed() {
				 logp.Critical("Failed to serialize the event: %v   ", err)
			 } else {
				 logp.Warn("Failed to serialize the event: %v--", err)
			 }
 
			 dropped++
			 continue
		 }
		
		 // 写入到mysql
		 data := out.extractData(string(serializedEvent))
		 if data == nil {
			 dropped++
			 continue
		 }

		//fmt.Printf("%v-%v-%v", data[0], data[1], data[2], data[3])
		 _, err = stmt.Exec(data...)
		 if err != nil {
			fmt.Println("Exec error:", err)
			//panic(err)
			continue
		 }
		 
		 writeBytesLen += len(serializedEvent) + 1
	 }
 
	 err = stmt.Close()
	 if err != nil {
		fmt.Println("stmt close error:", err)
		panic(err)
	 }
 
	 err = tx.Commit()
	 if err != nil {
		fmt.Println("commit error:", err)
		panic(err)
	 }
	 
	 st.WriteBytes(writeBytesLen)
	 st.Dropped(dropped)
	 st.Acked(len(events) - dropped)
 
	 return nil
 }
 