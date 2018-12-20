package acdc

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	s := Sum(1, 2, 3)
	if s != 6 {
		t.Error("got", s, "expected", 6)
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3))
	// Output:
	// 6
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3)
	}
}
