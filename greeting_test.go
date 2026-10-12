package sandbox

import "testing"

func TestGreeting(t *testing.T) {
	if got := Greeting(); got != "GSA sandbox" {
		t.Fatalf("Greeting() = %q", got)
	}
}

