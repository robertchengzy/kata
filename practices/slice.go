package main

import (
	"fmt"
	"unsafe"
)

func main11() {
	var s []int
	for i := 1; i <= 8; i++ {
		s = append(s, i)
	}

	fmt.Println(cap(s))
	reverse(s)
	fmt.Println(s)

	reverse2(s)
	fmt.Println(s)

	reverse3(s)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s []int) {
	s = append(s, 999)
	for i, j := 0, len(s) - 1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func reverse3(s []int) {
	s = append(s, 999, 1000, 1001)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func reverse4(s []int) {
	newElem := 999
	for len(s) < cap(s) {
		fmt.Println("Adding an element:", newElem, "cap:", cap(s), "len:", len(s))
		s = append(s, newElem)
		newElem++
	}
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

type A struct {
	Ptr1 *B
	Ptr2 *B
	Val B
}

type B struct {
	Str string
}

func main() {
	a := A{
		Ptr1: &B{"ptr-str-1"},
		Ptr2: &B{"ptr-str-2"},
		Val: B{"val-str"},
	}
	fmt.Println(a.Ptr1)
	fmt.Println(a.Ptr2)
	fmt.Println(a.Val)
	demo(a)
	fmt.Println(a.Ptr1)
	fmt.Println(a.Ptr2)
	fmt.Println(a.Val)
}

func demo(a A) {
	// Update a value of a pointer and changes will persist
	a.Ptr1.Str = "new-ptr-str1"
	// Use an entirely new B object and changes won't persist
	a.Ptr2 = &B{"new-ptr-str-2"}
	a.Val.Str = "new-val-str"
}

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}