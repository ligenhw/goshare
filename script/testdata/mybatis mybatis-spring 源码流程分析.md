# 前言

实际项目中通常使用mybatis-spring获得mapper的bean对象。本文通过 mybatis 和 mybatis-spring 的源码流程 了解其实现方式。

![mybatis-logo](http://upload-images.jianshu.io/upload_images/3991621-4b107dfb4bca99f5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## mybatis  doc
[http://www.mybatis.org/mybatis-3/zh/getting-started.html](http://www.mybatis.org/mybatis-3/zh/getting-started.html)


## mybatis-spring doc
[http://www.mybatis.org/spring/zh/index.html](http://www.mybatis.org/spring/zh/index.html)

## mybatis-spring github
[https://github.com/mybatis/spring](https://github.com/mybatis/spring)

---
## mybatis 源码流程

### session & binding

```
SqlSession session = sqlSessionFactory.openSession();
try {
  BlogMapper mapper = session.getMapper(BlogMapper.class);
  Blog blog = mapper.selectBlog(101);
} finally {
  session.close();
}
```

第一部分：构造sessionFactory

第二部分 查询
* 1. sqlsession.getMapper() 返回的是接口的 动态代理 MapperProxy
* 2. 接口调用时 MapperProxy.invoke()调用，根据方法名获取MapperMethod
* 3. 执行MapperMethod.execute()  根据mapper标签区分CRUD操作，调用sqlSession。


### transaction
* 1. Transaction接口 包装了connection, 管理其声明周期，创建，提交，回滚，关闭。  
* 2. 有两种实现， JDBC事务 和 Managed交由spring处理。


## datasource
* 1.  POOLED 带连接池的 数据源
PooledDataSource.getconnection() 调用popConnection()从空闲连接中获取**PooledConnection**对象。
如果idleConnections非空，直接取其中的一个返回。
如果activeConnections 活跃连接数小于 配置的上限，则创建一个新的连接返回。
如果无可用连接，等待后再次获取。
最后 获取将准备返回的conn记录到activeConnections中。

PooledConnection是Connection的一个动态代理。主要为了在调用connection对象的接口前，检查连接是否过期。另外，如果调用close，会将connect归还连接池。

* 2. UNPOOLED 不带连接池 的数据源
UnpooledDataSource.getconnection() 使用JDBC API 从 DriverManager获取连接。
PooledDataSource也是通过UnpooledDataSource获取底层的连接。
---
## plugin
```
// ExamplePlugin.java
@Intercepts({@Signature(
  type= Executor.class,
  method = "update",
  args = {MappedStatement.class,Object.class})})
public class ExamplePlugin implements Interceptor {
  public Object intercept(Invocation invocation) throws Throwable {
    return invocation.proceed();
  }
  public Object plugin(Object target) {
    return Plugin.wrap(target, this);
  }
  public void setProperties(Properties properties) {
  }
}
```

* 1. 实现Interceptor 接口，并通过Intercepts注解 指明需要拦截的方法。可以拦截 Executor ，ParameterHandler ，ResultSetHandler ，StatementHandler 的方法。
* 2.  解析配置文件时，解析plugin标签，生成一个拦截器的列表 InterceptorChain对象。
* 3. 通过Configuration工厂方法 newExecutor  newStatementHandler  newParameterHandler  newResultSetHandler 中会 通过 interceptorChain.pluginAll 应用所有拦截器，得到代理后的对象。供后续流程使用。 pluginall迭代调用所有插件的plugin方法。
* 4. Plugin提供工具方法wrap用于，生成传入对象的动态代理。 第一个参数为 被代理对象。第二个参数为 Interceptor （plugin）对象本身。
* 5. 当拦截方法调用时，会执行Interceptor的intercept方法，传入的Invocation 对象 封装了代理方法的调用过程，通过调用proceed方法 使之执行。

---

## binding
生成mapper接口的 动态代理。
实现方式与 retrofit 相同。 [https://www.jianshu.com/p/122859d42f4f](https://www.jianshu.com/p/122859d42f4f)

* 1. MapperProxyFactory<T> 用于生成mapper的动态代理对象。
* 2.  MapperProxy<T> mapper的动态代理。保存了接口中的所有方法map methodCache。
```
@Override
  public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
    try {
      if (Object.class.equals(method.getDeclaringClass())) {
        return method.invoke(this, args);
      } else if (isDefaultMethod(method)) {
        return invokeDefaultMethod(proxy, method, args);
      }
    } catch (Throwable t) {
      throw ExceptionUtil.unwrapThrowable(t);
    }
    final MapperMethod mapperMethod = cachedMapperMethod(method);
    return mapperMethod.execute(sqlSession, args);
  }
```
接口调用时invoke根据传入的method 在methodCache中查找方法，调用MapperMethod.execute()

* 3. MapperMethod 根据 interface method 到configuration中找到MappedStatement。根据MappedStatement的类型，query,insert,update,delete 在execute时 将方法路由到sqlsession 相应的实现上。

---

## executor

sql的执行器，封装了对jdbc的调用。
sqlsession调用executor实现 对底层sql的执行。

* 1. statement 子包， 封装了jdbc Statement 对象 的执行。StatementHandler接口。

---

## mybatis-spring 源码流程
* 1. SqlSessionTemplate类
线程安全的SqlSession实现
使用内部类SqlSessionInterceptor作为动态代理，invoke调用时
首先从事务管理器获取sqlsession，如果是在一个事务中则sql是ThreadLocal中存储的sqlsession。
如果不在事务中，则通过sqlSessionFactory.openSession()新创建一个sqlsession对象。
最后调用registerSessionHolder(sqlsession) , 如果开启事务，将sqlsession注册。

## transaction
实现mybatis包中的事务接口， 提供 spring 管理的 事务。
使用了 spring-jdbc spring-tx