package sqlo

/**
select
*/
type Selecter interface {
	Select(cols ...string) db //select
	From(table string) db     //from db
	Where(col string) db
}

type Inserter interface {
	Insert(table string) db
	Cols(cols ...string) db
}
