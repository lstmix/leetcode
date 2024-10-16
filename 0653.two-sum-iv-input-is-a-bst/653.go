/**
 * Given the root of a Binary Search Tree and a target number k, return true if there exist two elements in the BST such that their sum is equal to the given target.
 *
 * Example 1:
 * (https://assets.leetcode.com/uploads/2020/09/21/sum_tree_1.jpg)
 * Input: root = [5,3,6,2,4,null,7], k = 9
 * Output: true
 *
 * Example 2:
 * (https://assets.leetcode.com/uploads/2020/09/21/sum_tree_2.jpg)
 * Input: root = [5,3,6,2,4,null,7], k = 28
 * Output: false
 *
 * Example 3:
 *
 * Input: root = [2,1,3], k = 4
 * Output: true
 *
 * Example 4:
 *
 * Input: root = [2,1,3], k = 1
 * Output: false
 *
 * Example 5:
 *
 * Input: root = [2,1,3], k = 3
 * Output: true
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [1, 10^4].
 * 	-10^4 <= Node.val <= 10^4
 * 	root is guaranteed to be a valid binary search tree.
 * 	-10^5 <= k <= 10^5
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkInMap(root *TreeNode, target int, m map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := m[target-root.Val]; ok {
		return true
	}
	m[root.Val] = struct{}{}
	return checkInMap(root.Left, target, m) || checkInMap(root.Right, target, m)
}

func findTarget(root *TreeNode, k int) bool {
	return checkInMap(root, k, make(map[int]struct{}))
}
