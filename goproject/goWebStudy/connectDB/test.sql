CREATE DATABASE `test`;

CREATE TABLE `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(32) UNIQUE NOT NULL,
    `password` VARCHAR(32) NOT NULL,
    `e_mail` VARCHAR(32)
);