CREATE TABLE ingredients (
    id INT AUTO_INCREMENT COMMENT '食材ID',
    ingredient_name VARCHAR(255) NOT NULL UNIQUE COMMENT '食材名称',
    ingredient_image_content VARCHAR(255) NOT NULL DEFAULT '图片地址' COMMENT '食材图片地址',
    ingredient_description TEXT NOT NULL DEFAULT '食材描述' COMMENT '食材描述',
    PRIMARY KEY (id)
);

CREATE TABLE recipe_ingredients (
    id INT AUTO_INCREMENT COMMENT '菜谱ID',
    ingredient_id INT COMMENT '食材ID',
    UNIQUE (id, ingredient_id),
    PRIMARY KEY(id)
);
