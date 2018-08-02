package codewars

import "strings"

/*Count the number of Duplicates
Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that
occur more than once in the input string. The input string can be assumed to contain only alphabets (both uppercase and
lowercase) and numeric digits.

Example
"abcde" -> 0 # no characters repeats more than once
"aabbcde" -> 2 # 'a' and 'b'
"aabBcde" -> 2 # 'a' occurs twice and 'b' twice (bandB)
"indivisibility" -> 1 # 'i' occurs six times
"Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
"aA11" -> 2 # 'a' and '1'
"ABBA" -> 2 # 'A' and 'B' each occur twice*/

func DuplicateCount(s string) (c int) {
	h := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}

func duplicate_count1(s1 string) int {
	counter := make(map[string]int)

	for _, s := range s1 {
		sl := strings.ToLower(string(s))
		counter[sl] += 1
	}

	sum := 0
	for _, v := range counter {
		if v > 1 {
			sum += 1
		}
	}

	return sum
}
