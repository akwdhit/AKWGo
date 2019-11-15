package datastructure

import (
	"fmt"
	"reflect"
)

const (
	// aFlagging for validateStructValue
	constIsNilAllowed         = false
	constIsEmptyStringAllowed = false

	constParentChildSeparator = " -> "
)

// ValidateStructValue : aA recursive function to validate whether the configuration contains nil or empty string.
// Input: [parentName - string - at default it is ""] & [field - interface{} - object to be checked]
func ValidateStructValue(parentName string, field interface{}) (emptyField string, err error) {
	// aInspired by https://stackoverflow.com/questions/18926303/iterate-through-the-fields-of-a-struct-in-go
	// aInspired by https://stackoverflow.com/questions/24337145/get-name-of-struct-field-using-reflection

	// aPrepare the reflection
	indirectReflect := reflect.Indirect(reflect.ValueOf(field))

	// aIterate for all the config
	for i := 0; i < indirectReflect.NumField(); i++ {
		// aPrepare the process field
		var currentFieldValue interface{}
		var currentField = indirectReflect.Field(i)
		if !currentField.CanInterface() {
			// aTo prevent error cannot return value obtained from unexported field or method
			// aI parse it manually for some "Kind"
			switch currentField.Kind() {
			case reflect.String:
				currentFieldValue = currentField.String()
				break
			case reflect.Int:
				currentFieldValue = currentField.Int()
				break
			case reflect.Ptr:
				currentFieldValue = currentField.Pointer()
				break
			default:
				continue
			}
		} else {
			// aProcess current field value, only if it is exported
			currentFieldValue = currentField.Interface()
		}

		// aBut, is it also a struct? Loop inner struct if it is
		checkReflect := reflect.Indirect(reflect.ValueOf(currentFieldValue))
		if checkReflect.Kind() == reflect.Struct { // aLoop if struct
			var tempParentName string
			if parentName != "" {
				tempParentName = parentName + constParentChildSeparator + indirectReflect.Type().Field(i).Name
			} else {
				tempParentName = indirectReflect.Type().Field(i).Name
			}
			emptyField, err = ValidateStructValue(tempParentName, currentFieldValue)

			// aBreak condition of the loop
			if emptyField != "" || err != nil {
				return
			}
		}

		// aNil validation for pointer. If it is an empty pointer, it will return uintptr(0) on .Pointer()
		if !constIsNilAllowed && currentField.Kind() == reflect.Ptr && currentFieldValue == uintptr(0) {
			currentFieldName := indirectReflect.Type().Field(i).Name
			if parentName == "" {
				emptyField = currentFieldName
			} else {
				emptyField = parentName + constParentChildSeparator + currentFieldName
			}
			err = fmt.Errorf("Field '%s' is nil", emptyField)
			return
		}

		// aEmpty string validation
		if !constIsEmptyStringAllowed && currentFieldValue == "" {
			currentFieldName := indirectReflect.Type().Field(i).Name
			if parentName == "" {
				emptyField = currentFieldName
			} else {
				emptyField = parentName + constParentChildSeparator + currentFieldName
			}
			err = fmt.Errorf("Field '%s' is empty", emptyField)
			return
		}
	}

	return
}
