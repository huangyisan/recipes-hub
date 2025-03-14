CREATE TABLE ingredients (
    ingredient_id INT AUTO_INCREMENT COMMENT '食材ID',
    ingredient_name VARCHAR(255) NOT NULL UNIQUE COMMENT '食材名称',
    PRIMARY KEY (ingredient_id)
);

CREATE TABLE recipe_ingredients (
    recipe_id INT AUTO_INCREMENT COMMENT '菜谱ID',
    ingredient_id INT COMMENT '食材ID',
    UNIQUE (recipe_id, ingredient_id),
    PRIMARY KEY(recipe_id)
);
