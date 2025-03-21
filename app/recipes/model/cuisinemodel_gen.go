// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.8.1

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cuisineFieldNames          = builder.RawFieldNames(&Cuisine{})
	cuisineRows                = strings.Join(cuisineFieldNames, ",")
	cuisineRowsExpectAutoSet   = strings.Join(stringx.Remove(cuisineFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cuisineRowsWithPlaceHolder = strings.Join(stringx.Remove(cuisineFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cuisineModel interface {
		Insert(ctx context.Context, data *Cuisine) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Cuisine, error)
		FindOneByCuisineName(ctx context.Context, cuisineName string) (*Cuisine, error)
		Update(ctx context.Context, data *Cuisine) error
		SelectBuilder() squirrel.SelectBuilder
		FindAll(context.Context, squirrel.SelectBuilder, string) ([]*Cuisine, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Cuisine, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error

		Delete(ctx context.Context, id int64) error
	}

	defaultCuisineModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Cuisine struct {
		Id          int64          `db:"id"`           // 菜系ID
		CuisineName string         `db:"cuisine_name"` // 菜系名称
		Description sql.NullString `db:"description"`  // 菜系描述
	}
)

func newCuisineModel(conn sqlx.SqlConn) *defaultCuisineModel {
	return &defaultCuisineModel{
		conn:  conn,
		table: "`cuisine`",
	}
}

func (m *defaultCuisineModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCuisineModel) FindOne(ctx context.Context, id int64) (*Cuisine, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cuisineRows, m.table)
	var resp Cuisine
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

func (m *defaultCuisineModel) FindOneByCuisineName(ctx context.Context, cuisineName string) (*Cuisine, error) {
	var resp Cuisine
	query := fmt.Sprintf("select %s from %s where `cuisine_name` = ? limit 1", cuisineRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, cuisineName)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCuisineModel) Insert(ctx context.Context, data *Cuisine) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, cuisineRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CuisineName, data.Description)
	return ret, err
}

func (m *defaultCuisineModel) Update(ctx context.Context, newData *Cuisine) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cuisineRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.CuisineName, newData.Description, newData.Id)
	return err
}

func (m *defaultCuisineModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultCuisineModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Cuisine, error) {

	builder = builder.Columns(cuisineRows)

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

	var resp []*Cuisine

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultCuisineModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Cuisine, error) {

	builder = builder.Columns(cuisineRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	// query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	query, values, err := builder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Cuisine

	err = m.conn.QueryRowsCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultCuisineModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.conn.TransactCtx(ctx, fn)
	//return m.conn.TransactCtx(ctx,func(ctx context.Context,session sqlx.Session) error {
	//        return  fn(ctx,session)
	//})

}

func (m *defaultCuisineModel) tableName() string {
	return m.table
}
