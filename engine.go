package sqlo

import (
	"database/sql"
)

type Engine struct {
	db  *sql.DB
	s   string
	err error
}

func Connect(dns string) (d Engine, err error) {
	d = Engine{}
	open, err := sql.Open("mysql", dns)
	d.db = open //连接数据库
	d.err = err
	return
}

func (e Engine) Ping() error {
	return e.db.Ping()
}

func (e *Engine) Close() error {
	return e.db.Close()
}

func (e *Engine) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return e.db.Query(query, args)
}

func (e *Engine) DB() *sql.DB {
	return e.db
}
