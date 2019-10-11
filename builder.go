package sqlo

import (
	"bytes"
	"fmt"
	"strings"
)

func (d db) And(fields ...string) db {
	return andOr(d, " and ", fields...)
}

func (d db) Or(fields ...string) db {
	return andOr(d, "or", fields...)
}

func (d db) OrderBy(desc string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, " order by ", desc})

	//buf.WriteString(d.s)
	//buf.WriteString("order by ")
	//buf.WriteString(desc)
	d.s = buf.String()
	return d
}
func (d db) Limit(ps, pn int) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, fmt.Sprintf(" limit %d, %d", ps, pn)})
	d.s = buf.String()
	return d
}

func (d db) Count(cols string) db {
	var buf bytes.Buffer
	buf.WriteString(d.s)
	s := strings.TrimSpace(d.s)
	fmt.Println(s, len(s))
	if len(s) > 6 {
		buf.WriteString(fmt.Sprintf(",count(%s) AS %s ", cols, string(cols[0])))
	} else {
		buf.WriteString(fmt.Sprintf("count(%s) AS %s ", cols, string(cols[0])))
	}
	d.s = buf.String()
	return d
}

/**
SELECT * from runoob_tbl  WHERE runoob_author LIKE '%COM';
("select * from t_ally where ally_name like ?", "%" + allyName + "%")
*/
func (d db) Like(col string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, " and ", col, " like ? "})

	//buf.WriteString(d.s)
	//buf.WriteString("and ")
	//buf.WriteString(col)
	//buf.WriteString(" like ? ")
	d.s = buf.String()
	return d
}

func (d db) Union() db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, " union "})
	//
	//buf.WriteString(d.s)
	//buf.WriteString(" union ")
	d.s = buf.String()
	return d
}

func (d db) Alias(alias string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, " as ", alias})
	//
	//
	//buf.WriteString(d.s)
	//buf.WriteString(" as ")
	//buf.WriteString(alias)
	d.s = buf.String()
	return d
}

func (d db) GroupBy(cols ...string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, " group by "})
	buf = RangeS(buf, "", cols...)
	d.s = buf.String()
	return d
}

/**
SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a
INNER JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;
*/
func (d db) InnerJoin(table string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, " inner join ", table, " on "})
	d.s = buf.String()
	return d
}

func (d db) LeftJoin(table string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, " left join ", table, " on "})
	d.s = buf.String()
	return d
}

func (d db) RightJoin(table string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, " right join ", table, " on "})
	d.s = buf.String()
	return d
}

/**
v1 >= v2
v2 <= v1
条件前面加操作符号
*/
func (d db) On(col string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{d.s, col, "?"})
	d.s = buf.String()
	return d
}

func (d db) string() string {
	return d.s
}
