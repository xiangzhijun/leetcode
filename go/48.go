//[48]旋转图像
//给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：         [[7,4,1],[8,5,2],[9,6,3]]
// 1:(0,0)->(0,2) 2:(0,1)->(1,2) 3:(0,2)->(2,2)
// 4:(1,0)->(0,1) 5:(1,1)->(1,1) 6:(1,2)->(2,1)
// 7:(2,0)->(0,0)
//
//
// 示例 2：
//
//
//输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
//
//
//
//
// 提示：
//
//
// n == matrix.length == matrix[i].length
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000
//
//
//
//
// Related Topics 数组 数学 矩阵 👍 1569 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "encoding/json"

func rotate(matrix [][]int) {
	n := len(matrix)
	//for i := 0; i < n/2; i++ {
	//	for j := 0; j < (n+1)/2; j++ {
	//		matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
	//	}
	//}

	var l, r, temp int
	for i := 0; i < (n+1)/2; i++ {
		for j := i; j < n-i-1; j++ {
			l, r, temp = i, j, matrix[i][j]
			for true {
				l, r = r, n-l-1
				matrix[l][r], temp = temp, matrix[l][r]
				if l == i && r == j {
					break
				}
			}
		}
	}
}

func test(matrix [][]int) {
	bytes, _ := json.Marshal(matrix)
	println(string(bytes))
	rotate(matrix)
	bytes, _ = json.Marshal(matrix)
	println(string(bytes))
	println()
}

func main() {
	test([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}

//leetcode submit region end(Prohibit modification and deletion)
