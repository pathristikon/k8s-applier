package utils

import "testing"

func TestConstants(t *testing.T) {
	if ConfigDir == "" ||
		ConfigFile == "" ||
		ProjectName == "" ||
		KubectlArgument == "" ||
		HelmArgument == "" {
		t.Fatalf("Constants missing in constants.go!")
	}
}
