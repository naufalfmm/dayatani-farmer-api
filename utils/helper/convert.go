package helper

import (
	"reflect"
	"strconv"
)

func ConvertStringDefault[T any](data string, def T) T {
	var res any
	switch reflect.TypeOf(def).Kind() {
	case reflect.Int:
		i, err := strconv.Atoi(data)
		if err != nil {
			return def
		}

		res = i
	case reflect.Int16:
		i, err := strconv.ParseInt(data, 10, 16)
		if err != nil {
			return def
		}

		res = int16(i)
	case reflect.Int32:
		i, err := strconv.ParseInt(data, 10, 32)
		if err != nil {
			return def
		}

		res = int32(i)
	case reflect.Int64:
		i, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return def
		}

		res = int64(i)
	}

	return res.(T)
}
