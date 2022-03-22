SET NAMES utf8;

-- 导航
CREATE TABLE `navs` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
	`cate` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	`catename` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`name` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`desc` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`url` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`logo` VARCHAR(1024) NOT NULL COLLATE 'utf8_general_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;

-- 链接
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
ENGINE=InnoDB;

-- 用户
CREATE TABLE `users` (
	`id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'uid',
	`account` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '账号' COLLATE 'utf8mb4_general_ci',
	`passwd` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '密码' COLLATE 'utf8mb4_general_ci',
	`create_at` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB;