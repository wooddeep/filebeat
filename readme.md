<font color=black size=12 face="JitBrains Mono">黑体</font>
## 1. 安装好go的编译环境，设置好GOROOT和GOPATH，假设GOROOT为~/go
## 2. 添加~/go/src到GOPATH；
## 3. 创建目录~/go/src/github.com/elastic/ 
## 4. 下载beats.tar.gz到go的源码包~/go/src/github.com/elastic，解压
### 4.1 拷贝最根目录下的的mysql.go, config.go 到~/go/src/github.com/elastic/beats/libbeat/outputs/mysql文件夹下面
## 5. 关于入口mysql的代码在~/go/src/github.com/elastic/beats/libbeat/outputs/mysql

## 6. 进入filebeat的目录编译：
### 6.1 cd ~/go/src/github.com/elastic/beats/filebeat
### 6.2 make 
### 6.3 在~/go/src/github.com/elastic/beats/filebeat文件夹生成可执行文件filebeat

## 关于mysql的配置


