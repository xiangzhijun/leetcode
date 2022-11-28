# 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不
# 一定经过根节点。
#
#  路径和 是路径中各节点值的总和。
#
#  给你一个二叉树的根节点 root ，返回其 最大路径和 。
#
#
#
#  示例 1：
#
#
# 输入：root = [1,2,3]
# 输出：6
# 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
#
#  示例 2：
#
#
# 输入：root = [-10,9,20,null,null,15,7]
# 输出：42
# 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
#
#
#
#
#  提示：
#
#
#  树中节点数目范围是 [1, 3 * 10⁴]
#  -1000 <= Node.val <= 1000
#
#
#  Related Topics 树 深度优先搜索 动态规划 二叉树 👍 1682 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def maxPathSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        # min_val = float('-inf')

        # def dfs(node):
        #     if not node:
        #         return None
        #
        #     left_sum = dfs(node.left)
        #     right_sum = dfs(node.right)
        #
        #     not_conn_max = min_val
        #     left_conn, right_conn, all_conn = node.val, node.val, node.val
        #     if left_sum:
        #         if left_sum[0] > 0:
        #             left_conn += left_sum[0]
        #         not_conn_max = left_sum[1]
        #     if right_sum:
        #         if right_sum[0] > 0:
        #             right_conn += right_sum[0]
        #         not_conn_max = max(not_conn_max, right_sum[1])
        #
        #     not_conn_max = max(not_conn_max, left_conn + right_conn - node.val)
        #     conn_max = max(left_conn, right_conn)
        #
        #     return conn_max, not_conn_max
        #
        # max_tuple = dfs(root)
        # # print(max_tuple)
        # return max(max_tuple)
        self.max_sum = float('-inf')

        def dfs(node):
            if not node:
                return 0
            left_max = max(dfs(node.left), 0)
            right_max = max(dfs(node.right), 0)
            max_val = node.val + left_max + right_max

            if max_val > self.max_sum:
                self.max_sum = max_val
            return node.val + max(left_max, right_max)

        dfs(root)
        return self.max_sum

# leetcode submit region end(Prohibit modification and deletion)
