<font face="JetBrains Mono">1. 安装好go的编译环境，设置好GOROOT和GOPATH，假设GOROOT为~/go</font>
<br/>
<font face="JetBrains Mono">2. 添加~/go/src到GOPATH</font>
<br/>
<font face="JetBrains Mono">3. 创建目录~/go/src/github.com/elastic</font>
<br/>
<font face="JetBrains Mono">4. 下载beats.tar.gz到go的源码包~/go/src/github.com/elastic，解压</font>
<br/>
<font face="JetBrains Mono">5. 关于入口mysql的代码在~/go/src/github.com/elastic/beats/libbeat/outputs/mysql</font>

<br/>
<font face="JetBrains Mono">6. 进入filebeat的目录编译：</font>

<br/>
<font face="JetBrains Mono">6.1 cd ~/go/src/github.com/elastic/beats/filebeat</font>

<br/>
<font face="JetBrains Mono">6.2 make </font>

<br/>
<font face="JetBrains Mono">在~/go/src/github.com/elastic/beats/filebeat文件夹生成可执行文件filebeat</font>



