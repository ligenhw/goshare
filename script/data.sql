-- about界面的 blog_id = -1
insert into comment(blog_id, user_id, content)
 values(-1, 7, 'about comment 1')
-- link界面评论 blog_id = -2
insert into comment(user_id, content)
 values(-2, 7, 'link comment 1')

insert into comment (user_id, content, parent_id, reply_to) 
  values (7, 'about comment3 reply2', -1, 7);