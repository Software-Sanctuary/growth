package mapper

import (
	"errors"
	"fmt"
	"github.com/gokhantamkoc/growth-framework/logging"
	"reflect"
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
