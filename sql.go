package main

var sql = "
CREATE TABLE IF NOT EXISTS `navs`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `name` VARCHAR(255) NOT NULL,
   `desc` VARCHAR(255) NOT NULL,
   `url` VARCHAR(255) NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE TABLE `links` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`type` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1',
	`catename` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	`name` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`url` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`logo` VARCHAR(1024) NOT NULL COLLATE 'utf8_general_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=10025
;
"