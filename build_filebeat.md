# 1. 安装好go的编译环境，设置好GOROOT和GOPATH，假设GOPATH为~/go
# 2. 添加~/go/src到GOPATH；
# 3. 创建目录~/go/src/github.c/elastic/ 
# 4. 下载beats.tar.gz到go的源码包~/go/src/github.c/elastic，解压
# 5. 进入filebeat的目录编译：
## 5.1 cd ~/go/src/github.com/elastic/beats/filebeat
## 5.2 make 
## 5.3 在~/go/src/github.com/elastic/beats/filebeat文件夹生成可执行文件filebeat

# 6. 关于入口mysql的代码在~/go/src/github.com/elastic/beats/libbeat/outputs/mysql
