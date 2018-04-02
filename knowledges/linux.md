### 对服务器各种状态下的连接数分组并查询得到结果：
`netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}' `
`netstat -na|wc -l  `
`netstat -nat|grep ESTABLISHED|wc -l  `
