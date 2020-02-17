package utils

import "testing"

func TestConstants(t *testing.T) {
	if ConfigDir == "" ||
		ConfigFile == "" ||
		ProjectName == "" ||
		KubectlArgument == "" ||
		HelmArgument == "" ||
		K8sVersion == "" {
		t.Fatalf("Constants missing in constants.go!")
	}
}
