package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	fmt.Println("good good study, day day up")
	fmt.Println(HashCode("1234567889"))
}

// 获取hashcode
func HashCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
