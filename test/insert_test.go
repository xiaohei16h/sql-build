package sqlBuild

import (
	"fmt"
	"github.com/xiaohei16h/sql-build"
	"github.com/xiaohei16h/sql-build/debug"
	"strings"
	"testing"
	"time"
)

type JsonTime time.Time

const ctLayout = "2006-01-02 15:04:05"

func Now() JsonTime {
	return JsonTime(time.Now())
}

func (ct JsonTime) TrimQuote(b []byte) bool {
	return true
}

// UnmarshalJSON Parses the json string in the custom format
func (ct *JsonTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	*ct = JsonTime(nt)
	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *JsonTime) String() string {
	t := time.Time(*ct)
	// 如果时间为0，则默认为空字符串，默认为0001-01-01 00:00:00
	if t.IsZero() {
		return "\"\""
	}

	return fmt.Sprintf("%q", t.Format(ctLayout))
}

type Tab struct {
	Id   int       `insert:"id;auto;mycat:next value for MYCATSEQ_AGENT"`
	Name string    `insert:"name"`
	Age  int       `insert:"age"`
	Json string    `json:"json"`
	Time time.Time `json:"time"`
	StartTime JsonTime `json:"start_time"`
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
	var tab3 = Tab{Id: 0, Name: "pp", Age: 18, Json: "jj", Time: time.Now(), StartTime: Now()}
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
