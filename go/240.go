//[240]搜索二维矩阵II
//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
//
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
//
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 5
//输出：true
//
//
// 示例 2：
//
//
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 20
//输出：false
//
//
//
//
// 提示：
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= n, m <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// 每行的所有元素从左到右升序排列
// 每列的所有元素从上到下升序排列
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 二分查找 分治 矩阵 👍 1195 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func searchMatrix(matrix [][]int, target int) bool {
	//m行，n列
	m, n := len(matrix), len(matrix[0])
	//第i行，第j列
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][0] > target || matrix[i][j] < target {
			//target不在第i行，向下移动一行
			i++
		} else if matrix[i][j] > target || matrix[m-1][j] < target {
			//target不在第j列，向左移动一列
			j--
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	println(searchMatrix(matrix, 5))
	println(searchMatrix(matrix, 20))
	println(searchMatrix(matrix, 16))
}

//leetcode submit region end(Prohibit modification and deletion)
