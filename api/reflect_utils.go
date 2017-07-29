package utils

import (
	"reflect"
)

func GetFieldByName(object interface{}, name string) *reflect.Value {
	var foundField *reflect.Value
	objectType := reflect.TypeOf(object)
	objectValue := reflect.ValueOf(object)
	if objectValue.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
		objectValue = objectValue.Elem()
	}

	for i := 0; i < objectValue.NumField(); i++ {
		fieldType := objectType.Field(i)
		fieldValue := objectValue.Field(i)

		if fieldType.Name == name {
			foundField = &fieldValue
		} else if fieldType.Anonymous {
			foundField = GetFieldByName(fieldValue.Interface(), name)
		}

		if foundField != nil {
			break
		}
	}

	return foundField
}
