/**
 * Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.
 *
 *
 * Example 1:<br />
 *
 * Input: [0,1]
 * Output: 2
 * Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
 *
 *
 *
 * Example 2:<br />
 *
 * Input: [0,1,0]
 * Output: 2
 * Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
 *
 *
 *
 * Note:
 * The length of the given binary array will not exceed 50,000.
 *
 */

package leetcode

func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1

	var maxLen, count int
	for i := range nums {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}

		if prevIdx, ok := m[count]; ok {
			maxLen = max(maxLen, i-prevIdx)
		} else {
			m[count] = i
		}
	}
	return maxLen
}
