package sqlo

import (
	"bytes"
	"fmt"
	"strings"
)

//select
func (d db) Select(cols ...string) db {
	var buf bytes.Buffer
	buf.WriteString(" select ")
	buf = RangeS(buf, "", cols...)
	d.s = buf.String()
	return d
}

// from
func (d db) From(table string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)

	buf.WriteString(" from ")
	buf.WriteString(table)

	d.s = buf.String()
	return d
}

//where
func (d db) Where(col string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	buf.WriteString(" where ")

	col = strings.TrimSpace(col)
	if col != "" {
		buf.WriteString(col)
		buf.WriteString("=? ")
	}
	d.s = buf.String()
	return d
}

/**
update 字段
*/
func (d db) Fields(fields ...string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	buf = RangeS(buf, " =? ", fields...)

	d.s = buf.String()
	return d
}

/**
insert 时候的table 字段
*/
func (d db) Cols(cols ...string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	buf.WriteString(" (")
	place := ""
	for index, v := range cols {
		if len(cols)-1 != index {
			place += fmt.Sprintf("%s%s", "?", ",")
			buf.WriteString(v)
			buf.WriteString(",")
		} else {
			place += fmt.Sprintf("%s", "?")
			buf.WriteString(v)
		}
	}
	buf.WriteString(") values (")
	buf.WriteString(place)
	buf.WriteString(")")
	d.s = buf.String()
	return d
}
