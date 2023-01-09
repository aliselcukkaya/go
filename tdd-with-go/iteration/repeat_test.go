package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)
	expected := "aaaaa"

	if result != expected {
		t.Errorf("expected %q got %q result", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	//Output:aaaaa
}
