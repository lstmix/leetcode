/**
 * Given an unsorted array, find the maximum difference between the successive elements in its sorted form.
 *
 * Return 0 if the array contains less than 2 elements.
 *
 * Example 1:
 *
 *
 * Input: [3,6,9,1]
 * Output: 3
 * Explanation: The sorted form of the array is [1,3,6,9], either
 *              (3,6) or (6,9) has the maximum difference 3.
 *
 * Example 2:
 *
 *
 * Input: [10]
 * Output: 0
 * Explanation: The array contains less than 2 elements, therefore return 0.
 *
 * Note:
 *
 *
 * 	You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
 * 	Try to solve it in linear time/space.
 *
 *
 */

package leetcode

import "math"

// Method 1. Radix sort
func countSort(nums, aux []int, exp int) {
	n := len(nums)

	// radix is 10, could be something else
	count := make([]int, 10)

	// store count of occurrences of current digit
	for i := 0; i < n; i++ {
		count[(nums[i]/exp)%10]++
	}

	// modify count so it contains actual position
	// of this digit in sorted array (aux)
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// build output array
	for i := n - 1; i >= 0; i-- {
		// get the actual position and put a number there
		aux[count[(nums[i]/exp)%10]-1] = nums[i]
		// decrease position
		count[(nums[i]/exp)%10]--
	}

	copy(nums, aux)
}

func radixSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	// find the maximum number to know number of digits
	maxVal := nums[0]
	for i := range nums {
		maxVal = max(maxVal, nums[i])
	}

	aux := make([]int, len(nums))
	// exp is 10^i where i is current digit number
	// 1, 10, 100, 1000 ...
	for exp := 1; maxVal/exp > 0; exp *= 10 {
		countSort(nums, aux, exp)
	}
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	radixSort(nums)

	gap := 0
	for i := 1; i < len(nums); i++ {
		gap = max(gap, nums[i]-nums[i-1])
	}
	return gap
}

// Method 2. Bucket sort
func maximumGap2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	minVal, maxVal := nums[0], nums[0]
	for i := range nums {
		minVal = min(minVal, nums[i])
		maxVal = max(maxVal, nums[i])
	}

	if minVal == maxVal {
		return 0
	}

	// set the buckets to be smaller than (max−min)/(n−1)
	// Since the gaps (between elements) within the same bucket would only be ≤t,
	// we could deduce that the maximal gap would indeed occur
	// only between two adjacent buckets

	gap := int(math.Ceil(float64(maxVal-minVal) / float64(n-1)))
	// n buckets, it could be pairs, but two arrays is simpler
	bucketMin := make([]int, n)
	bucketMax := make([]int, n)
	for i := range bucketMin {
		bucketMin[i] = math.MaxInt64
	}

	// default values
	for i := 0; i < n; i++ {
		bucketMin[i] = math.MaxInt64
		bucketMax[i] = math.MinInt64
	}

	for _, num := range nums {
		// find the bucket
		idx := (num - minVal) / gap
		// update bucket min and max
		bucketMin[idx] = min(bucketMin[idx], num)
		bucketMax[idx] = max(bucketMax[idx], num)
	}

	for i := 0; i < n; i++ {
		// find the largest gap
		if bucketMin[i] != math.MaxInt64 {
			gap = max(gap, bucketMin[i]-minVal)
			minVal = bucketMax[i]
		}
	}
	return gap
}
