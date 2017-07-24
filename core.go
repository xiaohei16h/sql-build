package sqlBuild

import (
	"github.com/pkg/errors"
	"strings"
	"github.com/golyu/sql-build/debug"
	"reflect"
	"sync"
)

var (
	ErrTabName      = errors.New("The tabName can not be empty")
	ErrValueType    = errors.New("The Value not have need type")
	ErrInjection    = errors.New("Injection err")
	ErrCondition    = errors.New("Fail to meet the condition")
	ErrNoLimit      = errors.New("Need 'Offset' and 'Limit' are used together")
	ErrInsertColumn = errors.New("Not found Insert Column")
	ErrInsertValue  = errors.New("Not found Insert Data")
)

type BuildCore struct {
	tableName     string
	columnValues  []string
	whereValues   []string
	orderValues   []string
	limitValue    int
	offsetValue   int
	inMap         map[string][]string
	notinMap      map[string][]string
	groupByValues []string
	likeValues    []string

	//insert
	insertColumns     []string
	insertTags        []int
	insertValues      []string
	insertValuesMutex sync.Mutex

	err error
}
type Rule struct {
	IntValue     int
	Int8Value    int8
	Int16Value   int16
	Int32Value   int32
	Int64Value   int64
	UintValue    uint
	Uint8Value   uint8
	Uint16Value  uint16
	Uint32Value  uint32
	Uint64Value  uint64
	Float32Value float32
	Float64Value float64
	StringValue  string
}

func (b *BuildCore) setTabName(tabName string) {
	if CheckInjection(tabName) {
		b.err = ErrInjection
		return
	}
	if tabName == "" {
		b.err = ErrTabName
		debug.Warning("The tabName can not be empty")
	}
	b.tableName = tabName
}

func (b *BuildCore) orderBy(orderByValue string) {
	if b.err != nil {
		return
	}
	if orderByValue == "" {
		debug.Println("The orderByValue is nil")
		return
	}
	if CheckInjection(orderByValue) {
		b.err = ErrInjection
		return
	}
	b.orderValues = append(b.orderValues, orderByValue)
}

func (b *BuildCore) groupBy(groupByValue string) {
	if b.err != nil {
		return
	}
	if groupByValue == "" {
		debug.Println("The groupByValue is nil")
		return
	}
	if CheckInjection(groupByValue) {
		b.err = ErrInjection
		return
	}
	b.groupByValues = append(b.groupByValues, groupByValue)
}
func (b *BuildCore) like(likeValue, key string) {
	if b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The LikeKey can not be empty")
		return
	}
	if likeValue == "" {
		debug.Println("The likeValue is nil")
		return
	}
	if CheckInjection(likeValue) {
		b.err = ErrInjection
		return
	}
	if strings.Contains(likeValue, "%") {
		b.likeValues = append(b.likeValues, key+" like '"+likeValue+"'")
	} else {
		b.likeValues = append(b.likeValues, key+" like "+strings.Join([]string{"'%", "%'"}, likeValue))
	}
}

func (b *BuildCore) in(inValues interface{}, key string) {
	if b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The InKey can not be empty")
		return
	}
	if b.inMap == nil {
		b.inMap = make(map[string][]string)
	}
	result, err := GetInValues(inValues)
	if err != nil {
		b.err = err
		return
	}
	if len(result) > 0 {
		b.inMap[key] = result
	}
}

func (b *BuildCore) notin(notinValues interface{}, key string) {
	if b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The NotInKey can not be empty")
		return
	}
	if b.notinMap == nil {
		b.notinMap = make(map[string][]string)
	}
	result, err := GetInValues(notinValues)
	if err != nil {
		b.err = err
		return
	}
	if len(result) > 0 {
		b.notinMap[key] = result
	}
}

func (b *BuildCore) where(whereValue interface{}, key string, rule Rule) {
	if b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The WhereKey can not be empty")
		return
	}
	value, err := GetWhereValues(whereValue, rule)
	if err != nil {
		b.err = err
		return
	}
	if value != "" && value != "''" {
		if !strings.ContainsAny(key, ">=<") {
			key += " = "
		}
		b.whereValues = append(b.whereValues, key+value)
	}
}
func (b *BuildCore) where_(whereValue interface{}, key string, rule Rule) {
	if b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The WhereKey can not be empty")
		return
	}
	value, err := GetWhereValues(whereValue, rule)
	if err != nil {
		b.err = err
		return
	}
	if value != "" && value != "''" {
		if !strings.ContainsAny(key, ">=<") {
			key += " = "
		}
		b.whereValues = append(b.whereValues, key+value)
	} else {
		b.err = ErrCondition
	}
}



//func ( *BuildCore) getStrings(inValues interface{}) (strs []string, err error) {
//	switch values := inValues.(type) {
//	case []int:
//		for _, v := range []int(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []int8:
//		for _, v := range []int8(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []int16:
//		for _, v := range []int16(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []int32:
//		for _, v := range []int32(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []int64:
//		for _, v := range []int64(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []uint:
//		for _, v := range []uint(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []uint8:
//		for _, v := range []uint8(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []uint16:
//		for _, v := range []uint16(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []uint32:
//		for _, v := range []uint32(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []uint64:
//		for _, v := range []uint64(values) {
//			str := fmt.Sprintf("%d", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []float64:
//		for _, v := range []float64(values) {
//			str := fmt.Sprintf("%g", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []float32:
//		for _, v := range []float32(values) {
//			str := fmt.Sprintf("%g", v)
//			if str != "" {
//				strs = append(strs, str)
//			}
//		}
//	case []string:
//		for _, v := range []string(values) {
//			if v != "" {
//				if CheckInjection(v) {
//					err = ErrInjection
//					return
//				}
//				strs = append(strs, strings.Join([]string{"'", "'"}, v))
//			}
//		}
//	default:
//		err = ErrValueType
//	}
//	return
//}

func (b *BuildCore) limit(limitValue int) {
	if b.err != nil {
		return
	}
	if limitValue > 0 {
		b.limitValue = limitValue
	} else {
		debug.Warning("limit can not < 1")
	}
}
func (b *BuildCore) offset(offsetValue int) {
	if b.err != nil {
		return
	}
	if offsetValue > 0 {
		b.offsetValue = offsetValue
	} else {
		debug.Warning("offset can not < 1")
	}
}

func (b *BuildCore) column(column string) {
	if b.err != nil {
		return
	}
	if CheckInjection(column) {
		b.err = ErrInjection
		return
	}
	if column != "" {
		b.columnValues = append(b.columnValues, column)
	} else {
		debug.Println("column is nil")
	}
}

func (b *BuildCore) setValueColumns(ty reflect.Type, tag string) {
	for i := 0; i < ty.NumField(); i++ {
		name := ty.Field(i).Tag.Get(tag)
		if name != "" {
			b.insertColumns = append(b.insertColumns, name)
			b.insertTags = append(b.insertTags, i)
		}
	}
}

func (b *BuildCore) value(ind reflect.Value, rule Rule, wg... *sync.WaitGroup) {
	if len(wg)>0{
		defer wg[0].Done()
	}
	if b.err != nil {
		return
	}
	//val := reflect.ValueOf(in)
	//ind := val.Elem()
	var values []string
	for _, v := range b.insertTags {
		value, err := GetValue(ind.Field(v), rule)
		if err != nil {
			b.err = err
			return
		}
		values = append(values, value)
	}
	if len(values) > 0 {
		row := strings.Join(values, ",")
		func() {
			b.insertValuesMutex.Lock()
			defer b.insertValuesMutex.Unlock()
			b.insertValues = append(b.insertValues, row)
		}()
	} else {
		debug.Println("Insert no found data")
	}
}
