package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

//Practice: Caller can now specify how many times it should repeat the character
//Added another test case for fun
func TestRepeat(t *testing.T) {
	repeated_a := Repeat("a", 5)
	expected_a := "aaaaa"

	if repeated_a != expected_a {
		t.Errorf("Test 1 Failed: expected %q but got %q", expected_a, repeated_a)
	}

	repeated_b := Repeat("b", 4)
	expected_b := "bbbb"

	if repeated_b != expected_b {
		t.Errorf("Test 2 Failed: expected %q but got %q", expected_b, repeated_b)
	}
}