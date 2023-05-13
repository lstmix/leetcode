/**
 * Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
 * Letters are case sensitive, for example, "Aa" is not considered a palindrome here.
 *
 * Example 1:
 *
 * Input: s = "abccccdd"
 * Output: 7
 * Explanation:
 * One longest palindrome that can be built is "dccaccd", whose length is 7.
 *
 * Example 2:
 *
 * Input: s = "a"
 * Output: 1
 *
 * Example 3:
 *
 * Input: s = "bb"
 * Output: 2
 *
 *
 * Constraints:
 *
 * 	1 <= s.length <= 2000
 * 	s consits of lower-case and/or upper-case English letters only.
 *
 */

package leetcode

func longestPalindrome(s string) int {
	m := make(map[byte]struct{})
	count := 0

	for i := range s {
		if _, ok := m[s[i]]; ok {
			count += 2
			delete(m, s[i])
		} else {
			m[s[i]] = struct{}{}
		}
	}
	if len(m) > 0 {
		count++
	}
	return count
}
