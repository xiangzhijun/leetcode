# 给你一个二叉树的根节点 root ， 检查它是否轴对称。
#
#
#
#  示例 1：
#
#
# 输入：root = [1,2,2,3,4,4,3]
# 输出：true
#
#
#  示例 2：
#
#
# 输入：root = [1,2,2,null,3,null,3]
# 输出：false
#
#
#
#
#  提示：
#
#
#  树中节点数目在范围 [1, 1000] 内
#  -100 <= Node.val <= 100
#
#
#
#
#  进阶：你可以运用递归和迭代两种方法解决这个问题吗？
#
#  Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2065 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """

        def dfs(left_node, right_node):
            if left_node is None and right_node is None:
                return True
            if left_node is None or right_node is None:
                return False

            if left_node.val != right_node.val:
                return False

            if not bfs(left_node.left, right_node.right):
                return False
            if not bfs(left_node.right, right_node.left):
                return False

            return True

        return dfs(root.left, root.right)

# leetcode submit region end(Prohibit modification and deletion)
