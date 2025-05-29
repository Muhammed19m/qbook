package assert

import "testing"

func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("expected error, got: %s", err)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error, got: %s", err)
	}
}

func Len[T any](t *testing.T, sl []T, expectedLen int) {
	if len(sl) != expectedLen {
		t.Errorf("expected len %d, got len %d", expectedLen, len(sl))
	}
}
