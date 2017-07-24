### sql-build
sql-build是一个支持条件控制的go语言sql拼接库.共分为4个部分,select,insert,update和delete四个部分,生成的结果为完整sql语句,需要和beego,xorm以及其它支持原生sql语句的orm配置使用,sql-build只做拼接工作.
    
主要功能有
- 无序的链式拼接
- 支持string类型的值注入检测
- 通过条件控制,决定数据有效性
    
需要注意的是:
- 目前只支持单表操作,不支持联表
- 参数的值在前,列名在后(这个用起来有点别扭,主要是为了使用idea的template)
- 可能缺少一些sql关键字,如果有需要,可以pull request

### 使用说明

### *select*
select方法可以支持以下函数,除了Select和String函数放在语句的头尾处,其余的都可以无序设置
- Select(table string) SelectInf
- Column(column string) SelectInf
- Where(value interface{}, key string, rules ... Rule) SelectInf
- Where_(value interface{}, key string, rules ... Rule) SelectInf
- Like(value string, key string) SelectInf
- In(values interface{}, key string) SelectInf
- NotIn(values interface{}, key string) SelectInf
- OrderBy(orderBy string) SelectInf
- Limit(limit int) SelectInf
- Offset(offset int) SelectInf
- GroupBy(groupBy string) SelectInf
- String() (string, error)

##### String 生成sql的函数
```go
import (
	"testing"
	"github.com/golyu/sql-build"
)

func TestTable(t *testing.T) {
    myTab:="myTab"
	sql, err := sqlBuild.Select(myTab).String()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(sql)
	tableName = string(i)
}
```
输出sql会是:
```go
SELECT * FROM myTab
```
如果把myTab置为空,err则打印
```go
The tabName can not be empty
```
以下示例为了简洁,不显示import

##### Column 查询结果显示列

```go
func TestColumn(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		Column("aaa").
		Column("bbb as xx").
		Column("").
		String()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(sql)
}
```
sql打印
```go
SELECT aaa,bbb as xx FROM myTab
```
空的数据无法通过数据有效性检测,不予组装(这个数据有效性条件是可以控制的)

##### Where条件控制语句
```go
func TestWhere(t *testing.T) {
	//
	sql, err := sqlBuild.Select("myTab").
		Where("nameValue", "name").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").
		Where(0, "phone ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where("nameValue", "name").
		Where(12, "age > ").
		Where("", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
sql输入结果为:
```go
	SELECT * FROM myTab WHERE name = 'nameValue'
	SELECT * FROM myTab WHERE name = 'nameValue' and age > 12
	SELECT * FROM myTab WHERE name = 'nameValue' and age > 12
	SELECT * FROM myTab WHERE name = 'nameValue' and age > 12
```
对于key中没有带>=<符号的,sql-build会自定补上=.
第三,四,明明有些条件存在,为什么生成的sql后丢失了呢,这就是sql-build的条件控制功能了,默认的数值类型<=0的条件,string类型值为""的都会被过滤掉,也就是数据无效,如果你说,我们的业务里也有值为0的啊,怎么办,sql-build可以自定义数据有效性过滤条件,在后面会详细说明.

##### Where_ 条件控制语句
```go
func TestWhere_(t *testing.T) {
	//
	sql, err := sqlBuild.Select("myTab").
		Where_("nameValue", "name").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").
		Where_(0, "phone ").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Where_("nameValue", "name").
		Where_(12, "age > ").
		Where_("", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
sql输出结果为:
```go
SELECT * FROM myTab WHERE name = 'nameValue'
SELECT * FROM myTab WHERE name = 'nameValue' and age > 12
```
err输出结果为:
```go
Fail to meet the condition
Fail to meet the condition
```
可以看出,`Where_`功能上和`Where`基本一致,只是把过滤替换成了抛出异常,这是对于有些业务需要指定的条件,而该条件可能为空的情况,做出的判断

##### GroupBy
```go
func TestGroupBy(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		GroupBy("aa").
		GroupBy("bb").
		GroupBy("").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
sql输出
```go
SELECT * FROM myTab GROUP BY aa,bb
```
##### OrderBy
```go
func TestOrderBy(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		OrderBy("aa1").
		OrderBy("-bb1").
		OrderBy("").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
sql输出为:
```go
SELECT * FROM myTab ORDER BY aa1,-bb1
```
正序直接写入列名,倒序在前加上-号

##### Limit Offset
```go
func TestLimit(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		Limit(10).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Limit(10).
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
sql输出为:
```go
SELECT * FROM myTab LIMIT 10
SELECT * FROM myTab LIMIT 10 OFFSET 2
```
err输出为:
```go
Need 'Offset' and 'Limit' are used together
```
Limit可以单独使用,Offset需要配合Offset使用,可以先Offset再Limit,我们中间的功能语句无序,只需要出现过就行
##### In
```go
func TestIn(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		In([]string{"小明", "小华"}, "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		In([]string{"小明", "小华"}, "name").
		In([]int{12, 22, 33, 16}, "age").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		In("小花", "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
sql输出:
```go
SELECT * FROM myTab WHERE name IN ('小明','小华')
SELECT * FROM myTab WHERE name IN ('小明','小华') and age IN (12,22,33,16)
```
err输出:
```go
The Value not have need type
```
in语句的值不允许为基本类型切片外的其它任何类型,包括基础类型,上述语句可以改写为
```go
sql, err = sqlBuild.Select("myTab").
	In([]string{"小花"}, "name").
	String()
```
##### NotIn
```go
func TestNotin(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		NotIn([]string{"小明", "小华"}, "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		NotIn([]string{"小明", "小华"}, "name").
		NotIn([]int{12, 22, 33, 16}, "age").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		NotIn("小花", "name").
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
##### Like
```go
func TestLike(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		Like("我的家", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Like("%我的家", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//
	sql, err = sqlBuild.Select("myTab").
		Like("我的家%", "address").String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
sql输出为:
```go
SELECT * FROM myTab WHERE address like '%我的家%'
SELECT * FROM myTab WHERE address like '%我的家'
SELECT * FROM myTab WHERE address like '我的家%'
```
如果like的值没有%,自动在前后补上%

以上就是sql-build中select功能的大致使用,功能块的函数是可以混合使用的,上面的举例只是部分写法,那我们合起来写一次
```go
func TestAll(t *testing.T) {
	sql, err := sqlBuild.Select("myTab").
		Where("一班", "class").
		Where(0, "age>").
		Where("c", "").
		Where_("男", "sex").
		In([]string{"语文", "数学"}, "hobby").
		NotIn([]int{6, 7}, "xx").
		GroupBy("xxx").
		GroupBy("xxxx").
		OrderBy("-id").
		Limit(10).
		Offset(2).String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
	//injection
	sql, err = sqlBuild.Select("myTab").
		Where("一班", "class").
		Where(0, "age>").
		Where("c", "").
		Where_("男' and 0<>(select count(*) from myTab) and ''=''", "sex").
		In([]string{"语文", "数学"}, "hobby").
		NotIn([]int{6, 7}, "xx").
		GroupBy("xxx").
		GroupBy("xxxx").
		OrderBy("-id").
		Limit(10).
		Offset(2).
		String()
	if err != nil {
		t.Error(err.Error())
		err = nil
	}
	t.Log(sql)
}
```
sql输出:
```go
SELECT * FROM myTab WHERE class = '一班' and sex = '男' and hobby IN ('语文','数学') and xx NOT IN (6,7) GROUP BY xxx,xxxx ORDER BY -id LIMIT 10 OFFSET 2
```
err输出:
```go
Injection err
```
检查出string类型的值里面包含注入的关键字

### *insert*

insert方法支持以下函数

- Insert(table string) InsertInf
- Value(value interface{}, rules ... Rule) InsertInf
- Values(value interface{}, rules ... Rule) InsertInf
- String() (string, error)

> 注意:insert方法因为需要在insert的方法上面加上insert的tag,value和values方法不可以同时使用

##### value

```go
type Tab struct {
	Id   int `insert:"id"`
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
```
首先,给`filed`加上`insert`的tag,然后调用value,传入指针,同`select`一样,value方法也可以自定义规则传入使用.
sql打印:
```go
INSERT INTO xx(id,name,age) VALUES (DEFAULT,'yiersan',18)
```

##### values 批量插入

```go
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
```
sql打印:
```go
INSERT INTO xx(id,name,age) VALUES (DEFAULT,'pp',18),(DEFAULT,'xx',16),(DEFAULT,'yiersan',18)
```


