# 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
#
#  在「杨辉三角」中，每个数是它左上方和右上方的数的和。
#
#
#
#
#
#  示例 1:
#
#
# 输入: numRows = 5
# 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
#
#
#  示例 2:
#
#
# 输入: numRows = 1
# 输出: [[1]]
#
#
#
#
#  提示:
#
#
#  1 <= numRows <= 30
#
#
#  Related Topics 数组 动态规划 👍 819 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        result = []
        for i in range(numRows):
            row = []
            pre_row = result[i - 1] if i > 0 else None
            for j in range(i + 1):
                if j == 0 or j == i:
                    row.append(1)
                else:
                    row.append(pre_row[j - 1] + pre_row[j])
            result.append(row)
        return result
# leetcode submit region end(Prohibit modification and deletion)
