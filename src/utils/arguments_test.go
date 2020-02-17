package utils

import (
	"testing"
)

func TestDryRun(t *testing.T) {
	if appConfig.dryRun != false {
		t.Fatalf("dryRun must be false by default")
	}
}

func TestPushBuild(t *testing.T) {
	if appConfig.pushBuild != false {
		t.Fatalf("pushBuild must be false by default")
	}
}