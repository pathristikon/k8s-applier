package utils

import (
	"reflect"
	"strings"
)


/**
Reflect globalConfig struct and remove arguments inside

If an argument is present in globalConfig struct, then we need to remove
that argument from arguments slice
*/
func escapeArgumentsAlreadyInConfig(arguments []string) []string{
	reflectAppConfig := reflect.ValueOf(&appConfig).Elem()

	for x := 0; x < reflectAppConfig.NumField(); x++ {
		for index, argument := range arguments {
			field := reflectAppConfig.Type().Field(x)

			parsedArg := strings.ToLower(strings.Join(strings.Split(argument[1:], "-"), ""))

			if parsedArg == strings.ToLower(field.Name) {
				return RemoveIndexFromSlice(arguments, index)
			}
		}
	}

	return arguments
}


/** Remove slice element by its index */
func RemoveIndexFromSlice(sliceVariable []string, index int) []string {
	return append(sliceVariable[:index], sliceVariable[index+1:]...)
}