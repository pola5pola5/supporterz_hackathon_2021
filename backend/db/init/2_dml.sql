USE `trip_db`;
SET NAMES utf8mb4;

INSERT INTO `user_table` (`user_id`,`user_name`,`password`,`auth_token`) VALUES ("c8fed8a1-7e15-4516-838a-14a6bc1f703f","aaa","111","4f272392-a679-4a86-ba92-3ffd96c83ae8");
INSERT INTO `user_table` (`user_id`,`user_name`,`password`,`auth_token`) VALUES ("090ae791-9f53-4a84-ad61-0e84aee08634","bbb","222","0a5daf57-1fa8-452e-99d9-426718cabe8b");

INSERT INTO `trip_table` (`trip_id`,`user_id`) VALUES ("ade01599-4acb-40a8-870c-4b9eea358d22","090ae791-9f53-4a84-ad61-0e84aee08634");
INSERT INTO `trip_table` (`trip_id`,`user_id`) VALUES ("c3780053-40be-4bce-9597-887559d9beee","090ae791-9f53-4a84-ad61-0e84aee08634");

INSERT INTO `img_table` (`img_id`,`trip_id`,`img_url`,`latitude`,`longitude`,`date_time`) VALUES ("922a7cfe-9c95-478c-a3aa-08df6e7bca3c","ade01599-4acb-40a8-870c-4b9eea358d22","url1",139.88769444444443,35.774116666666664,"2021-05-29 18:59:19");
INSERT INTO `img_table` (`img_id`,`trip_id`,`img_url`,`latitude`,`longitude`,`date_time`) VALUES ("bd81f5a7-2aee-41e9-ad93-c1f3bf7c754e","c3780053-40be-4bce-9597-887559d9beee","url2",139.88769444444443,35.774116666666664,"2021-05-29 18:59:19");
