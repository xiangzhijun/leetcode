//[329]矩阵中的最长传递路径
//给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。
//
// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
//输出：4
//解释：最长递增路径为 [1, 2, 6, 9]。
//
// 示例 2：
//
//
//输入：matrix = [[3,4,5],[3,2,6],[2,2,1]]
//输出：4
//解释：最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
//
//
// 示例 3：
//
//
//输入：matrix = [[1]]
//输出：1
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// 0 <= matrix[i][j] <= 2³¹ - 1
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 拓扑排序 记忆化搜索 数组 动态规划 矩阵 👍 745 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n, n)
	}
	var dfs func(x, y, lastValue int) int
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	dfs = func(x, y, lastValue int) int {
		if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] <= lastValue {
			return 0
		}
		if cache[x][y] > 0 {
			return cache[x][y]
		}
		cache[x][y] = 1
		cache[x][y] = max(cache[x][y], 1+dfs(x-1, y, matrix[x][y]))
		cache[x][y] = max(cache[x][y], 1+dfs(x+1, y, matrix[x][y]))
		cache[x][y] = max(cache[x][y], 1+dfs(x, y-1, matrix[x][y]))
		cache[x][y] = max(cache[x][y], 1+dfs(x, y+1, matrix[x][y]))
		return cache[x][y]
	}

	rst := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rst = max(dfs(i, j, -1), rst)
		}
	}

	return rst
}

func main() {

}

//leetcode submit region end(Prohibit modification and deletion)
