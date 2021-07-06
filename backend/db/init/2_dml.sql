USE `trip_db`;
SET NAMES utf8mb4;

INSERT INTO `user_table` (`user_id`,`user_name`,`password`,`auth_token`) VALUES ("c8fed8a1-7e15-4516-838a-14a6bc1f703f","user1","password1","4f272392-a679-4a86-ba92-3ffd96c83ae8");
INSERT INTO `user_table` (`user_id`,`user_name`,`password`,`auth_token`) VALUES ("090ae791-9f53-4a84-ad61-0e84aee08634","user2","password2","0a5daf57-1fa8-452e-99d9-426718cabe8b");

INSERT INTO `trip_table` (`trip_id`,`user_id`) VALUES ("ade01599-4acb-40a8-870c-4b9eea358d22","090ae791-9f53-4a84-ad61-0e84aee08634");
INSERT INTO `trip_table` (`trip_id`,`user_id`) VALUES ("c3780053-40be-4bce-9597-887559d9beee","090ae791-9f53-4a84-ad61-0e84aee08634");
INSERT INTO `trip_table` (`trip_id`,`user_id`) VALUES ("38dadef3-db89-46f7-8697-873ea6c3db66","c8fed8a1-7e15-4516-838a-14a6bc1f703f");

INSERT INTO `img_table` (`img_id`,`trip_id`,`img_url`,`longitude`,`latitude`,`date_time`) VALUES ("922a7cfe-9c95-478c-a3aa-08df6e7bca3c","ade01599-4acb-40a8-870c-4b9eea358d22","https://photo-tabi.s3.ap-northeast-1.amazonaws.com/281c3066-3c96-44c3-a4f5-c77a69366f61",139.88769444444443,35.774116666666664,"2021-05-29 18:59:19");
INSERT INTO `img_table` (`img_id`,`trip_id`,`img_url`,`longitude`,`latitude`,`date_time`) VALUES ("bd81f5a7-2aee-41e9-ad93-c1f3bf7c754e","c3780053-40be-4bce-9597-887559d9beee","https://photo-tabi.s3.ap-northeast-1.amazonaws.com/281c3066-3c96-44c3-a4f5-c77a69366f61",139.88769444444443,35.774116666666664,"2021-05-29 18:59:19");
INSERT INTO `img_table` (`img_id`,`trip_id`,`img_url`,`longitude`,`latitude`,`date_time`) VALUES ("3c25695d-05b7-4af2-84c4-b201c619e608","38dadef3-db89-46f7-8697-873ea6c3db66","https://photo-tabi.s3.ap-northeast-1.amazonaws.com/3c25695d-05b7-4af2-84c4-b201c619e608",139.69642777777779,35.66039722222222,"2021-04-16 12:20:11");
INSERT INTO `img_table` (`img_id`,`trip_id`,`img_url`,`longitude`,`latitude`,`date_time`) VALUES ("66ea70bb-9260-4709-af03-7a5925b30389","38dadef3-db89-46f7-8697-873ea6c3db66","https://photo-tabi.s3.ap-northeast-1.amazonaws.com/66ea70bb-9260-4709-af03-7a5925b30389",139.69798333333333,35.662441666666666,"2021-04-15 12:33:39");
INSERT INTO `img_table` (`img_id`,`trip_id`,`img_url`,`longitude`,`latitude`,`date_time`) VALUES ("e53119b4-91a2-473c-8b37-e2d1a1174ba6","38dadef3-db89-46f7-8697-873ea6c3db66","https://photo-tabi.s3.ap-northeast-1.amazonaws.com/e53119b4-91a2-473c-8b37-e2d1a1174ba6",139.69880555555557,35.662330555555556,"2021-04-20 12:32:53");
