package sqlo

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

/**
计算日期函数 https://www.runoob.com/mysql/mysql-functions.html
*/
//	计算起始日期 d 加上 n 天的日期
//	SELECT ADDDATE("2017-06-15", INTERVAL 10 DAY);
//->2017-06-25
func (e Engine) AddDate(time string, day int) Engine {
	var buf bytes.Buffer
	buf.WriteString("ADDDATE ( ")
	buf.WriteString(time)
	buf.WriteString(", INTERVAL  ")
	buf.WriteString(strconv.Itoa(day))
	buf.WriteString("DAY) ")
	e.s = buf.String()
	return e
}

//insert into table ( name,pwd ,id ) values( ?,?,?)
//insert into  wx (a,b) values (?,?)
func (e Engine) AddNowTime(cols []string) Engine {
	var buf bytes.Buffer
	key := ""
	for _, v := range cols {
		key += fmt.Sprintf(",%s", v)
	}
	e.s = strings.ReplaceAll(e.s, ") values", fmt.Sprintf("%s) values", key))

	value := ""

	for i := 0; i < len(cols); i++ {
		value += fmt.Sprintf(",%s", "now()")
	}

	e.s = strings.ReplaceAll(e.s, "?)", fmt.Sprintf("?%s)", value))
	buf.WriteString(e.s)
	e.s = buf.String()
	return e
}
