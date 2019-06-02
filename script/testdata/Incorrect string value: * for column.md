## 起因

问题起因是，在插入mysql数据库时报错如下：

Incorrect string value: '\xF0\x9F\x90\x9F' for column 'name' at

## 检查表的字符集

show create table tablename;

发现是 DEFAULT CHARSET=utf8，将其改成utf8mb4。

utf8在mysql中实现为最多存储三个字符，可以编码绝大多数的字符集。而utf8mb4的实现是最多存储4个字节，包含了unicode的完整字符集。所以在涉及到存储 emoji表情，或一些生僻汉字时，需要将数据库的字符集设置为utf8mb4。

对于已经创建的表，通过如下命令修改。

需改表的编码
ALTER TABLE `tablename` DEFAULT CHARACTER SET utf8mb4;

修改单个字段编码
ALTER TABLE `tablename` CHANGE `columnname` VARCHAR(36) CHARACTER SET utf8mb4 NOT NULL;

修改字段的编码
alter table `tablename` convert to character set utf8mb4;

注意⚠️： 只修改表的编码，字段的编码不会变化。

## 检查数据库连接的编码

"xxxx:xxx@tcp(xxx)/mysql?charset=utf8mb4&parseTime=true"

设置连接参数charset=utf8mb4


--

通过修改以上两处内容，可以解决该问题。

>参考资料：
>https://dev.mysql.com/doc/refman/8.0/en/charset-charsets.html
