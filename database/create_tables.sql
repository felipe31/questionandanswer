CREATE DATABASE IF NOT EXISTS questionandanswer;
USE questionandanswer;

CREATE TABLE `Users` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(500) NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE `Questions` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(1000) NOT NULL,
	`body` VARCHAR(10000) NOT NULL,
	`user_fk` INT NOT NULL,
	FOREIGN KEY (`user_fk`) REFERENCES `Users`(`id`),
	PRIMARY KEY (`id`)
);


CREATE TABLE `Answers` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`body` VARCHAR(10000) NOT NULL,
	`user_fk` INT NOT NULL,
	`question_fk` INT NOT NULL,
	PRIMARY KEY (`id`),
	FOREIGN KEY (`user_fk`) REFERENCES `Users`(`id`),
	FOREIGN KEY (`question_fk`) REFERENCES `Questions`(`id`)
);
