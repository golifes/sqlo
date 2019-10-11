package sqlo

import "bytes"

func andOr(d db, op string, fields ...string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	for _, v := range fields {
		buf.WriteString(op)
		buf.WriteString(v)
		buf.WriteString("? ")
	}
	d.s = buf.String()
	return d
}

func RangeS(buf bytes.Buffer, op string, fields ...string) bytes.Buffer {

	for index, v := range fields {
		buf.WriteString(v)
		buf.WriteString(op)
		if index != len(fields)-1 {
			buf.WriteString(",")
		}
	}
	return buf
}

/**
var buf bytes.Buffer
buf.WriteString(d.s)
buf.WriteString(col)
buf.WriteString(" like ? ")
*/

func Join(buf bytes.Buffer, cols []string) bytes.Buffer {
	for _, v := range cols {
		buf.WriteString(v)
	}
	return buf
}
