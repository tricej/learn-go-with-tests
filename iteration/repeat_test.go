package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("*", 6)
	expected := "******"

	if expected != repeated {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}
