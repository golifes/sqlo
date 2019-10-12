package sqlo

import (
	"bytes"
)

/**
insert update delete
*/

/**
DELETE FROM table_name [WHERE Clause]
*/
func (e Engine) Delete(table string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{"delete from ", table})

	//buf.WriteString("delete from ")
	//buf.WriteString(table)
	e.s = buf.String()
	return e
}

/**
UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
*/
func (e Engine) Update(table string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{"update ", table, " set "})

	//buf.WriteString("update ")
	//buf.WriteString(table)
	//buf.WriteString(" set ")
	e.s = buf.String()
	return e
}

/**
INSERT INTO table_name ( field1, field2,...fieldN )
                       VALUES
                       ( value1, value2,...valueN );
*/

func (e Engine) Insert(table string) Engine {
	var buf bytes.Buffer
	buf = Join(buf, []string{"insert into  ", table})
	//
	//buf.WriteString("insert into ")
	//buf.WriteString(table)
	e.s = buf.String()
	return e
}
