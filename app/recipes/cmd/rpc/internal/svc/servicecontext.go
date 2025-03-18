package svc

import (
	"context"
	"github.com/huangyisan/recipes-hub/app/recipes/cmd/rpc/internal/config"
	"github.com/huangyisan/recipes-hub/app/recipes/model"
	"github.com/huangyisan/recipes-hub/pkg/cloudflare/r2"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config       config.Config
	RecipesModel model.RecipesModel
	S3Handler    r2.S3Handler
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := NewMysql(c)
	// rpc超时1分钟
	zrpc.SetClientSlowThreshold(time.Second * 60 * 1)
	return &ServiceContext{
		Config:       c,
		RecipesModel: model.NewRecipesModel(mysql),
		S3Handler:    NewS3(),
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

func NewS3() r2.S3Handler {
	return r2.NewR2Server()
}
