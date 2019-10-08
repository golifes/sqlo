package sqlo

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	db, _ := NewDb("root:123@tcp(127.0.0.1)/demo?charset=utf8&parseTime=True&loc=Local")
	i := db.
		Select("a", "b").From("wx").
		Where().
		And("name", "age").
		string()
	fmt.Println(i)

}
