package mstring

import (
	"fmt"
	"testing"
)

func TestIsBlank(t *testing.T) {
	s := "  "
	if !IsBlank(s) {
		t.Error(s)
	}
}

func TestRandomString(t *testing.T) {
	fmt.Println(Random(0))
	fmt.Println(Random(10))
}

func TestReverse(t *testing.T) {
	fmt.Println(Reverse("fjkewo3i232klfes"))
}

func TestLeftPad(t *testing.T) {
	fmt.Println(LeftPadInt(10,10))
	fmt.Println(LeftPadInt(1,10))
	fmt.Println(LeftPadInt(122121,10))
	fmt.Println(LeftPadInt(12345678919,10))
	fmt.Println(LeftPadInt(10,10))
}

func BenchmarkLeftPadInt(b *testing.B) {
	for i:=0;i<b.N;i++{
		LeftPadInt(12345,10)
	}
}