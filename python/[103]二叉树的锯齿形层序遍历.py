# 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
#
#
#
#  示例 1：
#
#
# 输入：root = [3,9,20,null,null,15,7]
# 输出：[[3],[20,9],[15,7]]
#
#
#  示例 2：
#
#
# 输入：root = [1]
# 输出：[[1]]
#
#
#  示例 3：
#
#
# 输入：root = []
# 输出：[]
#
#
#
#
#  提示：
#
#
#  树中节点数目在范围 [0, 2000] 内
#  -100 <= Node.val <= 100
#
#
#  Related Topics 树 广度优先搜索 二叉树 👍 680 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def zigzagLevelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return []

        result = [[]]
        queue = [root, True]
        left_2_right = True
        while queue:
            # 根据当前方向不同，从队首或队尾取下一个node
            if left_2_right:
                node = queue.pop(0)
            else:
                node = queue.pop(-1)
            if isinstance(node, bool):
                if len(queue) > 0:
                    # 遇到标识符后代表本层遍历结束，调转遍历方向开始下一层遍历
                    left_2_right = not left_2_right
                    if left_2_right:
                        queue.append(left_2_right)
                    else:
                        queue.insert(0, left_2_right)
                    result.append([])
                continue

            if left_2_right:
                # 从左往右遍历时追加在队尾
                if node.left: queue.append(node.left)
                if node.right: queue.append(node.right)
            else:
                # 从右往左遍历时追加在队首
                if node.right: queue.insert(0, node.right)
                if node.left: queue.insert(0, node.left)

            result[-1].append(node.val)

        return result

# leetcode submit region end(Prohibit modification and deletion)
