package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RecipeIngredientsModel = (*customRecipeIngredientsModel)(nil)

type (
	// RecipeIngredientsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecipeIngredientsModel.
	RecipeIngredientsModel interface {
		recipeIngredientsModel
		withSession(session sqlx.Session) RecipeIngredientsModel
	}

	customRecipeIngredientsModel struct {
		*defaultRecipeIngredientsModel
	}
)

// NewRecipeIngredientsModel returns a model for the database table.
func NewRecipeIngredientsModel(conn sqlx.SqlConn) RecipeIngredientsModel {
	return &customRecipeIngredientsModel{
		defaultRecipeIngredientsModel: newRecipeIngredientsModel(conn),
	}
}

func (m *customRecipeIngredientsModel) withSession(session sqlx.Session) RecipeIngredientsModel {
	return NewRecipeIngredientsModel(sqlx.NewSqlConnFromSession(session))
}
