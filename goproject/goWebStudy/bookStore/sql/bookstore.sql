CREATE DATABASE `bookstore`;

CREATE TABLE `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_name` VARCHAR(32) UNIQUE NOT NULL,
    `pass_word` VARCHAR(32) NOT NULL,
    `e_mail` VARCHAR(32)
);