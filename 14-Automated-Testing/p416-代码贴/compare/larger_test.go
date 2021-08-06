package compare

import (
	"fmt"
	"testing"
)

func TestFirstLarger(t *testing.T) {
	got := Larger(1, 2)
	want := 2
	if got != want {
		t.Error(helper(1, 2, got, want))
	}
}

func helper(a, b, got, want int) string {
	return fmt.Sprintf("Larger(%d, %d) = %d, want %d", a, b, got, want)
}
