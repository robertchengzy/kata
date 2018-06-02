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

##### •如何查看数据库中的冗余索引
`select * from sys.schema_redundant_indexes;`

##### •如何获取未使用的索引
`select * from sys.schema_unused_indexes;`

##### 查看表生成的DDL
`show create table table_name;`
