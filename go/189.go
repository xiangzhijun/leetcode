//[189]轮转数组
//给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//[1,2,3,4,5,6,7]
//[5,6,7,1,2,3,4]
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]
//
//
// 示例 2:
//
//
//输入：nums = [-1,-100,3,99], k = 2
//输出：[3,99,-1,-100]
//解释:
//向右轮转 1 步: [99,-1,-100,3]
//向右轮转 2 步: [3,99,-1,-100]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// 0 <= k <= 10⁵
//
//
//
//
// 进阶：
//
//
// 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
//
//
//
//
//
//
//
//
// Related Topics 数组 数学 双指针 👍 1669 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "encoding/json"

//最大公约数，O(log(n)+n)
//func rotate(nums []int, k int) {
//	n := len(nums)
//	k = k % n
//	if k <= 0 {
//		return
//	}
//
//	//求n和k的最大公约数，时间复杂的log(n)
//	x, y, r := n, k, n%k
//	for r != 0 {
//		x, y = y, r
//		r = x % y
//	}
//
//	//从0到最大公约数循环交换，时间复杂的n，空间复杂的1
//	for start := 0; start < y; start++ {
//		index, temp := start, nums[start]
//		//从start位置出发开始交换，最后到start位置结束
//		for true {
//			newIndex := (index + k) % n
//			index, temp, nums[newIndex] = newIndex, nums[newIndex], temp
//			if index == start {
//				break
//			}
//		}
//	}
//}

//count计数，时间复杂度O(n)，空间复杂的O(1)
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	count := 0
	for start := 0; start < n && count < n; start++ {
		index, temp := start, nums[start]
		//从start位置出发开始交换，最后到start位置结束
		for true {
			newIndex := (index + k) % n
			index, temp, nums[newIndex] = newIndex, nums[newIndex], temp
			count++
			if index == start {
				break
			}
		}
	}
}

func main() {
	test := func(nums []int, k int) {
		bytes, _ := json.Marshal(nums)
		println(string(bytes))
		rotate(nums, k)
		bytes, _ = json.Marshal(nums)
		println(string(bytes))
		println()
	}

	test([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	test([]int{-1, -100, 3, 99}, 2)
	test([]int{-1}, 0)
	test([]int{1, 2, 3, 4, 5, 6}, 4)
}

//leetcode submit region end(Prohibit modification and deletion)
