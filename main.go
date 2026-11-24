package main

import "fmt"

// LongestPalindrome returns the longest palindromic substring of s (the
// first one found at maximal length). Every palindrome has a center — a
// character or a gap — so expand around all 2n−1 centers.
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, maxLen := 0, 1
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > maxLen {
			start = l + 1
			maxLen = r - l - 1
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)   // odd-length palindromes centered on a character
		expand(i, i+1) // even-length palindromes centered on a gap
	}
	return s[start : start+maxLen]
}

func main() {
	fmt.Println(LongestPalindrome("babad"))
}
