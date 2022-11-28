# 给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
#
#
#
#  示例 1：
#
#
#
#
# 输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
# 输出：15
#
#
#  示例 2：
#
#
# 输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
# 输出：19
#
#
#
#
#  提示：
#
#
#  树中节点数目在范围 [1, 10⁴] 之间。
#  1 <= Node.val <= 100
#
#
#  Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 129 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def deepestLeavesSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """

        queue = [root, True]
        sum = 0
        while queue:
            node = queue.pop(0)
            if isinstance(node, bool):
                if len(queue) > 0:
                    sum = 0
                    queue.append(True)
                continue

            if node.left: queue.append(node.left)
            if node.right: queue.append(node.right)
            sum += node.val

        return sum
# leetcode submit region end(Prohibit modification and deletion)
