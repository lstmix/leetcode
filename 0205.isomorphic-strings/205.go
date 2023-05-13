/**
 * Given two strings s and t, determine if they are isomorphic.
 *
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 *
 * All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.
 *
 * Example 1:
 *
 *
 * Input: s = "egg", t = "add"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s = "foo", t = "bar"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: s = "paper", t = "title"
 * Output: true
 *
 * Note:<br />
 * You may assume both s and t have the same length.
 *
 */

package leetcode

// The idea is that we need to map a char to another one, for example, "egg" and "add", we need to
// constract the mapping 'e' -> 'a' and 'g' -> 'd'. Instead of directly mapping 'e' to 'a', another way
// is to mark them with same value, for example, 'e' -> 1, 'a'-> 1, and 'g' -> 2, 'd' -> 2, this works same.
func isIsomorphic(s string, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}
