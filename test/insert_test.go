package sqlBuild

import (
	"github.com/xiaohei16h/sql-build"
	"github.com/xiaohei16h/sql-build/debug"
	"testing"
	"time"
)

type Tab struct {
	Id   int       `insert:"id;auto;mycat:next value for MYCATSEQ_AGENT"`
	Name string    `insert:"name"`
	Age  int       `insert:"age"`
	Json string    `json:"json"`
	Time time.Time `json:"time"`
}

func TestValue(t *testing.T) {
	var tab = Tab{Id: 0, Name: "Drop ", Age: 18}
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
	var tab2 = Tab{Id: 0, Name: "xx's", Age: 16}
	var tab3 = Tab{Id: 0, Name: "pp", Age: 18, Json: "jj", Time: time.Now()}
	var tabs = []Tab{tab1, tab2, tab3}
	sql, err := sqlBuild.Insert("xx").
		Values(tabs).OrUpdate().NoOnDuplicateKeyUpdateOption("id").String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
func TestOption(t *testing.T) {
	debug.Debug = true
	var tab1 = Tab{Id: 0, Name: "yiersan", Age: 18}
	sql, err := sqlBuild.Insert("xx").
		Option("name").
		Value(&tab1).String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
func TestNoOption(t *testing.T) {
	debug.Debug = true
	var tab1 = Tab{Id: 0, Name: "yiersan", Age: 18}
	var tab2 = Tab{Id: 0, Name: "xx", Age: 16}
	var tab3 = Tab{Id: 0, Name: "pp", Age: 18}
	var tabs = []Tab{tab1, tab2, tab3}
	sql, err := sqlBuild.Insert("xx").
		NoOption("name").
		Values(tabs).OrUpdate().String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
func TestMycat(t *testing.T) {
	debug.Debug = true
	var tab = Tab{Id: 0, Name: "yiersan", Age: 18}
	sql, err := sqlBuild.Insert("xx").
		Value(&tab).
		String()
	if err != nil {
		t.Error(err)
	}
	t.Log(sql)
}
