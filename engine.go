package sqlo

import (
	"database/sql"
)

type db struct {
	engine *sql.DB
	s      string
	err    error
}

func NewDb(dns string) (d db, err error) {
	d = db{}
	open, err := sql.Open("mysql", dns)
	d.engine = open //连接数据库
	d.err = err
	return
}
