//[378]有序矩阵中第K小的元素
//给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。 请注意，它是 排序后 的第 k 小元素，而不是第
//k 个 不同 的元素。
//
// 你必须找到一个内存复杂度优于 O(n²) 的解决方案。
//
//
//
// 示例 1：
//
//
//输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
//[[ 1, 5, 9]
// [10,11,13]
// [12,13,15]]
//输出：13
//解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13
//
//
// 示例 2：
//
//
//输入：matrix = [[-5]], k = 1
//输出：-5
//
//
//
//
// 提示：
//
//
// n == matrix.length
// n == matrix[i].length
// 1 <= n <= 300
// -10⁹ <= matrix[i][j] <= 10⁹
// 题目数据 保证 matrix 中的所有行和列都按 非递减顺序 排列
// 1 <= k <= n²
//
//
//
//
// 进阶：
//
//
// 你能否用一个恒定的内存(即 O(1) 内存复杂度)来解决这个问题?
// 你能在 O(n) 的时间复杂度下解决这个问题吗?这个方法对于面试来说可能太超前了，但是你会发现阅读这篇文章（ this paper ）很有趣。
//
//
// Related Topics 数组 二分查找 矩阵 排序 堆（优先队列） 👍 933 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func kthSmallest(matrix [][]int, k int) int {
	//循环遍历，时间复杂度O(n*n),空间复杂度O(n)
	//n := len(matrix)
	//index := make([]int, n)
	//rst := 0
	//for i := 0; i < k; i++ {
	//	min := 1<<63 - 1
	//	minIndex := 0
	//	for j := 0; j < n; j++ {
	//		if index[j] < n && matrix[j][index[j]] < min {
	//			min = matrix[j][index[j]]
	//			minIndex = j
	//		}
	//	}
	//	rst = min
	//	index[minIndex]++
	//}
	//return rst

	//二分查找
	n := len(matrix)
	check := func(mid int) bool {
		i, j := n-1, 0
		num := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= mid {
				num += i + 1
				j++
			} else {
				i--
			}
		}
		return num >= k
	}

	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := (right-left)/2 + left
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func main() {
	println(kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 8))
}

//leetcode submit region end(Prohibit modification and deletion)
