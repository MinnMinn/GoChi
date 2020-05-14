CREATE DATABASE IF NOT EXISTS `golang` DEFAULT CHARACTER SET utf8mb4;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `golang`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));

DROP TABLE IF EXISTS `connection`;
CREATE TABLE `golang`.`connection` (
  `user_id` INT NOT NULL,
  `connect_id` INT NOT NULL,
  PRIMARY KEY (`user_id`, `connect_id`));

DROP TABLE IF EXISTS `follow`;
CREATE TABLE `golang`.`follow` (
  `user_id` INT NOT NULL,
  `follow_id` INT NOT NULL,
  PRIMARY KEY (`user_id`, `follow_id`));

DROP TABLE IF EXISTS `block`;
CREATE TABLE `golang`.`block` (
  `user_id` INT NOT NULL,
  `block_id` INT NOT NULL,
  PRIMARY KEY (`user_id`, `block_id`));

INSERT INTO `golang`.`user` (`email`) VALUES ('andy@example.com');
INSERT INTO `golang`.`user` (`email`) VALUES ('john@example.com');
INSERT INTO `golang`.`user` (`email`) VALUES ('lisa@example.com');
INSERT INTO `golang`.`user` (`email`) VALUES ('common@example.com');

INSERT INTO `golang`.`connection` (`user_id`, `connect_id`) VALUES ('1', '3');
INSERT INTO `golang`.`connection` (`user_id`, `connect_id`) VALUES ('3', '1');
INSERT INTO `golang`.`connection` (`user_id`, `connect_id`) VALUES ('2', '3');
INSERT INTO `golang`.`connection` (`user_id`, `connect_id`) VALUES ('3', '2');
