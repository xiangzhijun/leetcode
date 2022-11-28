# 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
#
#  单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
#
#
#
#  示例 1：
#
#
# 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
# "ABCCED"
# 输出：true
#
#
#  示例 2：
#
#
# 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
# "SEE"
# 输出：true
#
#
#  示例 3：
#
#
# 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
# "ABCB"
# 输出：false
#
#
#
#
#  提示：
#
#
#  m == board.length
#  n = board[i].length
#  1 <= m, n <= 6
#  1 <= word.length <= 15
#  board 和 word 仅由大小写英文字母组成
#
#
#
#
#  进阶：你可以使用搜索剪枝的技术来优化解决方案，使其在 board 更大的情况下可以更快解决问题？
#
#  Related Topics 数组 回溯 矩阵 👍 1403 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def exist(self, board, word):
        """
        :type board: List[List[str]]
        :type word: str
        :rtype: bool
        """
        m = len(board)
        n = len(board[0])
        z = len(word)

        cache = [[False] * n for i in range(m)]

        def backtrace(x, y, index):
            if m > x >= 0 and n > y >= 0 and not cache[x][y] and board[x][y] == word[index]:
                cache[x][y] = True
                if index == z - 1:
                    return True

                if backtrace(x, y - 1, index + 1): return True
                if backtrace(x, y + 1, index + 1): return True
                if backtrace(x + 1, y, index + 1): return True
                if backtrace(x - 1, y, index + 1): return True
                cache[x][y] = False

            return False

        for i in range(m):
            for j in range(n):
                if backtrace(i, j, 0):
                    return True

        return False


print(Solution().exist(board=[["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]], word="SEE"))

# leetcode submit region end(Prohibit modification and deletion)
