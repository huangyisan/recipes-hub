// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package types

type Recipe struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Instructions   string `json:"instructions"`
	CookingTime    int64  `json:"cooking_time"`
	Difficulty     string `json:"difficulty"`
	RecipeType     string `json:"recipe_type"`
	ImageContent   string `json:"image_content"`
	CreationDate   string `json:"creation_date"`
	LastUpdated    string `json:"last_updated"`
	LastCookedDate string `json:"last_cooked_date"`
}

type RecipeCreateReq struct {
	Name         string `form:"name"`
	Instructions string `form:"instructions,optional"`
	CookingTime  int64  `form:"cooking_time,optional"`
	Difficulty   string `form:"difficulty,optional"`
	RecipeType   string `form:"recipe_type,optional"`
	ImageContent string `form:"image_content,optional"`
	CreationDate string `form:"creation_date,optional"`
}

type RecipeCreateResp struct {
	Recipe Recipe `json:"data"`
}
