
create database Students

use Students

drop database Students

create table `classes`(
	id bigint unsigned auto_increment,
	class_name varchar(50),
	primary key(id)
)

create table `students`(
	id bigint unsigned auto_increment ,
	student_name varchar(50),
	class_id bigint unsigned,
	foreign key (class_id) references classes(id),
	primary key(id, class_id)
)

create table `subjects`(
	id bigint unsigned auto_increment,
	subject_name varchar(50),
	primary key(id)
)


create table `studentsubject`(
	subject_id bigint unsigned,
	student_id bigint unsigned,
	primary key (student_id,subject_id),
	foreign key (student_id) references students(id),
	foreign key (subject_id) references subjects(id)
	
)

insert into classes (id, class_name) values (1, 'Sales');
insert into classes (id, class_name) values (2, 'ProductManager');
insert into classes (id, class_name) values (3, 'Services');
insert into classes (id, class_name) values (4, 'Service');
insert into classes (id, class_name) values (5, 'HR');
insert into classes (id, class_name) values (6, 'BusinessDev');
insert into classes (id, class_name) values (7, 'Marketing');
insert into classes (id, class_name) values (8, 'Training');
insert into classes (id, class_name) values (9, 'Engineering');
insert into classes (id, class_name) values (10, 'DA');

DELETE  FROM student WHERE student_id=1

insert into students (id, student_name, class_id) values (1, 'Tonks', 1);
insert into students (id, student_name, class_id) values (2, 'Columbell', 2);
insert into students (id, student_name, class_id) values (3, 'Skittrall', 3);
insert into students (id, student_name, class_id) values (4, 'Dinsdale', 4);
insert into students (id, student_name, class_id) values (5, 'Kyrkeman', 5);
insert into students (id, student_name, class_id) values (6, 'Harmar', 6);
insert into students (id, student_name, class_id) values (7, 'Kineton', 7);
insert into students (id, student_name, class_id) values (8, 'Sacaze', 8);
insert into students (id, student_name, class_id) values (9, 'Meindl', 9);
insert into students (id, student_name, class_id) values (10, 'Renney', 10);

insert into subjects (id, subject_name) values (1, 'Gabcube');
insert into subjects (id, subject_name) values (2, 'Leexo');
insert into subjects (id, subject_name) values (3, 'Dabvine');
insert into subjects (id, subject_name) values (4, 'Eazzy');
insert into subjects (id, subject_name) values (5, 'Kaymbo');
insert into subjects (id, subject_name) values (6, 'Lajo');
insert into subjects (id, subject_name) values (7, 'Tavu');
insert into subjects (id, subject_name) values (8, 'Devshare');
insert into subjects (id, subject_name) values (9, 'Twiyo');
insert into subjects (id, subject_name) values (10, 'Babbleset');
insert into subjects (id, subject_name) values (11, 'Fiveclub');
insert into subjects (id, subject_name) values (12, 'Flipstorm');
insert into subjects (id, subject_name) values (13, 'Bubblemix');
insert into subjects (id, subject_name) values (14, 'Skinte');
insert into subjects (id, subject_name) values (15, 'Zoomcast');
insert into subjects (id, subject_name) values (16, 'Trilia');
insert into subjects (id, subject_name) values (17, 'Skibox');
insert into subjects (id, subject_name) values (18, 'Jazzy');
insert into subjects (id, subject_name) values (19, 'Blognation');

insert into studentsubject (subject_id, student_id) values (1, 1);
insert into studentsubject (subject_id, student_id) values (1, 2);
insert into studentsubject (subject_id, student_id) values (1, 3);
insert into studentsubject (subject_id, student_id) values (2, 4);
insert into studentsubject (subject_id, student_id) values (2, 5);
insert into studentsubject (subject_id, student_id) values (2, 6);
insert into studentsubject (subject_id, student_id) values (3, 7);
insert into studentsubject (subject_id, student_id) values (3, 8);
insert into studentsubject (subject_id, student_id) values (3, 9);
insert into studentsubject (subject_id, student_id) values (4, 10);

insert into studentsubject (student_id, subject_id) values (4, 11);
insert into studentsubject (student_id, subject_id) values (5, 12);
insert into studentsubject (student_id, subject_id) values (5, 13);
insert into studentsubject (student_id, subject_id) values (6, 14);
insert into studentsubject (student_id, subject_id) values (6, 15);
insert into studentsubject (student_id, subject_id) values (7, 16);
insert into studentsubject (student_id, subject_id) values (8, 17);
insert into studentsubject (student_id, subject_id) values (9, 18);
insert into studentsubject (student_id, subject_id) values (10, 19);

