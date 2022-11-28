# 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充
# 。
#
#
#
#
#
#
#
#  示例 1：
#
#
# 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O",
# "X","X"]]
# 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
# 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都
# 会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
#
#
#  示例 2：
#
#
# 输入：board = [["X"]]
# 输出：[["X"]]
#
#
#
#
#  提示：
#
#
#  m == board.length
#  n == board[i].length
#  1 <= m, n <= 200
#  board[i][j] 为 'X' 或 'O'
#
#
#  Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 851 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def solve(self, board):
        """
        :type board: List[List[str]]
        :rtype: None Do not return anything, modify board in-place instead.
        """

        m = len(board)
        n = len(board[0])

        def dfs(x, y):
            if x < 0 or x >= m or y < 0 or y >= n:
                return
            if board[x][y] == 'O':
                board[x][y] = 'Y'
                dfs(x - 1, y)
                dfs(x + 1, y)
                dfs(x, y - 1)
                dfs(x, y + 1)

        for i in range(m):
            if board[i][0] == 'O': dfs(i, 0)
            if board[i][n - 1] == 'O': dfs(i, n - 1)
        for j in range(n):
            if board[0][j] == 'O': dfs(0, j)
            if board[m - 1][j] == 'O': dfs(m - 1, j)
        for i in range(m):
            for j in range(n):
                if board[i][j] == 'Y':
                    board[i][j] = 'O'
                elif board[i][j] == 'O':
                    board[i][j] = 'X'

        return board


# print(Solution().solve([["X", "X", "X", "X"], ["X", "O", "O", "X"], ["X", "X", "O", "X"], ["X", "O", "X", "X"]]))
# print(Solution().solve([["O", "O", "O", "O", "X", "X"], ["O", "O", "O", "O", "O", "O"], ["O", "X", "O", "X", "O", "O"],
#                         ["O", "X", "O", "O", "X", "O"], ["O", "X", "O", "X", "O", "O"],
#                         ["O", "X", "O", "O", "O", "O"]]))

# leetcode submit region end(Prohibit modification and deletion)
