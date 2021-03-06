## 6 查询性能优化
> 查询优化、索引优化、库表结构优化

### 6.1 为什么查询速度会变慢
> 一些不必要的额外操作，某些操作被额外地重复了很多次，某些操作执行的太慢。

### 6.2 慢查询基础：优化数据访问
> 查询性能低下最基本的原因是访问的数据太多。
> 低效查询的分析步骤：
> 1. 确认应用程序是否在检索大量超过需要的数据。
> 2. 确认MySQL服务器层是否在分析大量超过需要的数据行。

#### 6.2.1 是否向数据库请求了不需要的数据
> 避免请求实际需要的数据，多余的数据会给MySQL服务器带来额外的负担，并增加网络的开销，也会消耗应用服务器的CPU和内存资源。

#### 6.2.2 MySQL是否在扫描额外的记录
> 最简单的衡量查询的开销的三个指标如下：
> * 相应时间--服务时间（数据库处理查询的时间）和排队时间（服务器等待某些资源的时间）
> * 扫描的行数
> * 返回的行数

> 一般MySQL能够使用如下三种方式应用WHERE条件，从好到坏依次为：
> * 索引中使用WHERE条件来过滤不匹配的记录。存储引擎层完成的。
> * 使用索引覆盖扫描（Extra中为Using index）来返回记录，直接从索引中过滤不需要的记录并返回命中的结果。MySQL服务器层完成的，无须在回表查询记录。
> * 从数据表中返回数据，然后过滤不满足条件的记录（Extra中为Using Where）。MySQL服务器层完成，MySQL需要先从数据表读出记录然后过滤。

> 查询需要扫描大量的数据但返回少量的行的优化技巧
> * 使用索引覆盖扫描，把所有需要用的列都放到索引中。
> * 
