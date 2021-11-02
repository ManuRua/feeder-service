package counter

import (
	"testing"
)

func TestCounter(t *testing.T) {
	counter := Counter{}

	v := counter.Value()
	if v != 0 {
		t.Errorf("Expected: %v, got: %v", 0, v)
	}

	counter.Inc()

	v = counter.Value()
	if v != 1 {
		t.Errorf("Expected: %v, got: %v", 1, v)
	}

	counter.Dec()

	v = counter.Value()
	if v != 0 {
		t.Errorf("Expected: %v, got: %v", 0, v)
	}
}
