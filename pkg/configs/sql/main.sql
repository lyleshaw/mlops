-- Users table
CREATE TABLE `users`
(
    `user_id`    int unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `user_name`  varchar(32)  NOT NULL COMMENT 'Username',
    `email`      varchar(64)  NOT NULL COMMENT 'Email',
    `password`   varchar(256) NOT NULL COMMENT 'Password',
    `avatar`     varchar(256)          DEFAULT NULL COMMENT 'Avatar',
    `is_active`  tinyint(1)   NOT NULL DEFAULT '0' COMMENT 'Is active',
    `is_admin`   tinyint(1)   NOT NULL DEFAULT '0' COMMENT 'Is admin',
    `is_deleted` tinyint(1)   NOT NULL DEFAULT '0' COMMENT 'Is deleted',
    `create_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create time',
    `update_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `user_name` (`user_name`),
    UNIQUE KEY `email` (`email`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT 'User table';

-- Applications table
CREATE TABLE `applications`
(
    `app_id`     int unsigned                                        NOT NULL AUTO_INCREMENT COMMENT 'Application ID',
    `app_name`   varchar(64)                                         NOT NULL COMMENT 'Application name',
    `app_status` enum ('development', 'online', 'offline', 'closed') NOT NULL DEFAULT 'development' COMMENT 'Application status',
    `app_port`   int                                                 NOT NULL UNIQUE COMMENT 'Application port',
    `app_image`  varchar(256)                                        NOT NULL COMMENT 'Application docker image',
    `user_id`    int unsigned                                        NOT NULL COMMENT 'User ID',
    `is_deleted` tinyint(1)                                          NOT NULL DEFAULT '0' COMMENT 'Is deleted',
    `create_at`  datetime                                            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create time',
    `update_at`  datetime                                            NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Update time',
    PRIMARY KEY (`app_id`),
    KEY `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT 'Application table';