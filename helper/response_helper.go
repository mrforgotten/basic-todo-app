package helper

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

func ToRes(inputs ...interface{}) gin.H {
	result := gin.H{}
	for _, input := range inputs {
		inputVal := reflect.ValueOf(input)
		inputType := inputVal.Type()

		// Handle error types separately
		if inputType.Kind() == reflect.Ptr && inputVal.IsNil() {
			continue
		}

		if inputType.Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			err, _ := input.(error)
			result["error"] = err.Error()
		} else {
			result[inputType.Name()] = input
		}
	}
	return result
}
