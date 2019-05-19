# 前言

实际工作中会使用mybatis作为业务的 持久层框架，mybatis封装好了对JDBC的调用。便于开发者使用。而要深入理解mybatis，就需要对JDBC有所了解。

![JDBC](http://upload-images.jianshu.io/upload_images/3991621-0cfdbbdb02918e74.gif?imageMogr2/auto-orient/strip)

## 获取数据库连接

### 方式一 通过 DriverManager
```
val url = "jdbc:mysql://10.252.28.152:3306/test?serverTimezone=UTC&useUnicode=true&characterEncoding=utf8&useSSL=false"
val conn = DriverManager.getConnection(url, "gen", "123")
```

* 1. DriverManager 初始化时，首先读取 系统属性 jdbc.drivers ，用于加载驱动类。如果 驱动已SPI形式打包，会自动加载。
![mysql 驱动](https://upload-images.jianshu.io/upload_images/3991621-55fca24cf56cbb34.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

* 2. DriverManager.getConnection 通过url前缀与已注册的驱动匹配，调用driver.connect() 获取Connection。

### 方式二 通过 DataSource

作为DriverManager工具的替代，数据源对象是获取连接的首选方法。

通常用于实现连接池 及 分布式 事务。


## 执行sql

```
val stat = conn.getStatement()
// insert , update, delete
stat.executeUpdate(sqlcmd)
// query
val rs : ResultSet = stat.executeQuery(sqlcmd)

stat.close()
```

* 1. 创建Statement对象，dml语句调用executeUpdate 查询调用 executeQuery


预备语句PreparedStatement，用？表示参数，每次执行时绑定不同参数，可以反复使用，提高效率


## 处理结果集

### 通用结果集 ResultSet

```
while(rs.next()) {
 // by index
 val idbn = rs.getString(1)
// by column name
 val price = rs.getDouble("price")
}
```
可以通过下标遍历遍历列，从1开始。
也可以通过列名遍历。
另外 可以获取结果集的元数据。

### 获取自动生成键

```
stat.executeUpdate(insertStatement, Statement.RETURN_GENERATED_KEYS)
val rs = stat.generatedKeys
if ( rs.next() ) {
  val key = rs.getInt(1)
}
```

### 可滚动的结果集

```
val stat = conn.createStatement(type, concurrency)
```

### 行集 RowSet
与ResultSet不同，使用RowSet可以缓存结果，不需要与数据库始终保持连接。


## 元数据

```
val meta : DatabaseMetaData = conn.metaData

val rsmeta : ResultSetMetaData  = rs.metaData
```
DatabaseMetaData用于提供数据库相关的元数据
ResultSetMetaData用于提供结果集相关的元数据

## 事务

```
conn.setAutoCommit(false)
val stat = conn.createStatement()
stat.executeUpdate(cmd1)
stat.executeUpdate(cmd2)

conn.commit() // conn.rollback()

```
默认下 数据库连接处于自动提交模式。


----
>参考：
> [https://docs.oracle.com/javase/tutorial/jdbc/overview/index.html](https://docs.oracle.com/javase/tutorial/jdbc/overview/index.html)
