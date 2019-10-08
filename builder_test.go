package sqlo

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	engine, _ := Connect("root:123@tcp(127.0.0.1)/demo?charset=utf8&parseTime=True&loc=Local")
	//select a,b from wx where name=? and age = ? order by id desc limit 0 ,10
	i := engine.
		Select("a", "b").Count("c").From("wx").
		Where().
		And("name", "age").
		OrderBy("id desc").Limit(0, 10).
		string()
	fmt.Println(i)

}
