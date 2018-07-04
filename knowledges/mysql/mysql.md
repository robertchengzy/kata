##### SQL来查找一个数据库的重复索引和冗余索引
`SELECT
	a.TABLE_SCHEMA,
	a.TABLE_NAME,
	a.COLUMN_NAME,
	a.INDEX_NAME AS 'index1',
	b.INDEX_NAME AS 'index2' 
FROM
	information_schema.STATISTICS a
	JOIN information_schema.STATISTICS b ON a.TABLE_SCHEMA = b.TABLE_SCHEMA 
	AND a.TABLE_NAME = b.TABLE_NAME 
	AND a.SEQ_IN_INDEX = b.SEQ_IN_INDEX 
	AND a.COLUMN_NAME = b.COLUMN_NAME 
WHERE
	a.SEQ_IN_INDEX = 1 
	AND a.INDEX_NAME <> b.INDEX_NAME;`
	
##### 死锁日志	
`SHOW ENGINE INNODB STATUS;`

##### 查询 正在执行的事务：
`SELECT * FROM information_schema.INNODB_TRX;`

##### 查看正在锁的事务
`SELECT * FROM INFORMATION_SCHEMA.INNODB_LOCKS;`

##### 查看等待锁的事务
`SELECT * FROM INFORMATION_SCHEMA.INNODB_LOCK_WAITS;`

##### 查看表的列数
`SELECT
 	COUNT(*) 
 FROM
 	information_schema.COLUMNS 
 WHERE
 	TABLE_SCHEMA = 'jdkopen' 
 	AND table_name = 'jdk_course'`	
 	
##### 查看表的索引：
`SHOW INDEX FROM table_name（表名）`

`
FLUSH TABLE table_name;
FLUSH STATUS;
SELECT sql;
SHOW STATUS LIKE 'hander_read%';
`

##### 刷新表缓存并清除状态计数器
`FLUSH TABLE t1;
 FLUSH STATUS;
 SELECT COUNT(*) FROM t1 WHERE i1 = 3 AND d = '2000-01-01';
 SHOW STATUS LIKE 'handler_read%'`
 
##### 查看已打开的表数
 `SHOW GLOBAL STATUS LIKE 'Opened_tables';`
 
##### 查看表的碎片
`SELECT
 	table_name,
 	ENGINE,
 	table_rows,
 	( data_length + index_length ) length,
 	DATA_FREE 
 FROM
 	information_schema.TABLES 
 WHERE
 	TABLE_SCHEMA = 'jdkopen' 
 	AND DATA_FREE > 0;`
  	
##### InnoDB清理碎片
`ALTER TABLE jdk_course_calendar ENGINE = INNODB;`

##### 如何查看数据库中的冗余索引
`select * from sys.schema_redundant_indexes;`

##### 如何获取未使用的索引
`select * from sys.schema_unused_indexes;`

##### 查看表生成的DDL
`show create table table_name;`

##### 查看字符集
`SHOW VARIABLES WHERE variable_name LIKE 'character\_set\_%' OR variable_name LIKE 'collation%';`

##### 查看表信息
`show table status like 'table_name'`

##### 查询事务和锁等待的关系
`SELECT
 	R.TRX_ID WAITING_TRX_ID,
 	R.TRX_MYSQL_THREAD_ID WAITING_THREAD,
 	R.TRX_QUERY WATING_QUERY,
 	B.TRX_ID BLOCKING_TRX_ID,
 	B.TRX_MYSQL_THREAD_ID BLOCKING_THREAD,
 	B.TRX_QUERY BLOCKING_QUERY 
 FROM
 	INFORMATION_SCHEMA.INNODB_LOCK_WAITS W
 	INNER JOIN INFORMATION_SCHEMA.INNODB_TRX B ON B.TRX_ID = W.BLOCKING_TRX_ID
 	INNER JOIN INFORMATION_SCHEMA.INNODB_TRX R ON R.TRX_ID = W.REQUESTING_TRX_ID;`
 	
##### 固定字段state 状态 created_at 创建时间 updated_at 更新时间

##### 查询库中表的备注信息
`SELECT TABLE_NAME,table_comment FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'jdkopen';`

我们知道InnoDB一个page的默认大小是16k。由于是Btree组织，要求叶子节点上一个page至少要包含两条记录（否则就退化链表了）。

所以一个记录最多不能超过8k。

又由于InnoDB的聚簇索引结构，一个二级索引要包含主键索引，因此每个单个索引不能超过4k （极端情况，primay-key和某个二级索引都达到这个限制）。

由于需要预留和辅助空间，扣掉后不能超过3500，取个“整数”就是 (1024bytes*3=3072bytes)。
