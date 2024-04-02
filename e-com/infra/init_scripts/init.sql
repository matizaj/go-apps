CREATE DATABASE ECOM;
       USE ECOM;
           CREATE TABLE USERS(
            StudentId int not null AUTO_INCREMENT,
            FirstName varchar(50) not null,
            PRIMARY KEY(StudentId)
        );
INSERT INTO USERS(FirstName) VALUES("John")