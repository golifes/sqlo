package sqlo

/**
select
*/
type Selecter interface {
	Select(cols ...string) Engine //select
	From(table string) Engine     //from Engine
	Where(col string) Engine
}

type Inserter interface {
	Insert(table string) Engine
	Cols(cols ...string) Engine
}
