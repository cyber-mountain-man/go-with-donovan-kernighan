package stringsutil

import "testing"

func TestReverse(t *testing.T) {
	got := Reverse("morrison")
	want := "nosirrom"

	if got != want {
		t.Fatalf("Reverse() = %q, want %q", got, want)
	}
}
