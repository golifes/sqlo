package sqlo

import (
	"testing"
)

var engine Engine

func init() {
	engine, _ = Connect("root:123@tcp(127.0.0.1)/demo?charset=utf8&parseTime=True&loc=Local")

}

func TestCount(t *testing.T) {

	//select a,b from wx where name=? and age = ? order by id desc limit 0 ,10
	sql := engine.
		Select("a", "b").From("wx").Alias("w").
		Where("id").
		And("name", "age").
		OrderBy("id desc").Limit(0, 10).
		String()
	t.Log(sql)
	//两种执行sql的方式
	//sqlo封装的query
	engine.Query(sql, "aaa", "bbb")
	//database/sql原生的query
	engine.DB().Query(sql, "aaa", "bbb")
}

func TestDb_Delete(t *testing.T) {
	sql := engine.
		Delete("wx").
		Where("id").
		And("name").
		String()
	t.Log(sql)
}

func TestDb_Update(t *testing.T) {
	sql := engine.Update("wx").Fields("a", "b").Where("c").String()
	t.Log(sql)
}
func TestDb_Insert(t *testing.T) {
	sql := engine.Insert("wx").
		Cols("a", "b").
		String()
	t.Log(sql)
}

func TestDb_Like(t *testing.T) {
	sql := engine.Select("a", "b").From("wx").Where("id").Like("name").String()
	t.Log(sql)
}

// SELECT name, COUNT(*) FROM   employee_tbl GROUP BY name;
func TestDb_GroupBy(t *testing.T) {
	sql := engine.Select("id").Count("name").From("wx").GroupBy("name", "id").String()
	t.Log(sql)
}

/**
	SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a
INNER JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;
*/
func TestDb_InnerJoin(t *testing.T) {
	sql := engine.Select("a.runoob_id", "a.runoob_author", "b.runoob_count").
		From(" runoob_tbl a").
		InnerJoin("tcount_tbl b").On("a.runoob_author=").And("a=", "b=").String()
	t.Log(sql)
}

func TestEngine_Insert(t *testing.T) {
	sql := engine.Insert("wx").Cols("a", "b").AddNowTime([]string{"ctime", "mtime"}).String()
	t.Log(sql)
}
