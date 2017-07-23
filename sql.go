package sqlBuild

type SelectInf interface {
	Select(table string) SelectInf
	Column(column string) SelectInf
	Where(value interface{}, key string, rules ... Rule) SelectInf
	Where_(value interface{}, key string, rules ... Rule) SelectInf
	Like(value string, key string) SelectInf
	In(values interface{}, key string) SelectInf
	NotIn(values interface{}, key string) SelectInf
	OrderBy(orderBy string) SelectInf
	Limit(limit int) SelectInf
	Offset(offset int) SelectInf
	GroupBy(groupBy string) SelectInf
	String() (string, error)
}

type InsertInf interface {
	Insert(table string) InsertInf
	Value(value interface{}, key string) InsertInf
	Value_(value interface{}, key string) InsertInf
}

type UpdateInf interface {
	Update(table string) UpdateInf
	Column(column string) UpdateInf
	Where(value interface{}, key string) UpdateInf
	Where_(value interface{}, key string) UpdateInf
	Like(value string, key string) UpdateInf
	In(values interface{}, key string) UpdateInf
	NotIn(values interface{}, key string) UpdateInf
	OrderBy(orderBy string) UpdateInf
	Limit(limit int) UpdateInf
	Offset(offset int) UpdateInf
	GroupBy(groupBy string) UpdateInf
	String() (string, error)
}

type DeleteInf interface {
	Delete(table string) DeleteInf
	Column(column string) DeleteInf
	Where(value interface{}, key string) DeleteInf
	Where_(value interface{}, key string) DeleteInf
	Like(value string, key string) DeleteInf
	In(values interface{}, key string) DeleteInf
	NotIn(values interface{}, key string) DeleteInf
	OrderBy(orderBy string) DeleteInf
	Limit(limit int) DeleteInf
	Offset(offset int) DeleteInf
	GroupBy(groupBy string) DeleteInf
	String() (string, error)
}
