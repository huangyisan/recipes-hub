CREATE TABLE recipes (
                         id INT AUTO_INCREMENT COMMENT '菜谱ID',
                         recipe_name VARCHAR(255) NOT NULL DEFAULT '番茄炒蛋' COMMENT '菜名，唯一约束',
                         instructions TEXT COMMENT '烹饪步骤',
                         cooking_time INT NOT NULL DEFAULT 5 COMMENT '烹饪时间（以分钟为单位）',
                         difficulty VARCHAR(50) NOT NULL DEFAULT '简单' COMMENT '难度等级',
                         recipe_type TINYINT NOT NULL DEFAULT 1 COMMENT '菜谱类型',
                         image_content VARCHAR(255) COMMENT '图片的URL地址',
                         creation_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         last_updated DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         last_cooked_date DATETIME COMMENT '最后一次烹饪时间',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `idx_recipe_name` (`recipe_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';

CREATE TABLE cuisine (
                         id INT AUTO_INCREMENT COMMENT '菜系ID',
                         cuisine_name VARCHAR(255) NOT NULL UNIQUE COMMENT '菜系名称',
                         description TEXT COMMENT '菜系描述',
                         PRIMARY KEY(`id`),
                        UNIQUE KEY `idx_cuisine_name` (`cuisine_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='菜单表';