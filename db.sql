CREATE TABLE IF NOT EXISTS `bb_catalog` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `ident` varchar(255) NOT NULL UNIQUE,
    `name` varchar(255) NOT NULL,
    `resume` varchar(255) NOT NULL,
    `display_order` integer NOT NULL,
    `img_url` varchar(255) NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE IF NOT EXISTS `bb_blog` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `ident` varchar(255) NOT NULL UNIQUE,
   `title` varchar(255) NOT NULL,
   `keywords` varchar(255),
   `catalog_id` bigint NOT NULL,
   `blog_content_id` bigint NOT NULL UNIQUE,
   `blog_content_last_update` bigint NOT NULL,
   `type` tinyint NOT NULL,
   `status` tinyint NOT NULL,
   `views` bigint NOT NULL,
   `created` datetime NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
CREATE INDEX `bb_blog_catalog_id` ON `bb_blog` (`catalog_id`);


CREATE TABLE IF NOT EXISTS `bb_blog_content` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `content` longtext NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE IF NOT EXISTS `newemployee_catalog` (
    `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `ident` varchar(255) NOT NULL UNIQUE,
    `name` varchar(255) NOT NULL,
    `resume` varchar(255) NOT NULL,
    `display_order` integer NOT NULL,
    `img_url` varchar(255) NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE IF NOT EXISTS `newemployee_blog` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `ident` varchar(255) NOT NULL UNIQUE,
   `title` varchar(255) NOT NULL,
   `keywords` varchar(255),
   `catalog_id` bigint NOT NULL,
   `blog_content_id` bigint NOT NULL UNIQUE,
   `blog_content_last_update` bigint NOT NULL,
   `type` tinyint NOT NULL,
   `status` tinyint NOT NULL,
   `views` bigint NOT NULL,
   `created` datetime NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
CREATE INDEX `newemployee_blog_catalog_id` ON `newemployee_blog` (`catalog_id`);


CREATE TABLE IF NOT EXISTS `newemployee_blog_content` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `content` longtext NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE IF NOT EXISTS `capability_map` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `capability_id` bigint NOT NULL,
   `capability` varchar(255) NOT NULL,
   `class` varchar(255) NOT NULL,
   `sub_class` varchar(255) NOT NULL,
   `description` varchar(255),
   `expert_pool` varchar(255)  
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
CREATE INDEX `capability_map_catalog_id` ON `capability_map` (`capability_id`);


CREATE TABLE IF NOT EXISTS `capabilities` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `user_id` bigint NOT NULL,
   `user_name` varchar(255) NOT NULL,
   `capability_id` bigint NOT NULL,
   `level` varchar(255)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;




#training 表项创建
CREATE TABLE IF NOT EXISTS `training_schedule_publish` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `title` varchar(255) NOT NULL,
   `teacher` varchar(255) NOT NULL,
   `content` varchar(255),
   `description` varchar(255),
   `start_date` datetime,
   `room` varchar(255) NOT NULL,
   `status` varchar(255)  NOT NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


CREATE TABLE IF NOT EXISTS `training_schedule_collect` (
   `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
   `title` varchar(255) NOT NULL,
   `description` varchar(255) NOT NULL,
   `fields` varchar(255) NOT NULL,
   `important_level` varchar(255) NOT NULL,
   `expect_date` datetime
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
