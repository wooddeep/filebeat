<font color=black size=12 face="JitBrains Mono">黑体</font>
<font color=black size=12 face="JitBrains Mono"> 1. 安装好go的编译环境，设置好GOROOT和GOPATH，假设GOROOT为~/go</font>
<font color=black size=12 face="JitBrains Mono"> 2. 添加~/go/src到GOPATH；</font>
<font color=black size=12 face="JitBrains Mono"> 3. 创建目录~/go/src/github.com/elastic/ </font>
<font color=black size=12 face="JitBrains Mono"> 4. 下载beats.tar.gz到go的源码包~/go/src/github.com/elastic，解压</font>
<font color=black size=12 face="JitBrains Mono"> 4.1 拷贝最根目录下的的mysql.go, config.go 到~/go/src/github.com/elastic/beats/libbeat/outputs/mysql文件夹下面</font>
<font color=black size=12 face="JitBrains Mono"> 5. 关于入口mysql的代码在~/go/src/github.com/elastic/beats/libbeat/outputs/mysql</font>

<font color=black size=12 face="JitBrains Mono"> 6. 进入filebeat的目录编译：</font>
<font color=black size=12 face="JitBrains Mono"> 6.1 cd ~/go/src/github.com/elastic/beats/filebeat</font>
<font color=black size=12 face="JitBrains Mono"> 6.2 make </font>
<font color=black size=12 face="JitBrains Mono"> 6.3 在~/go/src/github.com/elastic/beats/filebeat文件夹生成可执行文件filebeat</font>

## 关于mysql的配置


