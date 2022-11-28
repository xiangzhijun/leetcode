# 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
# 返回其根节点。
#
#
#
#  示例 1:
#
#
# 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
# 输出: [3,9,20,null,null,15,7]
#
#
#  示例 2:
#
#
# 输入: preorder = [-1], inorder = [-1]
# 输出: [-1]
#
#
#
#
#  提示:
#
#
#  1 <= preorder.length <= 3000
#  inorder.length == preorder.length
#  -3000 <= preorder[i], inorder[i] <= 3000
#  preorder 和 inorder 均 无重复 元素
#  inorder 均出现在 preorder
#  preorder 保证 为二叉树的前序遍历序列
#  inorder 保证 为二叉树的中序遍历序列
#
#
#  Related Topics 树 数组 哈希表 分治 二叉树 👍 1695 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def buildTree(self, preorder, inorder):
        """
        :type preorder: List[int]
        :type inorder: List[int]
        :rtype: TreeNode
        """

        index_map = {inorder[i]: i for i in range(len(inorder))}

        def loop(pre_range, in_range):
            root = TreeNode(preorder[pre_range[0]])
            root_index = index_map[root.val]

            left_size = root_index - in_range[0]
            right_size = in_range[1] - root_index - 1

            if left_size > 0:
                root.left = loop([pre_range[0] + 1, pre_range[0] + left_size + 1],
                                 [in_range[0], in_range[0] + left_size])

            if right_size > 0:
                root.right = loop([pre_range[0] + 1 + left_size, pre_range[1]],
                                  [in_range[0] + left_size + 1, in_range[1]])

            return root

        return loop([0, len(preorder)], [0, len(inorder)])

# leetcode submit region end(Prohibit modification and deletion)
