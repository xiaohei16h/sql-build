package sqlBuild

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func CheckInjection(val string) (injection bool) {
	return
	//if val != "" {
	//	val = strings.ToLower(val)
	//	injection = strings.Contains(val, "select ") ||
	//		strings.Contains(val, "update ") ||
	//		strings.Contains(val, "delete ") ||
	//		strings.Contains(val, "insert ") ||
	//		strings.Contains(val, "declare ") ||
	//		strings.Contains(val, "drop ") ||
	//		strings.Contains(val, "create ") ||
	//		strings.Contains(val, "alter ")
	//	if injection {
	//		debug.Error("Injection <" + val + ">")
	//	}
	//}
	//return
}
func GetInValues(inValues interface{}) (strs []string, err error) {
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
				if CheckInjection(v) {
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

func GetWhereSetFuncValues(values interface{}, rule Rule) (value string,
	err error) {
	return getWhereSetValues(values, rule, func(value string) string {
		return string(value)
	})
}

func GetWhereSetValues(values interface{}, rule Rule) (value string,
	err error) {
	return getWhereSetValues(values, rule, func(value string) string {
		return strings.Join([]string{"'", "'"}, string(value))
	})
}

func getWhereSetValues(values interface{}, rule Rule, f func(value string) string) (value string,
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
			return fmt.Sprintf("%f", float64(value)), nil
		}
	case float32:
		if float32(value) > rule.Float32Value {
			return fmt.Sprintf("%f", float32(value)), nil
		}
	case string:
		if string(value) != rule.StringValue {
			if CheckInjection(string(value)) {
				return "", ErrInjection
			}
			return f(string(value)), nil
		}
	default:
		err = ErrValueType
	}
	return
}

type JsonInterface interface {
	TrimQuote(jsonStr []byte) bool
}

// 得到数据类型
func GetValue(value reflect.Value, rule Rule) (string, error) {
	toJson := func(v reflect.Value) (string, error) {
		if txtByte, err := json.Marshal(v.Interface()); err != nil {
			return "", err
		} else {
			if jsonInterface, ok := v.Interface().(JsonInterface); ok {
				if jsonInterface.TrimQuote(txtByte) {
					txtByte = bytes.Trim(txtByte, `"`)
				}
			}
			return strings.Join([]string{"'", "'"}, Escape(string(txtByte))), nil
		}
	}

	switch value.Kind() {
	case reflect.Bool:
		if value.Bool() {
			return "1", nil
		} else {
			return "0", nil
		}
	case reflect.Int:
		temp := value.Int()
		if int(temp) > rule.IntValue {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Int8:
		temp := value.Int()
		if int8(temp) > rule.Int8Value {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Int16:
		temp := value.Int()
		if int16(temp) > rule.Int16Value {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Int32:
		temp := value.Int()
		if int32(temp) > rule.Int32Value {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Int64:
		temp := value.Int()
		if temp > rule.Int64Value {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Uint:
		temp := value.Uint()
		if uint(temp) > rule.UintValue {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Uint8:
		temp := value.Uint()
		if uint8(temp) > rule.Uint8Value {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Uint16:
		temp := value.Uint()
		if uint16(temp) > rule.Uint16Value {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Uint32:
		temp := value.Uint()
		if uint32(temp) > rule.Uint32Value {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Uint64:
		temp := value.Uint()
		if uint64(temp) > rule.Uint64Value {
			return fmt.Sprintf("%d", temp), nil
		}
	case reflect.Float32:
		temp := value.Float()
		if float32(temp) > rule.Float32Value {
			return fmt.Sprintf("%f", temp), nil
		}
	case reflect.Float64:
		temp := value.Float()
		if float64(temp) > rule.Float64Value {
			return fmt.Sprintf("%f", temp), nil
		}
	case reflect.String:
		temp := value.String()
		if temp != rule.StringValue {
			if CheckInjection(temp) {
				return "", ErrInjection
			}
			return strings.Join([]string{"'", "'"}, Escape(temp)), nil
		}
	case reflect.Slice:
		if value.Len() > rule.SliceLength {
			return toJson(value)
		}
	case reflect.Array:
		if value.Len() > rule.ArrayLength {
			return toJson(value)
		}
	case reflect.Map:
		if value.Len() > rule.MapLength {
			return toJson(value)
		}
	case reflect.Struct:
		if v, ok := value.Interface().(time.Time); ok {
			if v.IsZero() {
				break
			}
			return strings.Join([]string{"'", "'"}, v.Format("2006-01-02 15:04:05.000")), nil
		}

		if rule.StructForce || !reflect.DeepEqual(reflect.New(value.Type()).Elem().Interface(), value.Interface()) {
			return toJson(value)
		}
	case reflect.Interface:
		return toJson(value)
	case reflect.Ptr:
		if !value.IsNil() {
			return GetValue(value.Elem(), rule)
		}
	}

	return "DEFAULT", nil
}

//func MysqlRealEscapeString(value string) string {
//	replace := map[string]string{"\\":"\\\\", "'":`\'`, "\\0":"\\\\0", "\n":"\\n", "\r":"\\r", `"`:`\"`, "\x1a":"\\Z"}
//
//	for b, a := range replace {
//		value = strings.Replace(value, b, a, -1)
//	}
//
//	return value
//}

func Escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
			break
		case '\n': /* Must be escaped for logs */
			escape = 'n'
			break
		case '\r':
			escape = 'r'
			break
		case '\\':
			escape = '\\'
			break
		case '\'':
			escape = '\''
			break
		case '"': /* Better safe than sorry */
			escape = '"'
			break
		case '\032': /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}
