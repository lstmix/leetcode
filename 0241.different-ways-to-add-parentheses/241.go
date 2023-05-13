/**
 * Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. The valid operators are +, - and *.
 *
 * Example 1:
 *
 *
 * Input: "2-1-1"
 * Output: [0, 2]
 * Explanation:
 * ((2-1)-1) = 0
 * (2-(1-1)) = 2
 *
 * Example 2:
 *
 *
 * Input: "2*3-4*5"
 * Output: [-34, -14, -10, -10, 10]
 * Explanation:
 * (2*(3-(4*5))) = -34
 * ((2*3)-(4*5)) = -14
 * ((2*(3-4))*5) = -10
 * (2*((3-4)*5)) = -10
 * (((2*3)-4)*5) = 10
 *
 */

package leetcode

func diffWays(input string, cache map[string][]int) []int {
	res := make([]int, 0)

	for i := range input {
		char := input[i]

		if char == '+' || char == '-' || char == '*' {
			left := input[:i]
			right := input[i+1:]

			leftRes, ok := cache[left]
			if !ok {
				leftRes = diffWays(left, cache)
			}

			rightRes, ok := cache[right]
			if !ok {
				rightRes = diffWays(right, cache)
			}

			for _, leftInt := range leftRes {
				for _, rightInt := range rightRes {
					r := 0
					switch char {
					case '+':
						r = leftInt + rightInt
					case '-':
						r = leftInt - rightInt
					case '*':
						r = leftInt * rightInt
					}
					res = append(res, r)
				}
			}
		}
	}

	// base case of recursion, we passed the for loop
	// and didn't compute anything, because there was only one part, left or right
	if len(res) == 0 {
		val := 0
		// we know for sure only integers are left
		for i := range input {
			val *= 10
			val += int(input[i] - '0')
		}

		res = append(res, val)
	}
	cache[input] = res

	return res
}

func diffWaysToCompute(input string) []int {
	cache := make(map[string][]int)
	return diffWays(input, cache)
}
