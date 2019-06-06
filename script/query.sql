
--查询评论列表 包含回复
select c1.id parent_id, c1.content parent_content, c1.time parent_time,
 c2.id, c2.user_id, c2.reply_to, c2.content, c2.time from
 comment c1 left join comment c2 on c2.parent_id = c1.id
 where c1.parent_id is null and c1.blog_id = ? ( c2.blog_id = ? or c2.blog_id is null)
 order by c1.time desc, c2.time ;

-- 两侧父子关系查询 按照c1.id排序 确保子回复按照主评论分组。
-- 存在单独的父评论，所以要使用左连接， 主评论记录 on c2.parent_id = c1.id 条件不满足时会将c2所有列为null join到主评论记录下。所以要加入c2.blog_id is null的判断，从而包含所有父子评论
-- 单条结果表示一个子评论以及它的主评论，如果有1个主评论下有n个子评论，结果集中会有n条记录，其中每条结果集的c1部分代表父评论信息，n条记录都相同。c2部分是各自的子评论内容。 
-- c1为父评论所以where条件是c1.parent_id is null

select * from
 comment c1 left join comment c2 on c2.parent_id = c1.id
 where c1.parent_id is null and c1.blog_id = 4 and ( c2.blog_id = 4 or c2.blog_id is null)
 order by c1.id desc, c1.time desc, c2.time ;


-- 简化之后
-- left join 确保单独的父评论 会包含在结果集中
-- c1代表父评论，所以where中条件是 c1.parent_id is null
-- order 先根据主评论时间倒排， 在根据子评论时间正排
select * from
 comment c1 left join comment c2 on c2.parent_id = c1.id
 where c1.parent_id is null and c1.blog_id = 4 
 order by c1.time desc, c2.time;



