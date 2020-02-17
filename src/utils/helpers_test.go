package utils

import (
	"reflect"
	"testing"
)

func TestHelperEscapeArgumentsAlreadyInConfig(t *testing.T) {
	args := []string{"-dry-run", "another", "flags"}
	parsedArguments := escapeArgumentsAlreadyInConfig(args)
	expectedResult := []string{"another", "flags"}

	if !reflect.DeepEqual(parsedArguments, expectedResult) {
		t.Fatalf("Something went wrong on comparing Helper!")
	}
}
