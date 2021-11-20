package stringx

import (
	"encoding/json"
	"strconv"
	"strings"
	"unsafe"
)

func NewWithByte(b []byte) Str {
	return *(*Str)(unsafe.Pointer(&b))
}

// Str 字符串类型转换
type Str string

func (str Str) String() string {
	return string(str)
}

// Bytes 转换为[]byte
func (str Str) Bytes() []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

// Bool 转换为bool
func (str Str) Bool() (bool, error) {
	strBool, err := strconv.ParseBool(str.String())
	if err != nil {
		return false, err
	}
	return strBool, nil
}

// DefaultBool 转换为bool，如果出现错误则使用默认值
func (str Str) DefaultBool(defaultVal bool) bool {
	strBool, err := str.Bool()
	if err != nil {
		return defaultVal
	}
	return strBool
}

// Int64 转换为int64
func (str Str) Int64() (int64, error) {
	strInt, err := strconv.ParseInt(str.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return strInt, nil
}

// DefaultInt64 转换为int64，如果出现错误则使用默认值
func (str Str) DefaultInt64(defaultVal int64) int64 {
	strInt, err := str.Int64()
	if err != nil {
		return defaultVal
	}
	return strInt
}

// Int 转换为int
func (str Str) Int() (int, error) {
	strInt, err := str.Int64()
	if err != nil {
		return 0, err
	}
	return int(strInt), nil
}

// DefaultInt 转换为int，如果出现错误则使用默认值
func (str Str) DefaultInt(defaultVal int) int {
	strInt, err := str.Int()
	if err != nil {
		return defaultVal
	}
	return strInt
}

// Uint64 转换为uint64
func (str Str) Uint64() (uint64, error) {
	strInt, err := strconv.ParseUint(str.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return strInt, nil
}

// DefaultUint64 转换为uint64，如果出现错误则使用默认值
func (str Str) DefaultUint64(defaultVal uint64) uint64 {
	strInt, err := str.Uint64()
	if err != nil {
		return defaultVal
	}
	return strInt
}

// Uint 转换为uint
func (str Str) Uint() (uint, error) {
	strInt, err := str.Uint64()
	if err != nil {
		return 0, err
	}
	return uint(strInt), nil
}

// DefaultUint 转换为uint，如果出现错误则使用默认值
func (str Str) DefaultUint(defaultVal uint) uint {
	strInt, err := str.Uint()
	if err != nil {
		return defaultVal
	}
	return strInt
}

// Float64 转换为float64
func (str Str) Float64() (float64, error) {
	strFloat, err := strconv.ParseFloat(str.String(), 64)
	if err != nil {
		return 0, err
	}
	return strFloat, nil
}

// DefaultFloat64 转换为float64，如果出现错误则使用默认值
func (str Str) DefaultFloat64(defaultVal float64) float64 {
	strFloat, err := str.Float64()
	if err != nil {
		return defaultVal
	}
	return strFloat
}

// Float32 转换为float32
func (str Str) Float32() (float32, error) {
	strFloat, err := str.Float64()
	if err != nil {
		return 0, err
	}
	return float32(strFloat), nil
}

// DefaultFloat32 转换为float32，如果出现错误则使用默认值
func (str Str) DefaultFloat32(defaultVal float32) float32 {
	strFloat, err := str.Float32()
	if err != nil {
		return defaultVal
	}
	return strFloat
}

// ToJSON 转换为JSON
func (str Str) ToJSON(value interface{}) error {
	return json.Unmarshal(str.Bytes(), value)
}

// SubStr 截取
func (str Str) SubStr(pos int, length int) string {
	runes := []rune(str.String())
	lens := pos + length
	if lens > len(runes) {
		lens = len(runes)
	}
	return string(runes[pos:lens])
}

// GetParentDir 获取父级目录
func (str Str) GetParentDir() string {
	return str.SubStr(0, strings.LastIndex(str.String(), "/"))
}

// ToInt 字符串转 int
func ToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// StructToJsonStr struct 转 json 字符串
func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

// JsonStrToMap json 字符串 转 map
func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}

// StructToMap struct 转 map
func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}
