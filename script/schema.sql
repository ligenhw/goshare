
create database goshare;

use goshare;

create table user (
    id Int auto_increment,
    user_name varchar(20) not null,
    password varchar(20) not null,
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (id),
    UNIQUE KEY (user_name)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table profile (
    id Int auto_increment,
    user_id Int,
    gender varchar(20),
    age int,
    address varchar(40),
    email varchar(20),
    PRIMARY KEY (id),
    foreign key (user_id) REFERENCES user(id) ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table blog (
    id Int auto_increment,
    user_id Int,
    title varchar(20) not null,
    content mediumtext not null,
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    INDEX inx_time (time)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table comment (
    id Int auto_increment,
    blog_id Int,
    user_id Int,
    content varchar(40),
    time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (blog_id) REFERENCES blog(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    INDEX inx_time (time)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

