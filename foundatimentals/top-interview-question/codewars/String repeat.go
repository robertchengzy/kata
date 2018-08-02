package codewars

import "strings"

/*Write a function called repeatStr which repeats the given string string exactly n times.

repeatStr(6, "I") // "IIIIII"
repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"*/

func RepeatStr(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}
