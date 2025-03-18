package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RecipesModel = (*customRecipesModel)(nil)

type (
	// RecipesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecipesModel.
	RecipesModel interface {
		recipesModel
		withSession(session sqlx.Session) RecipesModel
	}

	customRecipesModel struct {
		*defaultRecipesModel
	}
)

// NewRecipesModel returns a model for the database table.
func NewRecipesModel(conn sqlx.SqlConn) RecipesModel {
	return &customRecipesModel{
		defaultRecipesModel: newRecipesModel(conn),
	}
}

func (m *customRecipesModel) withSession(session sqlx.Session) RecipesModel {
	return NewRecipesModel(sqlx.NewSqlConnFromSession(session))
}
