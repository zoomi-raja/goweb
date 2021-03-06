CREATE TABLE `goDatabase`.`users`
(
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(45) NOT NULL UNIQUE,
  `user_name` VARCHAR
(45) NOT NULL,
  `password` VARCHAR
(150) NOT NULL,
  `avatar` VARCHAR
(150) NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`));

CREATE TABLE `goDatabase`.`posts`
(
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `title` VARCHAR
(45) NOT NULL,
  `body` LONGTEXT NOT NULL,
  `intro` VARCHAR(250) NOT NULL,
  `meta_data` VARCHAR
(150) NULL,
  `claps` INT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`));

CREATE TABLE `goDatabase`.`comments`
(
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `parent_id` INT NULL DEFAULT 0,
  `post_id` INT NOT NULL,
  `approved` TINYINT NULL DEFAULT 1,
  `text` VARCHAR
(250) NOT NULL,
  `debth_level` INT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY
(`id`));