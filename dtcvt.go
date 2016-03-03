package dtcvt

import (
	// "bytes"
	// "encoding/binary"
	"strconv"
)

/****** To String ******/

/**
 * 转String
 * @param interface{} data 源数据
 * @return (bool, string)
 */
func String(data interface{}) (ok bool, result string) {
	ok = true
	switch data.(type) {
	case string:
		result = data.(string)
		break
	case int:
		result = strconv.Itoa(data.(int))
		break
	case int8:
		result = strconv.Itoa(int(data.(int8)))
		break
	case int16:
		result = strconv.Itoa(int(data.(int16)))
		break
	case int32:
		result = strconv.Itoa(int(data.(int32)))
		break
	case int64:
		result = strconv.Itoa(int(data.(int64)))
		break
	case uint8:
		result = strconv.FormatUint(uint64(data.(uint8)), 10)
		break
	case uint16:
		result = strconv.FormatUint(uint64(data.(uint16)), 10)
		break
	case uint32:
		result = strconv.FormatUint(uint64(data.(uint32)), 10)
		break
	case uint64:
		result = strconv.FormatUint(data.(uint64), 10)
		break
	case float32:
		result = strconv.FormatFloat(data.(float64), 'f', -1, 32)
		break
	case float64:
		result = strconv.FormatFloat(data.(float64), 'f', -1, 64)
		break
	case []byte:
		result = string(data.([]byte))
	case bool:
		if data.(bool) {
			result = "true"
		} else {
			result = "false"
		}
		break
	default:
		ok = false
		result = ""
		break
	}
	return ok, result
}

/**
 * 转String
 * @param interface{} data 源数据
 * @param string defaultValue 默认值
 * @return string
 */
func MustString(data interface{}, defaultValue ...string) string {
	ok, r := String(data)
	if ok {
		return r
	} else {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return ""
		}
	}
}

/****** To Int ******/

//TODO:后面加四舍五入、ceil等
/**
 * 转Int64
 * @param interface{} data 源数据
 * @return (bool, int64)
 */
func Int64(data interface{}) (ok bool, result int64) {
	var err error
	ok = true
	switch data.(type) {
	case string:
		result, err = strconv.ParseInt(data.(string), 10, 64)
		if err != nil {
			ok = false
		}
		break
	case int:
		result = int64(data.(int))
		break
	case int8:
		result = int64(data.(int8))
		break
	case int16:
		result = int64(data.(int16))
		break
	case int32:
		result = int64(data.(int32))
		break
	case int64:
		result = int64(data.(int64))
		break
	case uint8:
		result = int64(data.(uint8))
		break
	case uint16:
		result = int64(data.(uint16))
		break
	case uint32:
		result = int64(data.(uint32))
		break
	case uint64:
		result = int64(data.(uint64))
		break
	case float32:
		result = int64(data.(float32))
		break
	case float64:
		result = int64(data.(float64))
		break
		//TODO
	case []byte:
		//1
		// b_tmp := data.([]byte)
		// b_buf := bytes.NewBuffer(b_tmp)
		// var result_tmp int32
		// err = binary.Read(b_buf, binary.BigEndian, &result_tmp)
		// if err != nil {
		// 	ok = false
		// 	result = 0
		// } else {
		// 	result = int64(result_tmp)
		// }

		//2
		r_str := string(data.([]byte))
		result, err = strconv.ParseInt(r_str, 10, 64)
		if err != nil {
			ok = false
		}
		break
	case bool:
		if data.(bool) {
			result = 1
		} else {
			result = 0
		}
		break
	default:
		ok = false
		result = 0
		break
	}
	return ok, result
}

/**
 * 转Int
 * @param interface{} data 源数据
 * @return (bool, int)
 */
func Int(data interface{}) (bool, int) {
	ok, result := Int64(data)
	return ok, int(result)
}

/**
 * 转int
 * @param interface{} data 源数据
 * @param int defaultValue 默认值
 * @return int
 */
func MustInt(data interface{}, defaultValue ...int) int {
	ok, r := Int(data)
	if ok {
		return r
	} else {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return 0
		}
	}
}

/****** To Float ******/

/**
 * 转Float64
 * @param interface{} data 源数据
 * @return (bool, float64)
 */
func Float64(data interface{}) (ok bool, result float64) {
	var err error
	ok = true
	switch data.(type) {
	case string:
		result, err = strconv.ParseFloat(data.(string), 64)
		if err != nil {
			ok = false
		}
		break
	case int:
		result = float64(data.(int))
		break
	case int8:
		result = float64(data.(int8))
		break
	case int16:
		result = float64(data.(int16))
		break
	case int32:
		result = float64(data.(int32))
		break
	case int64:
		result = float64(data.(int64))
		break
	case uint8:
		result = float64(data.(uint8))
		break
	case uint16:
		result = float64(data.(uint16))
		break
	case uint32:
		result = float64(data.(uint32))
		break
	case uint64:
		result = float64(data.(uint64))
		break
	case float32:
		result = float64(data.(float32))
		break
	case float64:
		result = data.(float64)
		break
		//TODO
	case []byte:
		r_str := string(data.([]byte))
		result, err = strconv.ParseFloat(r_str, 64)
		if err != nil {
			ok = false
		}
		break
	case bool:
		if data.(bool) {
			result = 1
		} else {
			result = 0
		}
		break
	default:
		ok = false
		result = 0
		break
	}
	return ok, result
}

/**
 * 转float32
 * @param interface{} data 源数据
 * @return (bool, float32)
 */
func Float32(data interface{}) (bool, float32) {
	ok, result := Float64(data)
	return ok, float32(result)
}

/**
 * 转float64
 * @param interface{} data 源数据
 * @param float64 defaultValue 默认值
 * @return float64
 */
func MustFloat64(data interface{}, defaultValue ...float64) float64 {
	ok, r := Float64(data)
	if ok {
		return r
	} else {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		} else {
			return 0
		}
	}
}
