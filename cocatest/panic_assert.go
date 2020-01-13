package cocatest

import "testing"

func AssertPanic(t *testing.T, panicFunc func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	panicFunc()
}
