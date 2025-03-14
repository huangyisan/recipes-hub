package svc

import (
	"context"
	"github.com/huangyisan/recipes-hub/ingredients/cmd/rpc/internal/config"
	"github.com/huangyisan/recipes-hub/ingredients/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config                 config.Config
	IngredientsModel       model.IngredientsModel
	RecipeIngredientsModel model.RecipeIngredientsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := NewMysql(c)
	return &ServiceContext{
		Config:                 c,
		IngredientsModel:       model.NewIngredientsModel(mysql),
		RecipeIngredientsModel: model.NewRecipeIngredientsModel(mysql),
	}
}

func NewMysql(c config.Config) sqlx.SqlConn {
	sqlConn := sqlx.NewMysql(c.MysqlConf.DataSource)
	db, err := sqlConn.RawDB()
	if err != nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.MysqlConf.Timeout)*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return sqlConn
}
