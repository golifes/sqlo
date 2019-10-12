package sqlo

import "bytes"

func andOr(e Engine, op string, fields ...string) Engine {
	var buf bytes.Buffer
	buf.WriteString(e.s)
	for _, v := range fields {
		buf.WriteString(op)
		buf.WriteString(v)
		buf.WriteString("? ")
	}
	e.s = buf.String()
	return e
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
buf.WriteString(e.s)
buf.WriteString(col)
buf.WriteString(" like ? ")
*/

func Join(buf bytes.Buffer, cols []string) bytes.Buffer {
	for _, v := range cols {
		buf.WriteString(v)
	}
	return buf
}
