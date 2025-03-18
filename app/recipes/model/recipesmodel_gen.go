// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.8.1

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	recipesFieldNames          = builder.RawFieldNames(&Recipes{})
	recipesRows                = strings.Join(recipesFieldNames, ",")
	recipesRowsExpectAutoSet   = strings.Join(stringx.Remove(recipesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	recipesRowsWithPlaceHolder = strings.Join(stringx.Remove(recipesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	recipesModel interface {
		Insert(ctx context.Context, data *Recipes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Recipes, error)
		FindOneByRecipeName(ctx context.Context, recipeName string) (*Recipes, error)
		Update(ctx context.Context, data *Recipes) error
		SelectBuilder() squirrel.SelectBuilder
		FindAll(context.Context, squirrel.SelectBuilder, string) ([]*Recipes, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Recipes, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error

		Delete(ctx context.Context, id int64) error
	}

	defaultRecipesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Recipes struct {
		Id             int64          `db:"id"`               // 菜谱ID
		RecipeName     string         `db:"recipe_name"`      // 菜名，唯一约束
		Instructions   sql.NullString `db:"instructions"`     // 烹饪步骤
		CookingTime    int64          `db:"cooking_time"`     // 烹饪时间（以分钟为单位）
		Difficulty     string         `db:"difficulty"`       // 难度等级
		RecipeType     int64          `db:"recipe_type"`      // 菜谱类型
		ImageContent   sql.NullString `db:"image_content"`    // 图片的URL地址
		CreationDate   time.Time      `db:"creation_date"`    // 创建时间
		LastUpdated    time.Time      `db:"last_updated"`     // 更新时间
		LastCookedDate sql.NullTime   `db:"last_cooked_date"` // 最后一次烹饪时间
	}
)

func newRecipesModel(conn sqlx.SqlConn) *defaultRecipesModel {
	return &defaultRecipesModel{
		conn:  conn,
		table: "`recipes`",
	}
}

func (m *defaultRecipesModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRecipesModel) FindOne(ctx context.Context, id int64) (*Recipes, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", recipesRows, m.table)
	var resp Recipes
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRecipesModel) FindOneByRecipeName(ctx context.Context, recipeName string) (*Recipes, error) {
	var resp Recipes
	query := fmt.Sprintf("select %s from %s where `recipe_name` = ? limit 1", recipesRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, recipeName)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRecipesModel) Insert(ctx context.Context, data *Recipes) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, recipesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RecipeName, data.Instructions, data.CookingTime, data.Difficulty, data.RecipeType, data.ImageContent, data.CreationDate, data.LastUpdated, data.LastCookedDate)
	return ret, err
}

func (m *defaultRecipesModel) Update(ctx context.Context, newData *Recipes) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, recipesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.RecipeName, newData.Instructions, newData.CookingTime, newData.Difficulty, newData.RecipeType, newData.ImageContent, newData.CreationDate, newData.LastUpdated, newData.LastCookedDate, newData.Id)
	return err
}

func (m *defaultRecipesModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultRecipesModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Recipes, error) {

	builder = builder.Columns(recipesRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	// query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Recipes

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRecipesModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Recipes, error) {

	builder = builder.Columns(recipesRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	// query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	query, values, err := builder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Recipes

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRecipesModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, fn)
	//return m.conn.TransactCtx(ctx,func(ctx context.Context,session sqlx.Session) error {
	//        return  fn(ctx,session)
	//})

}

func (m *defaultRecipesModel) tableName() string {
	return m.table
}
