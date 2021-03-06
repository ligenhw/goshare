
create database goshare;

use goshare;

create table user (
    id Int auto_increment,
    user_name varchar(20) not null,
    password varchar(20) not null,
    avatar_url varchar(100) DEFAULT '',
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (id),
    UNIQUE KEY (user_name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table profile (
    id Int auto_increment,
    openid varchar(50),
    auth_type varchar(20),
    user_id Int,
    content varchar(200),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table blog (
    id Int auto_increment,
    user_id Int,
    title varchar(60) not null,
    content mediumtext not null,
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    INDEX inx_time (time)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table comment (
    id Int auto_increment,
    blog_id Int,
    user_id Int,
    parent_id Int,
    reply_to Int,
    content varchar(500),
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (blog_id) REFERENCES blog(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES comment(id) ON DELETE CASCADE,
    FOREIGN KEY (reply_to) REFERENCES user(id) ON DELETE CASCADE,
    INDEX inx_time (time),
    INDEX inx_blog_id (blog_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table tag (
    id Int auto_increment,
    name varchar(20) not null,
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table tag_blog_rel (
    id Int auto_increment,
    tag_id Int,
    blog_id Int,
    primary key (id),
    FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE,
    FOREIGN KEY (blog_id) REFERENCES blog(id) ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


