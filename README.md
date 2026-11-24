# Longest Palindromic Substring

> Striver SDE Sheet 093 · Day 15 String — every palindrome has a center (character or gap): expand all 2n−1 of them and keep the widest. Go, stdlib only, tested.

## Problem

Find the longest contiguous palindrome inside a string. `"babad"` → `"bab"` (or equally `"aba"`).

## Approach

Instead of testing substrings (O(n³)) or filling a DP table (O(n²) *space*), enumerate palindromes by their centers: a palindrome is symmetric around either a character (odd length) or the gap between two (even length) — 2n−1 centers in all. From each, grow outward while the ends match; matching is O(1) and each expansion step is paid for by palindrome length actually found. The even/odd split is the detail that loses people: missing the gap centers silently drops `"bb"`-style answers. Manacher's algorithm does O(n) by reusing symmetry between centers — the named upgrade to mention, rarely worth implementing live.

**Complexity:** O(n²) worst case, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
