package iota_usage

import "testing"

func TestGetValueOfOperation(t *testing.T) {
	t.Run("Create is equal to 1", func(t *testing.T) {
		val := GetValueOfOperation(Create)
		if val != 1 {
			t.Errorf("Expected: 1, Actual: %d", val)
		}
	})
	t.Run("Update is equal to 2", func(t *testing.T) {
		val := GetValueOfOperation(Update)
		if val != 2 {
			t.Errorf("Expected: 2, Actual: %d", val)
		}
	})
}
