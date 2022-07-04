create database RentBookGroupProject;

use RentBookGroupProject;

create table Users (
	user_id int primary key,
	user_name varchar(50) not null,
	user_email varchar(50) not null,
	password varchar(50)  not null
);

create table Books (
	book_id int primary key,
	user_id int,
	book_name varchar(50) not null,
	book_type varchar(50) not null,
	book_status bool not null,
	foreign key (user_id) references Users(user_id)
);

create table Rents (
	rent_id int primary key,
	user_id int,
	book_id int,
	created_at datetime not null
);