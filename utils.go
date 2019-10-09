package sqlo

import "bytes"

func andOr(d db, op string, fields ...string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	for _, v := range fields {
		buf.WriteString(op)
		buf.WriteString(v)
		buf.WriteString("=? ")
	}
	d.s = buf.String()
	return d
}
