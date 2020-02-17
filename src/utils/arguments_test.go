package utils

import (
	"testing"
)

func TestDryRun(t *testing.T) {
	if appConfig.dryRun != false {
		t.Fatalf("dryRun must be false by default")
	}
}
