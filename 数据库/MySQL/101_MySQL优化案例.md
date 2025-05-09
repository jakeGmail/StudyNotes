[toc]

# 1 慢查询优化

**背景**：在进行分页查询的时候，如果数据量很大，当分页到很后面的时候，mysql其实也是从第一条记录开始查询的。
就会导致很后面的分页的查询变慢，成为一个慢SQL. 

**解决方案**: 在每一次分页查询后，让下一页的分页查询返回上一次查询的最大的主键id. 那么在执行分页查询的时候，就添加条件只查询大于这个id的记录


# 2 数据量太大导致查询慢

**背景**：一个数据库实例中中有十几亿条数据，每次查询比较慢。但不想对数据库进行大的改动（例如分库、分表）

**解决方案**： 因为mysql是按页存储的，每一页16kb, 当我们查询到一条数据的时候，MYSQL会把这条数据所在的页加载到 buffer pool中。
下次查找时，会先在buffer pool中查询数据，如果在buffer pool中美元找到才会到磁盘中去查找。
因此把buffer pool的大小改大一些（几个G）