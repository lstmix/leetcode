/**
 * Given the root of a binary tree, return the preorder traversal of its nodes' values.
 *
 * Example 1:
 * (https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg)
 * Input: root = [1,null,2,3]
 * Output: [1,2,3]
 *
 * Example 2:
 *
 * Input: root = []
 * Output: []
 *
 * Example 3:
 *
 * Input: root = [1]
 * Output: [1]
 *
 * Example 4:
 * (https://assets.leetcode.com/uploads/2020/09/15/inorder_5.jpg)
 * Input: root = [1,2]
 * Output: [1,2]
 *
 * Example 5:
 * (https://assets.leetcode.com/uploads/2020/09/15/inorder_4.jpg)
 * Input: root = [1,null,2]
 * Output: [1,2]
 *
 *
 * Constraints:
 *
 * 	The number of nodes in the tree is in the range [0, 100].
 * 	-100 <= Node.val <= 100
 *
 *
 * Follow up:
 * Recursive solution is trivial, could you do it iteratively?
 *
 *
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	storage []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.storage = append(s.storage, node)
}

func (s *Stack) Pop() *TreeNode {
	node := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	return node
}

func (s *Stack) Empty() bool {
	return len(s.storage) == 0
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := &Stack{}

	for root != nil {
		res = append(res, root.Val)
		if root.Right != nil {
			stack.Push(root.Right)
		}
		root = root.Left
		if root == nil && !stack.Empty() {
			root = stack.Pop()
		}
	}
	return res
}
