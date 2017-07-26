package sqlBuild

import (
	"testing"
	"github.com/golyu/sql-build"
	"github.com/golyu/sql-build/debug"
)

type Tab struct {
	Id   int `insert:"id;auto"`
	Name string `insert:"name"`
	Age  int `insert:"age"`
}

func TestValue(t *testing.T) {
	var tab = Tab{Id: 0, Name: "yiersan", Age: 18}
	sql, err := sqlBuild.Insert("xx").
		Value(&tab).String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}

func TestValues(t *testing.T) {
	var tab1 = Tab{Id: 0, Name: "yiersan", Age: 18}
	var tab2 = Tab{Id: 0, Name: "xx", Age: 16}
	var tab3 = Tab{Id: 0, Name: "pp", Age: 18}
	var tabs = []Tab{tab1, tab2, tab3}
	sql, err := sqlBuild.Insert("xx").
		Values(tabs).String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}

func TestOrUpdate(t *testing.T) {
	debug.Debug = true
	var tab1 = Tab{Id: 0, Name: "yiersan", Age: 18}
	var tab2 = Tab{Id: 0, Name: "xx", Age: 16}
	var tab3 = Tab{Id: 0, Name: "pp", Age: 18}
	var tabs = []Tab{tab1, tab2, tab3}
	sql, err := sqlBuild.Insert("xx").
		Values(tabs).OrUpdate().String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}