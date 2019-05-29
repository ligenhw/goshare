
alter table user add column avatar_url varchar(40);
alter table user modify column avatar_url char(100);

alter table comment add column parent_id int;
alter table comment add column reply_to int;

alter table comment add constraint foreign key(parent_id) REFERENCES comment(id);
alter table comment add constraint foreign key(reply_to) REFERENCES user(id);





