/**
 * Given the root of a binary tree with N nodes, each node in the tree has node.val coins, and there are N coins total.
 *
 * In one move, we may choose two adjacent nodes and move one coin from one node to another.  (The move may be from parent to child, or from child to parent.)
 *
 * Return the number of moves required to make every node have exactly one coin.
 *
 *
 *
 * <div>
 * Example 1:
 *
 * (https://assets.leetcode.com/uploads/2019/01/18/tree1.png)
 *
 *
 * Input: <span id="example-input-1-1">[3,0,0]</span>
 * Output: <span id="example-output-1">2</span>
 * Explanation: From the root of the tree, we move one coin to its left child, and one coin to its right child.
 *
 *
 * <div>
 * Example 2:
 *
 * (https://assets.leetcode.com/uploads/2019/01/18/tree2.png)
 *
 *
 * Input: <span id="example-input-2-1">[0,3,0]</span>
 * Output: <span id="example-output-2">3</span>
 * Explanation: From the left child of the root, we move two coins to the root [taking two moves].  Then, we move one coin from the root of the tree to the right child.
 *
 *
 * <div>
 * Example 3:
 *
 * (https://assets.leetcode.com/uploads/2019/01/18/tree3.png)
 *
 *
 * Input: <span id="example-input-3-1">[1,0,2]</span>
 * Output: <span id="example-output-3">2</span>
 *
 *
 * <div>
 * Example 4:
 *
 * (https://assets.leetcode.com/uploads/2019/01/18/tree4.png)
 *
 *
 * Input: <span id="example-input-4-1">[1,0,0,null,3]</span>
 * Output: <span id="example-output-4">4</span>
 *
 *
 *
 *
 * <span>Note:</span>
 *
 * <ol>
 * 	1<= N <= 100
 * 	0 <= node.val <= N
 * </ol>
 * </div>
 * </div>
 * </div>
 * </div>
 */

package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func traverse(root *TreeNode, moves *int) int {
	if root == nil {
		return 0
	}
	left := traverse(root.Left, moves)
	right := traverse(root.Right, moves)
	*moves += abs(left) + abs(right)

	return root.Val + left + right - 1
}

func distributeCoins(root *TreeNode) int {
	var moves int
	traverse(root, &moves)
	return moves
}
