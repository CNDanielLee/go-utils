package utils

import "reflect"

// CopyProperty 复制结构体中的相同字段，source传结构体，target传结构体指针
func CopyProperty(source interface{}, target interface{}) {
	if source != nil {
		var copyObjV = reflect.ValueOf(target)
		//反射获取属性的类型和值
		var objT = reflect.TypeOf(source)
		var objV = reflect.ValueOf(source)
		//对应设置的数值
		for i := 0; i < objT.NumField(); i++ {
			if objT.Field(i).Anonymous {
				continue
			}
			property := copyObjV.Elem().FieldByName(objT.Field(i).Name)
			if property.IsValid() {
				property.Set(objV.Field(i))
			}
		}
	}
}