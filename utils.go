package sqlo

import "bytes"

func andOr(d db, op string, fields ...string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	count := 0
	for _, v := range fields {
		if len(fields)-1 != count {
			buf.WriteString(v)
			buf.WriteString("= ? ")
			buf.WriteString(op)

		} else {
			buf.WriteString(v)
			buf.WriteString("= ? ")
		}
		count += 1
	}
	d.s = buf.String()
	return d
}
