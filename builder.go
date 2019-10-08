package sqlo

import (
	"bytes"
	"fmt"
)

func (d db) Select(cols ...string) db {
	var buf bytes.Buffer
	buf.WriteString(" select ")
	count := 0
	for _, v := range cols {
		if len(cols)-1 != count {
			buf.WriteString(v)
			buf.WriteString(",")
		} else {
			buf.WriteString(v)
		}
	}
	d.s = buf.String()
	return d
}

func (d db) From(db string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)

	buf.WriteString(" from ")
	buf.WriteString(db)

	d.s = buf.String()
	return d
}

func (d db) Where() db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	buf.WriteString(" where ")
	d.s = buf.String()
	return d
}
func (d db) And(fields ...string) db {
	return andOr(d, " and ", fields...)
}

func (d db) Or(fields ...string) db {
	return andOr(d, "or", fields...)
}

func (d db) OrderBy(desc string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	buf.WriteString(" order by ")
	buf.WriteString(desc)
	d.s = buf.String()
	return d
}
func (d db) Limit(ps, pn int) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	buf.WriteString(fmt.Sprintf(" limit %d, %d", ps, pn))
	d.s = buf.String()
	return d
}

func (d db) Count(cols string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	buf.WriteString(fmt.Sprintf(" count(%s) AS %s ", cols, string(cols[0])))
	d.s = buf.String()
	return d
}
func (d db) string() string {
	return d.s
}
