package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CuisineModel = (*customCuisineModel)(nil)

type (
	// CuisineModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCuisineModel.
	CuisineModel interface {
		cuisineModel
		withSession(session sqlx.Session) CuisineModel
	}

	customCuisineModel struct {
		*defaultCuisineModel
	}
)

// NewCuisineModel returns a model for the database table.
func NewCuisineModel(conn sqlx.SqlConn) CuisineModel {
	return &customCuisineModel{
		defaultCuisineModel: newCuisineModel(conn),
	}
}

func (m *customCuisineModel) withSession(session sqlx.Session) CuisineModel {
	return NewCuisineModel(sqlx.NewSqlConnFromSession(session))
}
