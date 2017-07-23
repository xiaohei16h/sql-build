package sqlBuild

import (
	"github.com/pkg/errors"
	"strings"
	"github.com/golyu/sql-build/debug"
	"fmt"
)

var (
	ErrTabName   = errors.New("The tabName can not be empty")
	ErrValueType = errors.New("The Value not have need type")
	ErrInjection = errors.New("Injection err")
	ErrCondition = errors.New("Fail to meet the condition")
	ErrNoLimit   = errors.New("Need 'Offset' and 'Limit' are used together")
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
	err           error
	injection     bool
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
	if b.checkInjection(tabName) {
		return
	}
	if tabName == "" {
		b.err = ErrTabName
		debug.Warning("The tabName can not be empty")
	}
	b.tableName = tabName
}

func (b *BuildCore) orderBy(orderByValue string) {
	if b.injection || b.err != nil {
		return
	}
	if orderByValue == "" {
		debug.Println("The orderByValue is nil")
		return
	}
	if b.checkInjection(orderByValue) {
		return
	}
	b.orderValues = append(b.orderValues, orderByValue)
}

func (b *BuildCore) groupBy(groupByValue string) {
	if b.injection || b.err != nil {
		return
	}
	if groupByValue == "" {
		debug.Println("The groupByValue is nil")
		return
	}
	if b.checkInjection(groupByValue) {
		return
	}
	b.groupByValues = append(b.groupByValues, groupByValue)
}
func (b *BuildCore) like(likeValue, key string) {
	if b.injection || b.err != nil {
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
	if b.checkInjection(likeValue) {
		return
	}
	if strings.Contains(likeValue, "%") {
		b.likeValues = append(b.likeValues, key+" like '"+likeValue+"'")
	} else {
		b.likeValues = append(b.likeValues, key+" like "+strings.Join([]string{"'%", "%'"}, likeValue))
	}
}

func (b *BuildCore) in(inValues interface{}, key string) {
	if b.injection || b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The InKey can not be empty")
		return
	}
	if b.inMap == nil {
		b.inMap = make(map[string][]string)
	}
	result, err := b.getStrings(inValues)
	if err != nil {
		b.err = err
		return
	}
	if len(result) > 0 {
		b.inMap[key] = result
	}
}

func (b *BuildCore) notin(notinValues interface{}, key string) {
	if b.injection || b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The NotInKey can not be empty")
		return
	}
	if b.notinMap == nil {
		b.notinMap = make(map[string][]string)
	}
	result, err := b.getStrings(notinValues)
	if err != nil {
		b.err = err
		return
	}
	if len(result) > 0 {
		b.notinMap[key] = result
	}
}

func (b *BuildCore) where(whereValue interface{}, key string, rule Rule) {
	if b.injection || b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The WhereKey can not be empty")
		return
	}
	value, err := b.getKeyValues(whereValue, rule)
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
	if b.injection || b.err != nil {
		return
	}
	if key == "" {
		debug.Println("The WhereKey can not be empty")
		return
	}
	value, err := b.getKeyValues(whereValue, rule)
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

func (b *BuildCore) getKeyValues(values interface{}, rule Rule) (value string,
	err error) {
	switch value := values.(type) {
	case int:
		if int(value) > rule.IntValue {
			return fmt.Sprintf("%d", int(value)), nil
		}
	case int8:
		if int8(value) > rule.Int8Value {
			return fmt.Sprintf("%d", int8(value)), nil
		}
	case int16:
		if int16(value) > rule.Int16Value {
			return fmt.Sprintf("%d", int16(value)), nil
		}
	case int32:
		if int32(value) > rule.Int32Value {
			return fmt.Sprintf("%d", int32(value)), nil
		}
	case int64:
		if int64(value) > rule.Int64Value {
			return fmt.Sprintf("%d", int64(value)), nil
		}
	case uint:
		if uint(value) > rule.UintValue {
			return fmt.Sprintf("%d", uint(value)), nil
		}
	case uint8:
		if uint8(value) > rule.Uint8Value {
			return fmt.Sprintf("%d", uint8(value)), nil
		}
	case uint16:
		if uint16(value) > rule.Uint16Value {
			return fmt.Sprintf("%d", uint16(value)), nil
		}
	case uint32:
		if uint32(value) > rule.Uint32Value {
			return fmt.Sprintf("%d", uint32(value)), nil
		}
	case uint64:
		if uint64(value) > rule.Uint64Value {
			return fmt.Sprintf("%d", uint64(value)), nil
		}
	case float64:
		if float64(value) > rule.Float64Value {
			return fmt.Sprintf("%g", float64(value)), nil
		}
	case float32:
		if float32(value) > rule.Float32Value {
			return fmt.Sprintf("%g", float32(value)), nil
		}
	case string:
		if string(value) != rule.StringValue {
			if b.checkInjection(string(value)) {
				return "", ErrInjection
			}
			return strings.Join([]string{"'", "'"}, string(value)), nil
		}
	default:
		err = ErrValueType
	}
	return
}

func (b *BuildCore) getStrings(inValues interface{}) (strs []string, err error) {
	switch values := inValues.(type) {
	case []int:
		for _, v := range []int(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []int8:
		for _, v := range []int8(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []int16:
		for _, v := range []int16(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []int32:
		for _, v := range []int32(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []int64:
		for _, v := range []int64(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []uint:
		for _, v := range []uint(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []uint8:
		for _, v := range []uint8(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []uint16:
		for _, v := range []uint16(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []uint32:
		for _, v := range []uint32(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []uint64:
		for _, v := range []uint64(values) {
			str := fmt.Sprintf("%d", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []float64:
		for _, v := range []float64(values) {
			str := fmt.Sprintf("%g", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []float32:
		for _, v := range []float32(values) {
			str := fmt.Sprintf("%g", v)
			if str != "" {
				strs = append(strs, str)
			}
		}
	case []string:
		for _, v := range []string(values) {
			if v != "" {
				if b.checkInjection(v) {
					err = ErrInjection
					return
				}
				strs = append(strs, strings.Join([]string{"'", "'"}, v))
			}
		}
	default:
		err = ErrValueType
	}
	return
}

func (b *BuildCore) limit(limitValue int) {
	if limitValue > 0 {
		b.limitValue = limitValue
	} else {
		debug.Warning("limit can not < 1")
	}
}
func (b *BuildCore) offset(offsetValue int) {
	if offsetValue > 0 {
		b.offsetValue = offsetValue
	} else {
		debug.Warning("offset can not < 1")
	}
}

func (b *BuildCore) column(column string) {
	if b.checkInjection(column) {
		return
	}
	if column != "" {
		b.columnValues = append(b.columnValues, column)
	} else {
		debug.Println("column is nil")
	}
}

func (b *BuildCore) checkInjection(val string) bool {
	if val != "" {
		val = strings.ToLower(val)
		b.injection = strings.Contains(val, "select ") ||
			strings.Contains(val, "update ") ||
			strings.Contains(val, "delete ") ||
			strings.Contains(val, "insert ")||
			strings.Contains(val, "declare ")||
			strings.Contains(val, "drop ")||
			strings.Contains(val, "create ")||
			strings.Contains(val, "alter ")
		debug.Error("Injection <" + val + ">")
	}
	return b.injection
}
