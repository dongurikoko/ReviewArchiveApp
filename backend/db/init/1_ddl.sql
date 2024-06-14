-- MySQL Script generated by MySQL Workbench
-- Fri Feb 14 23:09:20 2020
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema review_archive_api
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `review_archive_api` DEFAULT CHARACTER SET utf8mb4 ;
USE `review_archive_api` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `review_archive_api`.`Users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `review_archive_api`.`Users` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT 'ユーザID',
  `uid` TEXT NOT NULL COMMENT 'UID',
  PRIMARY KEY (`id`))
ENGINE=InnoDB
COMMENT = 'ユーザ';

-- -----------------------------------------------------
-- Table `review_archive_api`.`Contents`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `review_archive_api`.`Contents`(
  `id` INT NOT NULL AUTO_INCREMENT COMMENT 'コンテンツID',
  `title` VARCHAR(255) NOT NULL COMMENT 'タイトル名',
  `before_code` TEXT  COMMENT '修正前コード',
  `after_code` TEXT  COMMENT '修正後コード',
  `review` TEXT  COMMENT 'レビュー内容',
  `memo` TEXT  COMMENT 'メモ',
  `user_id` INT NOT NULL COMMENT 'ユーザID',
  PRIMARY KEY (`id`),
  FOREIGN KEY (user_id) REFERENCES Users(id))
ENGINE = InnoDB
COMMENT = 'コンテンツ';

-- -----------------------------------------------------
-- Table `review_archive_api`.`Keywords`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `review_archive_api`.`Keywords` (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT 'キーワードID',
  `keyword` VARCHAR(255) NOT NULL COMMENT 'キーワード',
  PRIMARY KEY (`id`))
ENGINE=InnoDB
COMMENT = 'キーワード';

-- -----------------------------------------------------
-- Table `review_archive_api`.`Tagging`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `review_archive_api`.`Tagging` (
  `content_id` INT NOT NULL COMMENT 'コンテンツID',
  `keyword_id` INT NOT NULL COMMENT 'キーワードID',
  PRIMARY KEY (content_id, keyword_id),
  FOREIGN KEY (content_id) REFERENCES Contents(id),
  FOREIGN KEY (keyword_id) REFERENCES Keywords(id)
)
ENGINE=InnoDB
COMMENT = 'タグ';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
