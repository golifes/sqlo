package sqlo

import (
	"testing"
)

var engine db

func init() {
	engine, _ = Connect("root:123@tcp(127.0.0.1)/demo?charset=utf8&parseTime=True&loc=Local")

}

func TestCount(t *testing.T) {

	//select a,b from wx where name=? and age = ? order by id desc limit 0 ,10
	sql := engine.
		Select("a", "b").Count("c").From("wx").Alias("w").
		Where("id").
		And("name", "age").
		OrderBy("id desc").Limit(0, 10).
		string()
	t.Log(sql)
}

func TestDb_Delete(t *testing.T) {
	sql := engine.
		Delete("wx").
		Where("id").
		And("name").
		string()
	t.Log(sql)
}

func TestDb_Update(t *testing.T) {
	sql := engine.Update("wx").Fields("a", "b").Where("c").string()
	t.Log(sql)
}
func TestDb_Insert(t *testing.T) {
	sql := engine.Insert("wx").
		Cols("a", "b").
		string()
	t.Log(sql)
}

func TestDb_Like(t *testing.T) {
	sql := engine.Select("a", "b").From("wx").Where("id").Like("name").string()
	t.Log(sql)
}
