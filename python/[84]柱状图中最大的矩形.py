# 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
#
#  求在该柱状图中，能够勾勒出来的矩形的最大面积。
#
#
#
#  示例 1:
#
#
#
#
# 输入：heights = [2,1,5,6,2,3]
# 输出：10
# 解释：最大的矩形为图中红色区域，面积为 10
#
#
#  示例 2：
#
#
#
#
# 输入： heights = [2,4]
# 输出： 4
#
#
#
#  提示：
#
#
#  1 <= heights.length <=10⁵
#  0 <= heights[i] <= 10⁴
#
#
#  Related Topics 栈 数组 单调栈 👍 2103 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        heights.append(0)
        S = [0] * len(heights)
        stack = [[0, -1]]
        for i in range(len(heights)):
            while stack and stack[-1][0] > heights[i]:
                last_height, last_index = stack.pop()
                S[last_index] = last_height * (i - stack[-1][1] - 1)
            stack.append([heights[i], i])
        return max(S)


# print(Solution().largestRectangleArea([2, 1, 5, 6, 2, 3]))
# leetcode submit region end(Prohibit modification and deletion)
