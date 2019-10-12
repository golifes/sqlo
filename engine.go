package sqlo

import (
	"database/sql"
)

type Engine struct {
	engine *sql.DB
	s      string
	err    error
}

func Connect(dns string) (d Engine, err error) {
	d = Engine{}
	open, err := sql.Open("mysql", dns)
	d.engine = open //连接数据库
	d.err = err
	return
}
