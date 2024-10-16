/**
 * In a deck of cards, each card has an integer written on it.
 * Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of cards, where:
 *
 * 	Each group has exactly X cards.
 * 	All the cards in each group have the same integer.
 *
 *
 * Example 1:
 *
 * Input: deck = [1,2,3,4,4,3,2,1]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].
 *
 * Example 2:
 *
 * Input: deck = [1,1,1,2,2,2,3,3]
 * Output: false&acute;
 * Explanation: No possible partition.
 *
 * Example 3:
 *
 * Input: deck = [1]
 * Output: false
 * Explanation: No possible partition.
 *
 * Example 4:
 *
 * Input: deck = [1,1]
 * Output: true
 * Explanation: Possible partition [1,1].
 *
 * Example 5:
 *
 * Input: deck = [1,1,2,2,2,2]
 * Output: true
 * Explanation: Possible partition [1,1],[2,2],[2,2].
 *
 *
 * Constraints:
 *
 * 	1 <= deck.length <= 10^4
 * 	0 <= deck[i] < 10^4
 *
 */

package leetcode

func gcd(a, b int) int {
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}

func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	for _, num := range deck {
		m[num]++
	}

	res := 0
	for _, count := range m {
		res = gcd(count, res)
	}
	return res > 1
}
