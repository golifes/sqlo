package sqlo

import (
	"bytes"
	"fmt"
	"strings"
)

//select
func (e Engine) Select(cols ...string) Engine {
	var buf bytes.Buffer
	buf.WriteString(" select ")
	buf = RangeS(buf, "", cols...)
	e.s = buf.String()
	return e
}

// from
func (e Engine) From(table string) Engine {
	var buf bytes.Buffer

	buf = Join(buf, []string{e.s, " from ", table})

	//
	//buf.WriteString(e.s)
	//
	//buf.WriteString(" from ")
	//buf.WriteString(table)

	e.s = buf.String()
	return e
}

//where
func (e Engine) Where(col string) Engine {
	var buf bytes.Buffer

	buf = Join(buf, []string{e.s, " where "})

	//
	//buf.WriteString(e.s)
	//buf.WriteString(" where ")

	col = strings.TrimSpace(col)
	if col != "" {
		buf.WriteString(col)
	}
	e.s = buf.String()
	return e
}

/**
update 字段
*/
func (e Engine) Fields(fields ...string) Engine {
	var buf bytes.Buffer
	buf.WriteString(e.s)
	buf = RangeS(buf, " =? ", fields...)

	e.s = buf.String()
	return e
}

/**
insert 时候的table 字段
*/
func (e Engine) Cols(cols ...string) Engine {
	var buf bytes.Buffer
	buf.WriteString(e.s)
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
	e.s = buf.String()
	return e
}
