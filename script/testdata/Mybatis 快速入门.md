# 前言

介绍mybatis框架的基本使用方法。
更详细的源码学习 参考另外一篇文章: [mybatis mybatis-spring 源码流程分析](https://www.jianshu.com/p/434458ee9f9e)


## [官网](http://www.mybatis.org/mybatis-3/zh/sqlmap-xml.html)
---
## 创建maven工程
>pom.xml
#### 设置java版本
```
<properties>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
        <maven.compiler.source>9</maven.compiler.source>
        <maven.compiler.target>9</maven.compiler.target>
    </properties>
```
#### 添加mybatis, mysql
``` 
<dependency>
            <groupId>org.mybatis</groupId>
            <artifactId>mybatis</artifactId>
            <version>3.4.6</version>
</dependency>
<dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>8.0.12</version>
</dependency>
```
---
## 创建db.properties  (resource/)
```
username=gen
password=123
driver=com.mysql.cj.jdbc.Driver
url=jdbc:mysql://192.168.199.217:3306/awesome?useUnicode=true&characterEncoding=utf8&useSSL=true
```
---
## 创建 mybatis-config.xml  (resource/)
```
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE configuration
        PUBLIC "-//mybatis.org//DTD Config 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-config.dtd">
<configuration>

    <properties resource="jdbc.properties">
    </properties>
    
    <settings>
        <setting name="mapUnderscoreToCamelCase" value="true"/>
    </settings>

    <environments default="development">
        <environment id="development">
            <transactionManager type="JDBC"/>
            <dataSource type="POOLED">
                <property name="driver" value="${driver}"/>
                <property name="url" value="${url}"/>
                <property name="username" value="${username}"/>
                <property name="password" value="${password}"/>
            </dataSource>
        </environment>
    </environments>
    <mappers>
        <mapper resource="mapper/BlogMapper.xml"/>
    </mappers>
</configuration>
```
## 创建实体类 BlogEntity.java
```
package com.gen.mybatis.entity;

public class BlogEntity {

    private String id;
    private String userId;
    private String userName;
    private String name;
    private String summary;
    private String content;
    private String createAt;

    @Override
    public String toString() {
        return "BlogEntity{" +
                "id='" + id + '\'' +
                ", userId='" + userId + '\'' +
                ", userName='" + userName + '\'' +
                ", name='" + name + '\'' +
                ", summary='" + summary + '\'' +
                ", content='" + content + '\'' +
                ", createAt='" + createAt + '\'' +
                '}';
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getSummary() {
        return summary;
    }

    public void setSummary(String summary) {
        this.summary = summary;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public String getCreateAt() {
        return createAt;
    }

    public void setCreateAt(String createAt) {
        this.createAt = createAt;
    }
}
```
---
## 创建 BlogMapper.java 接口
```
package com.gen.mybatis.dao;

import com.gen.mybatis.entity.BlogEntity;

public interface BlogMapper {

    BlogEntity queryById(String id);
}
```
---
## 创建 BlogMapper.xml (resource/mapper/)

```
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper
        PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN"
        "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="com.gen.mybatis.dao.BlogMapper">
    <select id="queryById" resultType="com.gen.mybatis.entity.BlogEntity">
    select * from blogs where id = #{id}
  </select>
</mapper>
```
---
## 使用BlogMapper.java，创建BlogDao.java
```
package com.gen.mybatis.dao;

import com.gen.mybatis.entity.BlogEntity;
import org.apache.ibatis.io.Resources;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.apache.ibatis.session.SqlSessionFactoryBuilder;

import java.io.IOException;
import java.io.InputStream;

public class BlogDao {

    private SqlSessionFactory sqlSessionFactory;

    public BlogDao() {
        String res = "mybatis-config.xml";
        InputStream is = null;
        try {
            is = Resources.getResourceAsStream(res);
        } catch (IOException e) {
            e.printStackTrace();
        }
        sqlSessionFactory =
                new SqlSessionFactoryBuilder().build(is);
    }

    public BlogEntity queryById(String id) {
        try (SqlSession session = sqlSessionFactory.openSession()) {
            BlogMapper mapper = session.getMapper(BlogMapper.class);
            return mapper.queryById(id);
        }
    }

}
```
---
## 测试 BlogDaoTest.java
```
package com.gen.mybatis.dao;

import com.gen.mybatis.entity.BlogEntity;

import java.util.logging.Logger;

public class BlogDaoTest {

    private static final Logger logger = Logger.getLogger("BlogDaoTest");

    @org.junit.Test
    public void queryById() {
        BlogDao dao = new BlogDao();

        BlogEntity blogEntity = dao.queryById("1");

        logger.info(blogEntity.toString());
    }
}
```
