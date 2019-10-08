package sqlo

import (
	"bytes"
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
	var buf bytes.Buffer
	buf.WriteString(d.s)
	count := 0
	for _, v := range fields {
		if len(fields)-1 != count {
			buf.WriteString(v)
			buf.WriteString(" = ?  and ")
		} else {
			buf.WriteString(v)
			buf.WriteString(" = ?  ")
		}
		count += 1
	}
	d.s = buf.String()
	return d
}

func (d db) string() string {
	return d.s
}
