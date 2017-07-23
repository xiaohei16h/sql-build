package sqlBuild

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
)

type SqlBuilder struct {
	tableName string
	action    int
	set       string
	column    []string
	where     string
	order     string
	limit     string
	offset    string
	insertKey []string
	insertVal []string
	groupBy   string
	err       error
}

func (this *SqlBuilder) Set_(value interface{}, set string) (*SqlBuilder) {
	if set == "" {
		panic("no where arguments 'set'")
	}
	if !strings.Contains(set, "=") {
		set += "="
	}
	startLen := len(set)
	switch value := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		set = set + " " + fmt.Sprintf("%d", value)
	case float64, float32:
		set = set + " " + fmt.Sprintf("%f", value)
	case string:
		set = set + " '" + value + "'"
	default:
		panic("no type")
	}
	if len(set) > startLen {
		if this.set != "" {
			this.set += ","
		}
		this.set += set
	}
	return this
}

func (this *SqlBuilder) Set(value interface{}, set string) (*SqlBuilder) {
	if !strings.Contains(set, "=") {
		set += "="
	}
	startLen := len(set)
	switch value := value.(type) {
	case int:
		if int(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case int8:
		if int8(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case int16:
		if int16(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case int32:
		if int32(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case int64:
		if int64(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case uint:
		if uint(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case uint8:
		if uint8(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case uint16:
		if uint16(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case uint32:
		if uint32(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case uint64:
		if uint64(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%d", value)
		} else {
			set = ""
		}
	case float64:
		if float64(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%f", value)
		} else {
			set = ""
		}
	case float32:
		if float32(value) > 0 {
			set = " " + set + " " + fmt.Sprintf("%f", value)
		} else {
			set = ""
		}
	case string:
		if string(value) != "" {
			set = " " + set + " '" + value + "'"
		} else {
			set = ""
		}
	default:
		panic("no type")
	}
	if len(set) > startLen {
		if this.set != "" {
			this.set += ","
		}
		this.set += set
	}
	return this
}
func (this *SqlBuilder) Where_(value interface{}, why string) (*SqlBuilder) {
	if why == "" {
		panic("no where arguments 'why'")
	}
	if !strings.ContainsAny(why, ">=<") {
		why += "="
	}
	startLen := len(why)
	switch value := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		why = why + " " + fmt.Sprintf("%d", value)
	case float64, float32:
		why = why + " " + fmt.Sprintf("%f", value)
	case string:
		why = why + " '" + value + "'"
	default:
		panic("no type")
	}
	if len(why) > startLen {
		if this.where != "" {
			this.where += " AND "
		}
		this.where += why
	}
	return this
}
func (this *SqlBuilder) Insert(table string) (*SqlBuilder) {
	this.action = 1
	this.tableName = table
	return this
}
func (this *SqlBuilder) Delete(table string) (*SqlBuilder) {
	this.action = 2
	this.tableName = table
	return this
}
func (this *SqlBuilder) Update(table string) (*SqlBuilder) {
	this.action = 3
	this.tableName = table
	return this
}
func (this *SqlBuilder) Select(table string) (*SqlBuilder) {
	this.action = 4
	this.tableName = table
	return this
}

func (this *SqlBuilder) Column(column string) (*SqlBuilder) {
	if column != "" {
		if this.column == nil {
			this.column = make([]string, 0)
		}
		this.column = append(this.column, column)
	}
	return this
}
func (this *SqlBuilder) Where(value interface{}, why string) (*SqlBuilder) {
	if !strings.ContainsAny(why, ">=<") {
		why += "="
	}
	startLen := len(why)
	switch value := value.(type) {
	case int:
		if int(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case int8:
		if int8(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case int16:
		if int16(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case int32:
		if int32(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case int64:
		if int64(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case uint:
		if uint(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case uint8:
		if uint8(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case uint16:
		if uint16(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case uint32:
		if uint32(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case uint64:
		if uint64(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%d", value)
		} else {
			why = ""
		}
	case float64:
		if float64(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%f", value)
		} else {
			why = ""
		}
	case float32:
		if float32(value) > 0 {
			why = " " + why + " " + fmt.Sprintf("%f", value)
		} else {
			why = ""
		}
	case string:
		if string(value) != "" {
			why = " " + why + " '" + value + "'"
		} else {
			why = ""
		}
	default:
		panic("no type")
	}
	if len(why) > startLen {
		if this.where != "" {
			this.where += " AND"
		}
		this.where += why
	}
	return this
}

func (this *SqlBuilder) Like(value string, why string) *SqlBuilder {
	if why == "" {
		panic("no where arguments 'why'")
	}

	if strings.ContainsAny(why, ">=<") {
		panic("SQL 'like' not allow '>' or '=' or '<'")
	}
	if value != "" {
		if strings.Count(value, "%") == len(value) {
			return this
		}
		if !strings.Contains(value, "%") {
			value = "%" + value + "%"
		}

		if this.where != "" {
			this.where += " AND "
		}
		this.where += why + " like '" + value + "'"
	}
	return this
}

func (this *SqlBuilder) Value_(value interface{}, key string) (*SqlBuilder) {
	if key == "" {
		panic("no where arguments 'key'")
	}
	var val string
	switch value := value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		val = fmt.Sprintf("%d", value)
	case float64, float32:
		val = fmt.Sprintf("%f", value)
	case string:
		val = "'" + fmt.Sprintf("%s", value) + "'"
	default:
		panic("no type")
	}
	if len(val) > 0 {
		this.insertKey = append(this.insertKey, key)
		this.insertVal = append(this.insertVal, val)
	}
	return this
}
func (this *SqlBuilder) Value(value interface{}, key string) (*SqlBuilder) {
	if key == "" {
		return this
	}
	var val string
	switch value := value.(type) {
	case int:
		if int(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case int8:
		if int8(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case int16:
		if int16(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case int32:
		if int32(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case int64:
		if int64(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case uint:
		if uint(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case uint8:
		if uint8(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case uint16:
		if uint16(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case uint32:
		if uint32(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case uint64:
		if uint64(value) > 0 {
			val = fmt.Sprintf("%d", value)
		} else {
			return this
		}
	case float64:
		if float64(value) > 0 {
			val = fmt.Sprintf("%f", value)
		} else {
			return this
		}
	case float32:
		if float32(value) > 0 {
			val = fmt.Sprintf("%f", value)
		} else {
			return this
		}
	case string:
		if string(value) != "" {
			val = "'" + fmt.Sprintf("%s", value) + "'"
		} else {
			return this
		}
	default:
		panic("no type")
	}
	if len(val) > 0 {
		this.insertKey = append(this.insertKey, key)
		this.insertVal = append(this.insertVal, val)
	}
	return this
}

func (this *SqlBuilder) In(values interface{}, key string) (*SqlBuilder) {
	if key == "" {
		panic("no in arguments 'key'")
	}
	var val string
	ins := "("
	switch value := values.(type) {
	case []int:
		for _, v := range []int(value) {
			val = fmt.Sprintf("%d", v)
			if len(val) > 0 {
				if ins != "(" {
					ins += ","
				}
				ins += val
			}
		}
	case []string:
		for _, v := range []string(value) {
			val = "'" + fmt.Sprintf("%s", v) + "'"
			if len(val) > 0 {
				if ins != "(" {
					ins += ","
				}
				ins += val
			}
		}
	default:
		panic("no type (Will be perfect)")
	}
	if ins != "(" {
		ins += ")"
		if this.where != "" {
			this.where += " AND "
		}
		this.where += key + " in " + ins
	}
	return this
}

func (this *SqlBuilder) OrderBy(orderBy string) (*SqlBuilder) {
	if orderBy != "" {
		this.order = orderBy
	}
	if this.order != "" {
		this.order = " ORDER BY " + orderBy
	}
	return this
}
func (this *SqlBuilder) Limit(limit int) (*SqlBuilder) {
	if limit > 0 {
		this.limit = " LIMIT " + strconv.Itoa(limit)
	}
	return this
}
func (this *SqlBuilder) Offset(offset int) (*SqlBuilder) {
	if offset > -1 && this.limit != "" {
		this.offset = " OFFSET " + strconv.Itoa(offset)
	}
	return this
}

func (this *SqlBuilder) GroupBy(groupBy string) (*SqlBuilder) {
	if groupBy != "" {
		this.groupBy = " GROUP BY " + groupBy
	}
	return this
}

//LIMIT 0 OFFSET
func (this *SqlBuilder) String() (string, error) {
	if this.where != "" {
		this.where = " Where " + this.where
	}
	if this.set != "" {
		this.set = " SET " + this.set
	}
	switch this.action {
	case 1:
		keyLen := len(this.insertKey)
		valLen := len(this.insertVal)
		if keyLen > 0 && valLen > 0 && keyLen == valLen {
			keys := strings.Join(this.insertKey, ",")
			vals := strings.Join(this.insertVal, ",")
			return "INSERT INTO " + this.tableName + "(" + keys + ") VALUES(" + vals + ")", nil
		} else {
			return "", errors.New("data count not equal")
		}
	case 2:
		return "DELETE FROM " + this.tableName + this.where, nil
	case 3:
		if this.set == "" {
			return "", errors.New("not need update column")
		}
		return "UPDATE " + this.tableName + this.set + this.where+ this.order + this.limit + this.offset, nil
	case 4:
		if len(this.column) == 0 {
			return "SELECT * FROM " + this.tableName + this.where + this.groupBy + this.order + this.limit + this.offset, nil
		} else {
			c := ""
			for _, v := range this.column {
				if c != "" {
					c += ","
				}
				c += v
			}
			return "SELECT " + c + " FROM " + this.tableName + this.where + this.groupBy + this.order + this.limit +
				this.offset, nil
		}
	default:
		return "", errors.New("please input action(insert or update or delete)")
	}
}
