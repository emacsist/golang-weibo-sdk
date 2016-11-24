package helper

import "reflect"

// SetDefaultValues : 设置参数默认值
func SetDefaultValues(param interface{}) {
	v := reflect.ValueOf(param)
	fileds := v.Elem()
	for i := 0; i < fileds.NumField(); i++ {
		field := fileds.Field(i)
		// Ignore fields that don't have the same type as a string
		if field.Type().Name() == "string" {
			//如果是一个字符串对象
			if field.CanSet() {
				field.SetString("")
			}
		} else if field.Type().Name() == "int" || field.Type().Name() == "int64" || field.Type().Name() == "int32" || field.Type().Name() == "int8" || field.Type().Name() == "int16" {
			//如果是整型
			if field.CanSet() {
				field.SetInt(-1)
			}
		}
	}
}
