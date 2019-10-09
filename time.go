package sqlo

import (
	"bytes"
	"strconv"
)

/**
计算日期函数 https://www.runoob.com/mysql/mysql-functions.html
*/
//	计算起始日期 d 加上 n 天的日期
//	SELECT ADDDATE("2017-06-15", INTERVAL 10 DAY);
//->2017-06-25
func (d db) AddDate(time string, day int) db {
	var buf bytes.Buffer
	buf.WriteString("ADDDATE ( ")
	buf.WriteString(time)
	buf.WriteString(", INTERVAL  ")
	buf.WriteString(strconv.Itoa(day))
	buf.WriteString("DAY) ")
	d.s = buf.String()
	return d
}
