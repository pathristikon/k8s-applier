package utils

import "testing"

func TestAlerts(t *testing.T) {
	validAlertTypes := []string{"ERR", "NOTICE", "WARNING"}
	for _, t := range validAlertTypes {
		Alert(t, "Test", true)
	}
}

func TestAlertsPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Something went wrong, not panicked!")
		}
	}()

	Alert("TestingPanic", "Test", true)
}