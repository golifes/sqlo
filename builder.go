package sqlo

import (
	"bytes"
	"fmt"
	"strings"
)

func (e Engine) And(fields ...string) Engine {
	return andOr(e, " and ", fields...)
}

func (e Engine) Or(fields ...string) Engine {
	return andOr(e, "or", fields...)
}

func (e Engine) OrderBy(desc string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, " order by ", desc})

	//buf.WriteString(e.s)
	//buf.WriteString("order by ")
	//buf.WriteString(desc)
	e.s = buf.String()
	return e
}
func (e Engine) Limit(ps, pn int) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, fmt.Sprintf(" limit %d, %d", ps, pn)})
	e.s = buf.String()
	return e
}

func (e Engine) Count(cols string) Engine {
	var buf bytes.Buffer
	buf.WriteString(e.s)
	s := strings.TrimSpace(e.s)
	fmt.Println(s, len(s))
	if len(s) > 6 {
		buf.WriteString(fmt.Sprintf(",count(%s) AS %s ", cols, string(cols[0])))
	} else {
		buf.WriteString(fmt.Sprintf("count(%s) AS %s ", cols, string(cols[0])))
	}
	e.s = buf.String()
	return e
}

/**
SELECT * from runoob_tbl  WHERE runoob_author LIKE '%COM';
("select * from t_ally where ally_name like ?", "%" + allyName + "%")
*/
func (e Engine) Like(col string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, " and ", col, " like ? "})

	//buf.WriteString(e.s)
	//buf.WriteString("and ")
	//buf.WriteString(col)
	//buf.WriteString(" like ? ")
	e.s = buf.String()
	return e
}

func (e Engine) Union() Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, " union "})
	//
	//buf.WriteString(e.s)
	//buf.WriteString(" union ")
	e.s = buf.String()
	return e
}

func (e Engine) Alias(alias string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, " as ", alias})
	//
	//
	//buf.WriteString(e.s)
	//buf.WriteString(" as ")
	//buf.WriteString(alias)
	e.s = buf.String()
	return e
}

func (e Engine) GroupBy(cols ...string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, " group by "})
	buf = RangeS(buf, "", cols...)
	e.s = buf.String()
	return e
}

/**
SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a
INNER JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;
*/
func (e Engine) InnerJoin(table string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, " inner join ", table, " on "})
	e.s = buf.String()
	return e
}

func (e Engine) LeftJoin(table string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, " left join ", table, " on "})
	e.s = buf.String()
	return e
}

func (e Engine) RightJoin(table string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, " right join ", table, " on "})
	e.s = buf.String()
	return e
}

/**
v1 >= v2
v2 <= v1
条件前面加操作符号
*/
func (e Engine) On(col string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{e.s, col, "?"})
	e.s = buf.String()
	return e
}

func (e Engine) string() string {
	return e.s
}
