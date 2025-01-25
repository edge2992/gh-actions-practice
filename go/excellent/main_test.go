package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	t.Run("Test Even", func(t *testing.T) {
		result := EvenOrOdd(2)
		if result != "Even" {
			t.Errorf("EventOrOdd(2) = %s; want Even", result)
		}
	})

	t.Run("Test Odd", func(t *testing.T) {
		result := EvenOrOdd(3)
		if result != "Odd" {
			t.Errorf("EventOrOdd(3) = %s; want Odd", result)
		}
	})
}
