package sqlBuild

//func Select(table string) (*SqlBuilder) {
//	auto := new(SqlBuilder)
//	auto.Select(table)
//	return auto
//}

func Insert(table string) (*SqlBuilder) {
	auto := new(SqlBuilder)
	auto.Insert(table)
	return auto
}
func Update(table string) (*SqlBuilder) {
	auto := new(SqlBuilder)
	auto.Update(table)
	return auto
}
func Delete(table string) (*SqlBuilder) {
	auto := new(SqlBuilder)
	auto.Delete(table)
	return auto
}



