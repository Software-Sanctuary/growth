package mapper

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fatih/structs"
	"github.com/gokhantamkoc/growth-framework/logging"
)

func ReadTag(obj interface{}, fieldName, tagName string) (string, error) {
	logger := &logging.ConsoleLogger{}
	logger.NewConsoleLogger("mapper:ReadTag")
	field, hasField := reflect.TypeOf(obj).Elem().FieldByName(fieldName)
	if !hasField {
		errMessage := fmt.Sprintf(
			"cannot find field with name as %s",
			fieldName,
		)
		logger.Error(errMessage)
		return "", errors.New(errMessage)
	}
	fieldTag, hasTag := field.Tag.Lookup(tagName)
	if !hasTag {
		errMessage := fmt.Sprintf(
			"cannot find tag with name as %s",
			fieldName,
		)
		logger.Error(errMessage)
		return "", errors.New(errMessage)
	}
	return fieldTag, nil
}

func MapObject(objFrom interface{}, objTo interface{}) {
	logger := &logging.ConsoleLogger{}
	logger.NewConsoleLogger("mapper:Map")
	mapObjFrom := structs.Map(objFrom)
	for key, val := range mapObjFrom {
		if _, hasField := reflect.TypeOf(objTo).FieldByName(key); hasField {
			newValue := reflect.New(reflect.TypeOf(val))
			newValue.Elem().Set(reflect.ValueOf(val))
			reflect.ValueOf(objTo).FieldByName(key).Set(reflect.ValueOf(newValue))
		}
	}

	// objFromType := reflect.TypeOf(objFrom)
	// for fieldIndex := 0; fieldIndex < objFromType.NumField(); fieldIndex++ {
	//	fieldNameToMap := objFromType.Field(fieldIndex).Name
	//	if field, hasField := reflect.TypeOf(objTo).FieldByName(fieldNameToMap); hasField {
	//		fieldValue := reflect.ValueOf(field)
	//		newValue := reflect.New(reflect.TypeOf(objFromType.Field(fieldIndex)))
	//		newValue.Elem().Set(fieldValue)
	//		reflect.(&objTo).Elem().FieldByName(fieldNameToMap).Set(newValue)
	//	}
	// }
}
