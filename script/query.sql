
--查询评论列表 包含回复
select c1.id parent_id, c1.content parent_content, c1.time parent_time,
 c2.id, c2.user_id, c2.reply_to, c2.content, c2.time from
 comment c1 left join comment c2 on c2.parent_id = c1.id
 where c1.parent_id is null and c1.blog_id = ? ( c2.blog_id = ? or c2.blog_id is null)
 order by c1.time desc, c2.time ;

--两侧父子关系查询 按照c1.id排序 确保子回复按照主评论分组。
select * from
 comment c1 left join comment c2 on c2.parent_id = c1.id
 where c1.parent_id is null and c1.blog_id = 10 and ( c2.blog_id = 10 or c2.blog_id is null)
 order by c1.id ,c1.time desc, c2.time ;


