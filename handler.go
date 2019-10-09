package sqlo

/**
select
*/
type Selecter interface {
	Select(cols ...string) db //select
	From(db string) db        //from db
	Where(col string) db
}

type Inserter interface {
	Insert(db string) db
	Cols(cols ...string) db
}
