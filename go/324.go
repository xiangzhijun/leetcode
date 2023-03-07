//[324]摆动排序II
//给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
//
// 你可以假设所有输入数组都可以得到满足题目要求的结果。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,5,1,1,6,4]
//输出：[1,6,1,5,1,4]
//解释：[1,4,1,5,1,6] 同样是符合题目要求的结果，可以被判题程序接受。
//
//
// 示例 2：
//
//
//输入：nums = [1,3,2,2,3,1]
//输出：[2,3,1,3,1,2]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 5 * 10⁴
// 0 <= nums[i] <= 5000
// 题目数据保证，对于给定的输入 nums ，总能产生满足题目要求的结果
//
//
//
//
// 进阶：你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
//
// Related Topics 数组 分治 快速选择 排序 👍 528 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"encoding/json"
	"math/rand"
)

func wiggleSort(nums []int) {
	n := len(nums)
	mid := (n+1)/2 - 1
	target := findK(nums, 0, n-1, mid)

	//把mid左右两边等于target的值移动到mid周围
	for k, i, j := 0, 0, n-1; k <= j; k++ {
		//小于target的都在target左边，大于target的都在target右边
		//从左往右开始循环
		if nums[k] < target {
			//由于小于target的都在target(mid)左边，所以只有mid前面的循环可能进这个if
			//这个swap是为了把等于target的值交换到紧邻target左侧
			swap(nums, k, i)
			i++
		}
		if nums[k] > target {
			//由于大于target的都在target右边，所以只有mid后面的循环可能进这个if
			for j > k && nums[j] > target {
				j--
			}
			//这个swap是为了把等于target的值交换到紧邻target右侧
			swap(nums, k, j)
			j--
		}
	}

	//现在nums已经被分为三个分区，依次为小于target的、等于target的、大于target的
	temp := make([]int, n, n)
	copy(temp, nums)
	for i, j, k := 0, mid, n-1; i < n; i, j, k = i+2, j-1, k-1 {
		nums[i] = temp[j]
		if i+1 < n {
			nums[i+1] = temp[k]
		}
	}
}

func findK(nums []int, l, r, k int) int {
	randIndex := rand.Intn(r-l+1) + l
	swap(nums, randIndex, l)

	x := nums[l]
	i := r
	for j := r; j > l; j-- {
		if nums[j] > x {
			swap(nums, i, j)
			i--
		}
	}
	swap(nums, i, l)

	if i == k {
		return nums[i]
	}
	if i < k {
		return findK(nums, i+1, r, k)
	} else {
		return findK(nums, l, i-1, k)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func test(nums []int) {
	bytes, _ := json.Marshal(nums)
	println(string(bytes))
	wiggleSort(nums)
	bytes, _ = json.Marshal(nums)
	println(string(bytes))
	println()
}

func main() {
	test([]int{3, 5, 1, 2, 6, 3, 3})
	//wiggleSort([]int{4, 2, 5, 6, 3, 1})

}

//leetcode submit region end(Prohibit modification and deletion)
