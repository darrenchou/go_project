use test1;

create table users (
id int primary key auto_increment,
username varchar(100) unique not null,
password varchar(45) not null,
email varchar(45)
);

create table books(
id int primary key auto_increment,
title varchar(100) not null,
author varchar(100) not null,
publisher varchar(100) not null,
publish_date date not null,
rating int not null,
book_status varchar(20) not null,
price double(11,2) not null,
img_path varchar(100)
)