CREATE DATABASE ECOM;
       USE ECOM;
           CREATE TABLE TEST_TABLE(
            StudentId int not null AUTO_INCREMENT,
            FirstName varchar(50) not null,
            PRIMARY KEY(StudentId)
        );
INSERT INTO TEST_TABLE(FirstName) VALUES("John")