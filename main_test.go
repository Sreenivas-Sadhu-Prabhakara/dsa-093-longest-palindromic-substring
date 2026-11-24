package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"babad", "bab"}, // "aba" equally long; first found wins
		{"cbbd", "bb"},
		{"a", "a"},
		{"", ""},
		{"ac", "a"},
		{"racecar", "racecar"},
		{"aaaa", "aaaa"},
		{"abacdfgdcaba", "aba"},
	}
	for _, c := range cases {
		if got := LongestPalindrome(c.in); got != c.want {
			t.Errorf("LongestPalindrome(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
