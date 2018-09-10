### 对服务器各种状态下的连接数分组并查询得到结果：
`netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}' `
`netstat -na|wc -l  `
`netstat -nat|grep ESTABLISHED|wc -l  `

```
状态：描述 
CLOSED：无连接是活动的或正在进行 
LISTEN：服务器在等待进入呼叫 
SYN_RECV：一个连接请求已经到达，等待确认 
SYN_SENT：应用已经开始，打开一个连接 
ESTABLISHED：正常数据传输状态 
FIN_WAIT1：应用说它已经完成 
FIN_WAIT2：另一边已同意释放 
ITMED_WAIT：等待所有分组死掉 
CLOSING：两边同时尝试关闭 
TIME_WAIT：另一边已初始化一个释放 
LAST_ACK：等待所有分组死掉
```

1、netstat -tpln
然后观察两个tomcat的监听是怎么写的

2、iptables -vnL
去查看到防火墙是否有信任lo

3、getenforce
检查selinux配置

##### ab 
-n 测试会话中所执行的请求个数,默认仅执行一个请求
-c 一次产生的请求个数,即同一时间发出多少个请求,默认为一次一个
-t 测试所进行的最大秒数,默认为无时间限制....其内部隐含值是[-n 50000],它可以使对服务器的测试限制在一个固定的总时间以内
-p 包含了需要POST的数据的文件
-T POST数据所使用的Content-type头信息
-v 设置显示信息的详细程度
-w 以HTML表格的形式输出结果,默认是白色背景的两列宽度的一张表
-i 以HTML表格的形式输出结果,默认是白色背景的两列宽度的一张表
-x 设置<table>属性的字符串,此属性被填入<table 这里>
-y 设置<tr>属性的字符串
-z 设置<td>属性的字符串
-C 对请求附加一个Cookie行，其典型形式是name=value的参数对,此参数可以重复
-H 对请求附加额外的头信息,此参数的典型形式是一个有效的头信息行,其中包含了以冒号分隔的字段和值的对(如"Accept-Encoding: zip/zop;8bit")
-A HTTP验证,用冒号:分隔传递用户名及密码
-P 无论服务器是否需要(即是否发送了401认证需求代码),此字符串都会被发送
-X 对请求使用代理服务器
-V 显示版本号并退出
-k 启用HTTP KeepAlive功能,即在一个HTTP会话中执行多个请求,默认为不启用KeepAlive功能
-d 不显示"percentage served within XX [ms] table"的消息(为以前的版本提供支持)
-S 不显示中值和标准背离值,且均值和中值为标准背离值的1到2倍时,也不显示警告或出错信息,默认会显示最小值/均值/最大值等(为以前的版本提供支持)
-g 把所有测试结果写入一个'gnuplot'或者TSV(以Tab分隔的)文件
-e 产生一个以逗号分隔的(CSV)文件,其中包含了处理每个相应百分比的请求所需要(从1%到100%)的相应百分比的(以微妙为单位)时间
-h 显示使用方法
-k 发送keep-alive指令到服务器端

##### ab -n 1000 -c 1000 -T http://github.com

##### ln -s /usr/bin/ffmpeg /usr/local/bin/ffmpeg

##### sudo lsof -p [进程号] | wc -l