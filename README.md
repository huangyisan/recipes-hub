## 菜谱管理
* 添加菜谱: 用户可以添加新的菜谱，包括菜名、食材、步骤、烹饪时间等信息。
* 编辑菜谱: 用户可以修改已有的菜谱信息。
* 删除菜谱: 用户可以删除不再需要的菜谱。

### 菜谱功能
* 菜名
* 食材
* 步骤
* 烹饪时间
* 难度
* 创建时间
* 更新时间
* 最后一次烹饪时间
* 评分

```sql
CREATE TABLE recipes (
                         recipe_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '菜谱ID',
                         recipe_name VARCHAR(255) NOT NULL UNIQUE COMMENT '菜名，唯一约束',
                         instructions TEXT NOT NULL COMMENT '烹饪步骤',
                         cooking_time INT NOT NULL COMMENT '烹饪时间（以分钟为单位）',
                         difficulty VARCHAR(50) NOT NULL COMMENT '难度等级',
                         recipe_type VARCHAR(50) NOT NULL COMMENT '菜谱类型',
                         image_url VARCHAR(255) COMMENT '图片的URL地址',
                         creation_date DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         last_updated DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         last_cooked_date DATETIME COMMENT '最后一次烹饪时间'
);

CREATE INDEX idx_difficulty ON recipes(difficulty);  -- 为难度字段添加索引
CREATE INDEX idx_recipe_type ON recipes(recipe_type);  -- 为菜谱类型字段添加索引

CREATE TABLE ingredients (
    ingredient_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '食材ID',
    ingredient_name VARCHAR(255) NOT NULL UNIQUE COMMENT '食材名称',
    ingredient_image_url VARCHAR(255) COMMENT '食材图片地址',
    ingredient_description TEXT COMMENT '食材描述'
    
);

CREATE TABLE recipe_ingredients (
                                    recipe_id INT COMMENT '菜谱ID',
                                    ingredient_id INT COMMENT '食材ID',
                                    PRIMARY KEY (recipe_id, ingredient_id)
);

CREATE TABLE cuisine (
                              id INT PRIMARY KEY AUTO_INCREMENT COMMENT '菜系ID',
                              name VARCHAR(255) NOT NULL UNIQUE COMMENT '菜系名称',
                              description TEXT COMMENT '菜系描述'
);

CREATE TABLE recipe_ratings (
                                rating_id INT AUTO_INCREMENT PRIMARY KEY COMMENT '评分ID',
                                recipe_id INT NOT NULL COMMENT '被评分的菜谱ID',
                                rating FLOAT NOT NULL CHECK (rating >= 0 AND rating <= 5) COMMENT '评分（0-5分）',
                                review TEXT COMMENT '评分评价',
                                created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '评分时间'
);

```





## 食材管理
添加食材: 用户可以记录常用的食材，包括名称、数量、单位等。
编辑食材: 用户可以修改食材信息。
删除食材: 用户可以删除不再使用的食材。

## 菜谱分类
分类管理: 用户可以为菜谱添加分类（如：主菜、配菜、甜点等），方便查找和管理。
按分类浏览: 用户可以按分类浏览菜谱。
1. 收藏和评分
收藏菜谱: 用户可以收藏喜欢的菜谱，方便快速访问。
评分和评论: 用户可以对菜谱进行评分和评论，记录自己的烹饪体验。
1. 搜索和过滤
搜索功能: 用户可以通过菜名、食材等关键词搜索菜谱。
过滤功能: 用户可以根据分类、评分等条件过滤菜谱。
1. 购物清单
生成购物清单: 用户可以根据所选菜谱生成购物清单，列出所需的食材。
标记已购食材: 用户可以标记已购买的食材，方便管理。
1. 烹饪记录
记录烹饪过程: 用户可以记录每次烹饪的过程，包括时间、温度、注意事项等。
上传照片: 用户可以上传烹饪过程中的照片，记录成品效果。
1. 用户管理
用户注册和登录: 如果您希望支持多个用户，可以实现用户注册和登录功能。
个人资料管理: 用户可以管理自己的个人资料。
1. 数据备份和导出
数据备份: 用户可以备份自己的菜谱和食材数据。
导出功能: 用户可以将菜谱导出为 PDF 或其他格式，方便打印或分享。
1.  移动端支持
移动端适配: 如果您打算开发移动应用，可以考虑适配移动端界面，方便用户随时随地记录和查找菜谱。
1.  社交分享
分享功能: 用户可以将自己的菜谱分享至社交媒体，或与朋友分享。