package sqlBuild

import (
	"testing"
	"github.com/golyu/sql-build"
)

func TestInsert(t *testing.T) {
	//sql := sqlBuild.Insert("tabName").
	//	Value(0, "uid").
	//	Value("", "site_id")
	//t.Log(sql.String())
	//sql2 := sqlBuild.Insert("tabName").
	//	Value_(0, "uid").
	//	Value_("", "site_id")
	//t.Log(sql2.String())

}
func TestUpdate(t *testing.T) {
	sql := sqlBuild.Update("l_agent").
		Set("ssid", "ssid=").
		Set("logintimeeeee", "login_time=").
		Where("ooo", "site_id=").
		Where("username", "username=")
	t.Log(sql.String())
}
func TestSelect(t *testing.T) {
	b:=""
	sql := sqlBuild.Select("l_user_log").
		Where("xx", "site_id=").
		Column("count(1) as num").
		Where_("infe", "index_id").
		Where_("tt", "at_id").
		Where_(8, "ua_id").
		Where_(0, "sh_id").
		Where_(4, "username").
		Where_(8, "add_time>=").
		Where_(1, "add_time<=").
		Where_(1, "ip").
		Like(b+"%","username").
		In([]int{12,232,44},"ppp").
		Limit(int(6)).
		OrderBy("-add_time").
		Offset(int(6 * (6 - 1)))
	t.Log(sql.String())
}

//func TestCheckInjection(t *testing.T) {
//	sqlValue:=[]string{
//		"select * from",
//		"update * from",
//		"xxxx * from",
//		"insert into from",
//		"xxxx delete b",
//	}
//	for _, v := range sqlValue {
//		if sqlBuild.CheckInjection(v) {
//			t.Log("有注入:",v)
//		}
//	}
//
//}

