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

func TestWrap(t *testing.T) {
	fmt.Println(Wrap([]string{"a","b","c"},"'"))
}


func TestWrapString(t *testing.T) {
	fmt.Println(WrapString("a,b,c,",",","'"))
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

func TestCut(t *testing.T) {
	s1 := "如果List中的每个元素作为Pred函数的参数执行，结果都返回true，那么all函数返回true，"
	s2 := "如果List中的每个元素"
	s3 := "如果List中的每个元素作"
	fmt.Println(Cut(s1,12,"..."))
	fmt.Println(Cut(s2,12,"..."))
	fmt.Println(Cut(s3,12,".."))
}

func BenchmarkLeftPadInt(b *testing.B) {
	for i:=0;i<b.N;i++{
		LeftPadInt(12345,10)
	}
}