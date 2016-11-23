package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"

	"github.com/g0dgarden/ddd-api/conf"
)

// Executor はrepositoryの引数として使用するインタフェース
// Transaction操作をさせないように明示的にinterfaceを持たせていません
type Executor interface {
	Get(i interface{}, keys ...interface{}) (interface{}, error)
	Select(i interface{}, sql string, args ...interface{}) ([]interface{}, error)
	SelectOne(holder interface{}, query string, args ...interface{}) error
	Insert(list ...interface{}) error
	Update(list ...interface{}) (int64, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// Connector はDBのコネクションとして使用するインタフェース
type Connector interface {
	Executor
	Begin() error
	Commit() error
	Rollback() error
}

type connection struct {
	dbMap       *gorp.DbMap
	transaction *gorp.Transaction
	exec        Executor
}

// NewConnection はDBのコネクションとして使用するインタフェース
func NewConnection(dbConf *conf.SectionDB) Connector {
	dbMap, err := newDbMap(dbConf)
	if err != nil {
		panic(err)
	}
	return &connection{dbMap: dbMap, exec: dbMap}
}

// DbMapの生成
func newDbMap(dbConf *conf.SectionDB) (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", getConnectionString(dbConf))
	if err != nil {
		return nil, err
	}
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	associateTable(dbMap)
	return dbMap, nil
}

// 構造体とテーブルの関連付けを行う
func associateTable(dbMap *gorp.DbMap) {
	// TODO: 循環インポートとなってしまうため別の層でマッピングする
	// dbMap.AddTableWithName(users.User{}, "users").SetKeys(true, "Id")
}

// コネクションの接続文字列の取得
func getConnectionString(dbConf *conf.SectionDB) string {
	const sourceFormat string = "%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo"
	return fmt.Sprintf(sourceFormat, dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Database)
}

func (c *connection) Get(i interface{}, keys ...interface{}) (interface{}, error) {
	return c.exec.Get(i, keys...)
}

func (c *connection) Select(i interface{}, sql string, args ...interface{}) ([]interface{}, error) {
	return c.exec.Select(i, sql, args...)
}

func (c *connection) SelectOne(holder interface{}, query string, args ...interface{}) error {
	return c.exec.SelectOne(holder, query, args...)
}

func (c *connection) Insert(list ...interface{}) error {
	return c.exec.Insert(list...)
}

func (c *connection) Update(list ...interface{}) (int64, error) {
	return c.exec.Update(list...)
}

func (c *connection) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.exec.Exec(query, args...)
}

func (c *connection) Begin() error {
	t, err := c.dbMap.Begin()
	if err != nil {
		return err
	}
	c.transaction = t
	c.exec = t
	return nil
}

func (c *connection) Commit() error {
	err := c.transaction.Commit()
	c.transaction = nil
	c.exec = c.dbMap
	return err
}

func (c *connection) Rollback() error {
	err := c.transaction.Rollback()
	c.transaction = nil
	c.exec = c.dbMap
	return err
}
