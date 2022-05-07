CREATE DATABASE `meshnote`;
USE `meshnote`;

CREATE TABLE IF NOT EXISTS `user`(
   `id` INT UNSIGNED AUTO_INCREMENT UNIQUE,
   `name` VARCHAR(255) NOT NULL,
   `password` VARCHAR(63) NOT NULL,
   PRIMARY KEY(`id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `tree`(
   `id` INT UNSIGNED AUTO_INCREMENT UNIQUE,
   `name` VARCHAR(255) NOT NULL,
   `type` INT NOT NULL,
   `filename` VARCHAR(255) NOT NULL,
   `filesize` INT NOT NULL,
   `parent` INT NOT NULL,
   `uptime` VARCHAR(255) NOT NULL,
   PRIMARY KEY(`id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `log`(
   `id` INT UNSIGNED AUTO_INCREMENT UNIQUE,
   `user` VARCHAR(255) NOT NULL,
   `event` VARCHAR(255) NOT NULL,
   `detail` VARCHAR(65535) NOT NULL,
   PRIMARY KEY(`id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8;

INSERT INTO `tree` (`id`, `name`, `type`, `filename`, `filesize`, `parent`, `uptime`) VALUES (NULL, ?, ?, ?, ?, ?, ?);