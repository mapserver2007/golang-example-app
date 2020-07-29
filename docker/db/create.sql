create table users (
	id int not null auto_increment,
	name varchar(30) not null,
	age int not null,
  PRIMARY KEY (id)
);

create table items (
	id int not null auto_increment,
	name varchar(30) not null,
	price int not null,
  PRIMARY KEY (id)
);