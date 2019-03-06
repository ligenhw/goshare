
insert into user (user_name, password) values
 ('gen', '123'),
 ('root', '1234');

insert into profile (user_id, gender, age, address) values
 (1, 'male', 28, '北京');


select * from user, profile where user.id = 1;


# select * from user left join profile where user.id = 1;

insert into blog (user_id, title, content) values
 (1, 'react入门', 
 "
 ## 简介
 * js框架
 * 声明式
 * 组件化

## 例子
 ```
class HelloMessage extends React.Component {
  
  }
 ```
 "
 );

insert into blog (user_id, title, content) values
 (2, 'golang', 
 "
 ## 简介
 * 函数式
 * 声明式
 * 模块化

## 例子
 ```
 fmt.Println(\"hello\")
 ```
 "
 );

 insert into blog (user_id, title, content) values
 (1, 'test title', 
 "
 ## test content.
 "
 );