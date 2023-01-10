//[230]二叉搜索树中第K小的元素
//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,1,4,null,2], k = 1
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3
//
//
//
//
//
//
// 提示：
//
//
// 树中的节点数为 n 。
// 1 <= k <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
//
//
//
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 692 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	rst := 0
	//深度优先遍历，先找到最左节点，即最小值，然后开始回溯计数
	var dfs func(node *TreeNode, count int) int
	dfs = func(node *TreeNode, count int) int {
		if node == nil {
			return count
		}
		count = dfs(node.Left, count)
		if count == k {
			return count
		}
		count++
		if count == k {
			rst = node.Val
			return count
		}
		count = dfs(node.Right, count)
		return count
	}
	dfs(root, 0)

	return rst
}

func main() {

}

//leetcode submit region end(Prohibit modification and deletion)
