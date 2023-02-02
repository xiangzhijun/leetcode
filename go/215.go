//[215]数组中的第K个最大元素
//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
//输入: [3,2,1,5,6,4], k = 2
//输出: 5
//
//
// 示例 2:
//
//
//输入: [3,2,3,1,2,4,5,5,6], k = 4
//输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 1996 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	//快速选择排序
	var quicklySelect func(left, right int) int
	quicklySelect = func(start, end int) int {
		if end > start {
			randIndex := start + rand.Intn(end-start)
			//随机选择一个值用于快速选择，防止快速排序最坏情况
			nums[start], nums[randIndex] = nums[randIndex], nums[start]
		}

		i := start
		v := nums[i]

		left, right := start, end
		for left < right {
			for left < right {
				if nums[right] < v {
					nums[right], nums[i], i = nums[i], nums[right], right
					//right--
					break
				}
				right--
			}

			for left < right {
				if nums[left] > v {
					nums[left], nums[i], i = nums[i], nums[left], left
					//left++
					break
				}
				left++
			}
		}

		if i == k {
			return nums[i]
		} else if i < k {
			return quicklySelect(i+1, end)
		} else {
			return quicklySelect(start, i-1)
		}
	}

	return quicklySelect(0, len(nums)-1)
}
func main() {
	println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

//leetcode submit region end(Prohibit modification and deletion)
