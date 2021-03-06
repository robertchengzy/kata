### golang docker k8s

1. new 和 make 的区别

2. map/sync.map的实现原理，为什么是随机遍历 - https://blog.yiz96.com/golang-map/ https://blog.yiz96.com/golang-sync-map/

3. 内存管理和垃圾回收(CMS 三色标记法) https://blog.yiz96.com/golang-mm-gc/ https://www.cnblogs.com/ldaniel/p/8502867.html?utm_source=debugrun&utm_medium=referral

4. goroutine MPG 模型（抢占式调度） https://github.com/developer-learning/night-reading-go/blob/master/reading/20180802/README.md?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io https://github.com/developer-learning/night-reading-go/blob/master/reading/20180802/README.md?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io

5. channel CSP 模型 

6. interface 方法的集合和一种类型

7. slice 和 array  https://www.jianshu.com/p/030aba2bff41

8. reflect 反射

9. 基础数据类型和零值

10. 闭包

11. grpc 

12. protobuf

13. 单元测试

14. 调优

15. select

> https://tiancaiamao.gitbooks.io/go-internals/content/zh/02.3.html

### redis

1. 基础数据类型 string hash list set sorted

2. 持久化：AOF日志文件和RDB内存快照

3. 缓存策略

4. 淘汰策略（LRU算法）

5. 分布式锁

6. 同步机制

7. redis 集群

8. 性能优化

9. 内存碎片

10. 内存模型

11. 数据存储结构、rehash

12. 事务

### mysql

1. 慢查询优化

2. 共享锁和排他锁

3. 索引(B+Tree)

4. 分库分表分区

5. 主从复制原理

6. 读写分离

7. 事务的隔离级别

8. innodb和myisam

9. 表结构设计

10. 表碎片

### 数据结构和算法

1. 数组，字符串，链表，树，排序和搜索，动态规划，设计问题。

2. 时间复杂度和空间复杂度

### 计算机网络 https://www.cnblogs.com/Java3y/p/8444033.html

1. TCP和UDP的区别

2. 如何用UDP实现TCP

3. http1.0 http1.1 http2 https

4. 一次完整的HTTP请求（当你输入一个网址，实际会发生什么?）

5. 三次握手 四次挥手

### 工作经历

1. 支付功能（抢购问题实现，订单过期时间）

2. 短信/消息功能（消息集中发送问题）

3. 微信公众号，小程序等api（api调用）

4. 打卡系统（并发问题，重复打卡和缓存的使用）

5. redis共享session

6. 分析问题解决问题的能力

7. 学习能力

### 其他

1. 设计模式(单例模式)

2. hash一致性

3. 微服务

4. 分布式（事务，数据一致性）

5. 简历（star法则）

6. 缓冲区溢出

7. unix

8. ngnix

### 操作系统：

1. 线程和进程的区别

2. 并发和并行

3. 线程的生命周期

4. 什么是死锁？

### 计算机网络:

输入www.google.com 会发生什么？（confluent／houzz／yahoo）

1. TCP 三次握手，TCP/UDP 区别；

2. http/https 区别；http request：post／get ；http port 80 ssl;

3. Public key/Private key;

4. HTTP 401, 403, or 404 Error

5. Client/Server Model


### 答案

##### new和make的区别

> new：内置函数 new 分配空间。传递给 new 函数的是一个类型，不是一个值。返回值是指向这个新分配的零值的指针。

> make：内置函数 make 分配并且初始化 一个 slice, 或者 map 或者 chan 对象。 并且只能是这三种对象。 和 new 一样，第一个参数是 类型，
不是一个值。 但是make 的返回值就是这个类型（即使一个引用类型），而不是指针。 具体的返回值，依赖具体传入的类型。


##### golang 反射reflection
反射可大大提高程序的灵活性，使得interface{}有更大的发挥余地
反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
反射会将匿名字段作为独立字段(匿名字段本质)
想要利用反射修改对象状态，前提是interface.data是settable，即pointer-interface
通过反射可以“动态”调用方法

##### goroutine
一个Goroutine会以一个很小的栈启动（可能是2KB或4KB），当遇到深度递归导致当前栈空间不足时，Goroutine会根据需要动态地伸缩栈的大小
（主流实现中栈的最大值可达到1GB）。因为启动的代价很小，所以我们可以轻易地启动成千上万个Goroutine。
Go的运行时还包含了其自己的调度器，这个调度器使用了一些技术手段，可以在n
个操作系统线程上多工调度m个Goroutine。Go调度器的工作和内核的调度是相似
的，但是这个调度器只关注单独的Go程序中的Goroutine。Goroutine采用的是半抢
占式的协作调度，只有在当前Goroutine发生阻塞时才会导致调度；同时发生在用户
态，调度器会根据具体函数只保存必要的寄存器，切换的代价要比系统线程低得
多。运行时有一个 runtime.GOMAXPROCS 变量，用于控制当前运行正常非阻塞
Goroutine的系统线程数目。

##### 垃圾回收
Golang现在使用的垃圾回收算法是CMS，并行的mark-and-sweep，使用的是四色标注法（也有叫三色的）：

新建的内存节点标记为白色（为了防止误删，引入了“当前白色”和“其他白色”两个状态，所以为啥叫四色）；
每次从根节点进行扫描，遇到白色节点就将其变成灰色，放入灰色链表（初始化）；
遍历灰色链表，本元素标记为黑色；相邻元素如果是白色，则标记为灰色，放入灰色链表（mark）；
扫描所有节点，删除白色元素（sweep）

hashmap 通过一个 bucket 数组实现，所有元素将被 hash 到数组中的 bucket 中，bucket 填满后，将通过一个 overflow 指针来扩展一个 bucket 出来形成链表，
也就是解决冲突问题。这也就是一个基本的 hash 表结构，没什么新奇的东西，下面总结一些细节吧。

注意一个 bucket 并不是只能存储一个 key/value 对，而是可以存储8个 key/value 对。每个 bucket 由 header 和 data 两部分组成，
data 部分的内存大小为：(sizeof(key) + sizeof(value)) * 8，也就是要存储8对 key/value，这8对 key/value 在 data 内存中的存储顺序是：key0key1…key7value0value1…value7，
是按照顺序先依次存储8个 key 值，然后存储对应的8个 value。 为什么不是存储为 key0value0…key7value7 呢？主要是方便访问吧。
如果 key, value 的类型大小超过了128字节，将不会直接存储值，而是存储其指针。
bucket 的 header 部分有一个 uint8 tophash[8] 数组，这个数组将用来存储8个 key 的 hash 值的高8位值。比如：tophash[0] 存储的值就
是 hash(key0) » (64 - 8)。保存了一个 key 的 hash 高8位部分，在查找/删除/插入一个 key 的时候，可以先判断两个 key hash 的高8位是否相等，
如果不等，那就根本不用去比较 key 的内容。所以这里保存一下 hash 值的高8位可以作为第一步的粗略过滤，不少时候可以省掉比较两个 key 的内容，
因为比较两个 key 是否相等的代价远比两个 uint8 的代价高。当然，这里如果存储整个 hash 值，而不仅仅是高8位的话，判断效果将更好，但内存的占用就会多很多了。
bucket 的8个 key/value 空间如果都填满后，就会分配新的 bucket，通过 overflow 指针串联起来。注意这个链表指针被命名为 overflow，代表的正是 bucket 溢出了，
这个命名感觉很好，hash 表实现的时候我们应该努力避免 bucket overflow。
hashmap 是会自增长的，也就说随着插入的 kv 对越来越多，初始的 bucket 数组就可以需要增长、重新hash 所有元素，性能才会好。bucket
 数组增长的时机就是插入的元素个数大于了 bucket数组大小 * 6.5，为什么是6.5，这个在代码注释里有说明，主要是测试出来的经验值。
hashmap 每次增长，都是重新分配一个新的 bucket 数组，新 bucket 数组是之前 bucket 数组的2倍大小。 
7.hashmap 增长后，需要将老 bucket 数组中的元素拷贝到新的 bucket 数组，这个拷贝过程不是一口气立马完成的，而是采用了增量式的拷贝，
也就是说分配了新的 bucket 数组后，并没有立刻拷贝元素，而是等接下来每次插入一个元素的时候，才拷贝一点，随着插入的动作增多，
逐渐就将全部元素拷贝到了新的 bucket 数组中。
8.在 make 一个 map 对象的时候，如果不指定大小的话，bucket 数组默认就是1了，随着插入的元素增多，就会增长成2，4，8，16等。
可以看出不指定初始化大小的map，很可能要经历很多次的增长、元素拷贝。我们应该给 map 指定一个合适的大小值。

##### slice
并非所有时候都适合用切片代替数组，因为切片底层数组可能会在堆上分配内存，而且小数组在栈上拷贝的消耗也未必比 make 消耗大。
切片本身并不是动态数组或者数组指针。它内部实现的数据结构通过指针引用底层数组，设定相关属性将数据读写操作限定在指定的区域内。
切片本身是一个只读对象，其工作机制类似数组指针的一种封装。
给定项的切片索引可能比相关数组的相同元素的索引小。和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：切片是一个长度可变的数组。
Go 中切片扩容的策略是这样的：

如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。

一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一。

注意：扩容扩大的容量都是针对原来的容量而言的，而不是针对原来数组的长度而言的。

##### Redis如何做持久化的？

bgsave做镜像全量持久化，aof做增量持久化。因为bgsave会耗费较长时间，不够实时，在停机的时候会导致大量丢失数据，所以需要aof来配合使用。在redis实例重启时，
会使用bgsave持久化文件重新构建内存，再使用aof重放近期的操作指令来实现完整恢复重启之前的状态。

对方追问那如果突然机器掉电会怎样？取决于aof日志sync属性的配置，如果不要求性能，在每条写指令时都sync一下磁盘，就不会丢失数据。
但是在高性能的要求下每次都sync是不现实的，一般都使用定时sync，比如1s1次，这个时候最多就会丢失1s的数据。

对方追问bgsave的原理是什么？你给出两个词汇就可以了，fork和cow。fork是指redis通过创建子进程来进行bgsave操作，cow指的是copy on write，
子进程创建后，父子进程共享数据段，父进程继续提供读写服务，写脏的页面数据会逐渐和子进程分离开来。

##### Redis的五种数据类型
字符串string：字符串类型是Redis中最为基础的数据存储类型，是一个由字节组成的序列，他在Redis中是二进制安全的，这便意味着该类型可以接受任何格式的数据，
如JPEG图像数据货Json对象描述信息等，是标准的key-value，一般来存字符串，整数和浮点数。Value最多可以容纳的数据长度为512MB
应用场景：很常见的场景用于统计网站访问数量，当前在线人数等。incr命令(++操作)

列表list：Redis的列表允许用户从序列的两端推入或者弹出元素，列表由多个字符串值组成的有序可重复的序列，是链表结构。好比Java的linkedList，
在往两端插入和删除数据时，效率是非常高的，往中间插入数据效率是很低下的。List中可以包含的最大元素数量是4294967295。
应用场景：1.最新消息排行榜。2.消息队列，以完成多程序之间的消息交换。可以用push操作将任务存在list中（生产者），然后线程在用pop操作将任务取出进行执行。（消费者）

集合set：Redis的集合是无序不可重复的，和列表一样，在执行插入和删除和判断是否存在某元素时，效率是很高的。集合最大的优势在于可以进行交集并集差集操作。
Set可包含的最大元素数量是4294967295。
应用场景：1.利用交集求共同好友。2.利用唯一性，可以统计访问网站的所有独立IP。3.好友推荐的时候根据tag求交集，大于某个threshold（临界值的）就可以推荐。

散列hash：Redis中的散列可以看成具有String key和String value的map容器，可以将多个key-value存储到一个key中。每一个Hash可以存储4294967295个键值对。
应用场景：例如存储、读取、修改用户属性（name，age，pwd等）

有序集合zset：和set很像，都是字符串的集合，都不允许重复的成员出现在一个set中。他们之间差别在于有序集合中每一个成员都会有一个分数(score)与之关联，
Redis正是通过分数来为集合中的成员进行从小到大的排序。尽管有序集合中的成员必须是卫衣的，但是分数(score)却可以重复。
应用场景：可以用于一个大型在线游戏的积分排行榜，每当玩家的分数发生变化时，可以执行zadd更新玩家分数(score)，此后在通过zrange获取几分top ten的用户信息。

##### Redis的同步机制了解么？

Redis可以使用主从同步，从从同步。第一次同步时，主节点做一次bgsave，并同时将后续修改操作记录到内存buffer，待完成后将rdb文件全量同步到复制节点，
复制节点接受完成后将rdb镜像加载到内存。加载完成后，再通知主节点将期间修改的操作记录同步到复制节点进行重放就完成了同步过程。

##### http请求
1. 域名解析
2. 发起TCP三次握手
3. 发起HTTP请求
4. 服务器响应HTTP请求
5. 浏览器解析HTML代码，并请求HTML代码中的资源
6. 浏览器对页面进行渲染给用户

##### http1.1
a、默认持久连接节省通信量，只要客户端服务端任意一端没有明确提出断开TCP连接，就一直保持连接，可以发送多次HTTP请求
b、管线化，客户端可以同时发出多个HTTP请求，而不用一个个等待响应
c、断点续传原理

##### http2.0
a. 多路复用共享连接
b. 二进制分帧层
c. 请求优先级
d. 服务端推送
e. 首部压缩

