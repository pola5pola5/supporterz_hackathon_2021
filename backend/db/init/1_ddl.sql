DROP SCHEMA IF EXISTS `trip_db`;
CREATE SCHEMA `trip_db` DEFAULT CHARACTER SET utf8mb4 ;
USE `trip_db`;

DROP TABLE IF EXISTS `user_table`;
DROP TABLE IF EXISTS `trip_table`;
DROP TABLE IF EXISTS `img_table`;

CREATE TABLE IF NOT EXISTS `trip_db`.`user_table` (
    `user_id` VARCHAR(128) NOT NULL COMMENT "ユーザID",
    `user_name` VARCHAR(128) NOT NULL COMMENT "ユーザ名",
    `password` VARCHAR(128) NOT NULL COMMENT "パスワード",
    `auth_token` VARCHAR(128) NOT NULL COMMENT "認証トークン",
    PRIMARY KEY(`user_id`)
);

CREATE TABLE IF NOT EXISTS `trip_db`.`trip_table` (
    `trip_id` VARCHAR(128) NOT NULL COMMENT "旅ID",
    `user_id` VARCHAR(128) NOT NULL COMMENT "ユーザID",
    `start_date` DATETIME NOT NULL COMMENT "開始日",
    `end_date` DATETIME NOT NULL COMMENT "終了日",
    PRIMARY KEY(`trip_id`)
);

CREATE TABLE IF NOT EXISTS `trip_db`.`img_table` (
    `img_id` VARCHAR(128) NOT NULL COMMENT "画像ID",
    `trip_id` VARCHAR(128) NOT NULL COMMENT "旅ID",
    `img_url` VARCHAR(256) NOT NULL COMMENT "画像URL",
    `latitude` VARCHAR(128) NOT NULL COMMENT "緯度",
    `longitude` VARCHAR(128) NOT NULL COMMENT "経度",
    PRIMARY KEY(`img_id`)
);