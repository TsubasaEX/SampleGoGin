create database `Demo`;

create table `Demo`.users (
	id int,
    name varchar(45),
    password varchar(45),
    email varchar(45),
    primary key(id)
);

insert into `Demo`.users (id, name, password, email) values
(1, 'Wilson','123456', 'wilson@gmail.com'),
(2, 'Tom','123456', 'tom@gmail.com'),
(3, 'Sherry','123456', 'sherry@gmail.com');

select * from `Demo`.users;
