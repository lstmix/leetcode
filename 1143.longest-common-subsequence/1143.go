/**
 * Given two strings text1 and text2, return the length of their longest common subsequence.
 * A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.
 *
 * If there is no common subsequence, return 0.
 *
 * Example 1:
 *
 * Input: text1 = "abcde", text2 = "ace"
 * Output: 3
 * Explanation: The longest common subsequence is "ace" and its length is 3.
 *
 * Example 2:
 *
 * Input: text1 = "abc", text2 = "abc"
 * Output: 3
 * Explanation: The longest common subsequence is "abc" and its length is 3.
 *
 * Example 3:
 *
 * Input: text1 = "abc", text2 = "def"
 * Output: 0
 * Explanation: There is no such common subsequence, so the result is 0.
 *
 *
 * Constraints:
 *
 * 	1 <= text1.length <= 1000
 * 	1 <= text2.length <= 1000
 * 	The input strings consist of lowercase English characters only.
 *
 */

package leetcode

// 2D Dynamic programming solution
func longestCommonSubsequence(text1 string, text2 string) int {
	T1 := []byte(text1)
	T2 := []byte(text2)
	LT1 := len(T1)
	LT2 := len(T2)

	memo := make([][]int, LT1+1)
	for i := 0; i <= LT1; i++ {
		memo[i] = make([]int, LT2+1)
	}

	for i := 1; i <= LT1; i++ {
		for j := 1; j <= LT2; j++ {
			if T1[i-1] == T2[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i][j-1])
			}
		}
	}

	return memo[LT1][LT2]
}
