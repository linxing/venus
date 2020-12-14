CREATE TABLE IF NOT EXISTS `t_users` (
    `id`                    BIGINT(20)    PRIMARY KEY AUTO_INCREMENT NOT NULL,
    `user_name`             VARCHAR(255)  NOT NULL COMMENT '用户名',
    `nickname`              VARCHAR(255)  NOT NULL COMMENT '昵称',
    `avatar`                VARCHAR(255)  NOT NULL COMMENT '头像地址',
    `password`              VARCHAR(255)  NOT NULL COMMENT '密码',
    `phone_number`          VARCHAR(255)  NOT NULL COMMENT '手机号',
    `internal_phone_number` VARCHAR(255)  NOT NULL COMMENT '内部电话号码',
    `role_id`               INT(22)       NOT NULL COMMENT '角色ID',
    `position`              VARCHAR(255)  NOT NULL COMMENT '职位',
    `department_id`         INT(22)       NOT NULL COMMENT '部门ID',
    `department_name`       VARCHAR(255)  NOT NULL COMMENT '部门名',
    `last_login_at`         DATETIME      NULL     COMMENT '最近一次登录时间',
    `disable`               TINYINT       NOT NULL COMMENT '是否禁用该用户'
    `updated_at`            DATETIME      NOT NULL,
    `created_at`            DATETIME      NOT NULL,
    `deleted_at`            DATETIME      NOT NULL DEFAULT '0001-01-01 00:00:00',
    KEY `index_department` (`department_id`),
    UNIQUE KEY `unique_user_name` (`user_name`, `deleted_at`),
    UNIQUE KEY `unique_phone_number` (`phone_number`, `deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET utf8mb4;
