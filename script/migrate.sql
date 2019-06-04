
-- 修改列类型
alter table user modify column avatar_url char(100);

-- 添加列
alter table comment add column parent_id int;
alter table comment add column reply_to int;

-- 添加外建
alter table comment add constraint foreign key(parent_id) REFERENCES comment(id) ON DELETE CASCADE;
alter table comment add constraint foreign key(reply_to) REFERENCES user(id) ON DELETE CASCADE;

-- 查看外建名称
show create table comment;

-- 删除外建
alter table comment drop foreign key comment_ibfk_3;
alter table comment drop foreign key comment_ibfk_4;

-- 修改字符集
ALTER TABLE comment DEFAULT CHARACTER SET utf8mb4;
alter table comment convert to character set utf8mb4;

ALTER TABLE blog DEFAULT CHARACTER SET utf8mb4;
alter table blog convert to character set utf8mb4;
