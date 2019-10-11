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
func (d db) Delete(table string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{"delete from ", table})

	//buf.WriteString("delete from ")
	//buf.WriteString(table)
	d.s = buf.String()
	return d
}

/**
UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
*/
func (d db) Update(table string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{"update ", table, " set "})

	//buf.WriteString("update ")
	//buf.WriteString(table)
	//buf.WriteString(" set ")
	d.s = buf.String()
	return d
}

/**
INSERT INTO table_name ( field1, field2,...fieldN )
                       VALUES
                       ( value1, value2,...valueN );
*/

func (d db) Insert(table string) db {
	var buf bytes.Buffer
	buf = Join(buf, []string{"insert into  ", table})
	//
	//buf.WriteString("insert into ")
	//buf.WriteString(table)
	d.s = buf.String()
	return d
}
