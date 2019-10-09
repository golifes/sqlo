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
func (d db) Delete(db string) db {
	var buf bytes.Buffer

	buf.WriteString("delete from ")
	buf.WriteString(db)
	d.s = buf.String()
	return d
}

/**
UPDATE table_name SET field1=new-value1, field2=new-value2 [WHERE Clause]
*/
func (d db) Update(db string) db {
	var buf bytes.Buffer
	buf.WriteString("update ")
	buf.WriteString(db)
	buf.WriteString(" set ")
	d.s = buf.String()
	return d
}

/**
INSERT INTO table_name ( field1, field2,...fieldN )
                       VALUES
                       ( value1, value2,...valueN );
*/

func (d db) Insert(db string) db {
	var buf bytes.Buffer
	buf.WriteString("insert into ")
	buf.WriteString(db)
	d.s = buf.String()
	return d
}
