package convert

import (
	"fmt"
	"strconv"
)

// ToString converts any value into a string.
func ToString(value interface{}, args ...int) string {
	switch v := value.(type) {
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', optGet(args, 0, -1), optGet(args, 1, 32))
	case float64:
		return strconv.FormatFloat(v, 'f', optGet(args, 0, -1), optGet(args, 1, 64))
	case int:
		return strconv.FormatInt(int64(v), optGet(args, 0, 10))
	case int8:
		return strconv.FormatInt(int64(v), optGet(args, 0, 10))
	case int16:
		return strconv.FormatInt(int64(v), optGet(args, 0, 10))
	case int32:
		return strconv.FormatInt(int64(v), optGet(args, 0, 10))
	case int64:
		return strconv.FormatInt(v, optGet(args, 0, 10))
	case uint:
		return strconv.FormatUint(uint64(v), optGet(args, 0, 10))
	case uint8:
		return strconv.FormatUint(uint64(v), optGet(args, 0, 10))
	case uint16:
		return strconv.FormatUint(uint64(v), optGet(args, 0, 10))
	case uint32:
		return strconv.FormatUint(uint64(v), optGet(args, 0, 10))
	case uint64:
		return strconv.FormatUint(v, optGet(args, 0, 10))
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// ToInt converts any value into a int.
func ToInt(value interface{}) (int, error) {
	v, err := parseInt(value, 0)
	return int(v), err
}

// ToInt8 converts any value into a int8.
func ToInt8(value interface{}) (int8, error) {
	v, err := parseInt(value, 8)
	return int8(v), err
}

// ToInt16 converts any value into a int16.
func ToInt16(value interface{}) (int16, error) {
	v, err := parseInt(value, 16)
	return int16(v), err
}

// ToInt32 converts any value into a int32.
func ToInt32(value interface{}) (int32, error) {
	v, err := parseInt(value, 32)
	return int32(v), err
}

// ToInt64 converts any value into a int64.
func ToInt64(value interface{}) (int64, error) {
	v, err := parseInt(value, 64)
	return v, err
}

// ToUint converts any value into a uint.
func ToUint(value interface{}) (uint, error) {
	v, err := parseUint(value, 0)
	return uint(v), err
}

// ToUint8 converts any value into a uint8.
func ToUint8(value interface{}) (uint8, error) {
	v, err := parseUint(value, 8)
	return uint8(v), err
}

// ToUint16 converts any value into a uint16.
func ToUint16(value interface{}) (uint16, error) {
	v, err := parseUint(value, 16)
	return uint16(v), err
}

// ToUint32 converts any value into a uint32.
func ToUint32(value interface{}) (uint32, error) {
	v, err := parseUint(value, 32)
	return uint32(v), err
}

// ToUint64 converts any value into a uint64.
func ToUint64(value interface{}) (uint64, error) {
	v, err := parseUint(value, 64)
	return v, err
}

func parseInt(value interface{}, n int) (int64, error) {
	v, err := strconv.ParseInt(ToString(value), 10, n)
	return v, err
}

func parseUint(value interface{}, n int) (uint64, error) {
	v, err := strconv.ParseUint(ToString(value), 10, n)
	return v, err
}

func optGet(args []int, i int, opt int) int {
	if i >= 0 && i < len(args) {
		return args[i]
	}
	return opt
}
