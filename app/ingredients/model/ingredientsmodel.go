package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ IngredientsModel = (*customIngredientsModel)(nil)

type (
	// IngredientsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIngredientsModel.
	IngredientsModel interface {
		ingredientsModel
		withSession(session sqlx.Session) IngredientsModel
	}

	customIngredientsModel struct {
		*defaultIngredientsModel
	}
)

// NewIngredientsModel returns a model for the database table.
func NewIngredientsModel(conn sqlx.SqlConn) IngredientsModel {
	return &customIngredientsModel{
		defaultIngredientsModel: newIngredientsModel(conn),
	}
}

func (m *customIngredientsModel) withSession(session sqlx.Session) IngredientsModel {
	return NewIngredientsModel(sqlx.NewSqlConnFromSession(session))
}
