//[300]最长递增子序列
//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。
//
// 示例 1：
//
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
//
// 示例 3：
//
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
//
// Related Topics 数组 二分查找 动态规划 👍 2947 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	max := 0

	//动态规划
	//dp[i]表示以nums[i]结尾的最长递增子序列长度，
	//dp[i]=1+max(dp[nums[j]<nums[i]]),j<i
	//for i := 0; i < n; i++ {
	//	for j := i - 1; j >= 0 && j >= dp[i]; j-- {
	//		if nums[j] < nums[i] && dp[j] > dp[i] {
	//			dp[i] = dp[j]
	//		}
	//	}
	//
	//	dp[i] += 1
	//	if dp[i] > max {
	//		max = dp[i]
	//	}
	//}

	//动态规划+二分查找
	//dp[i]表示长度为i的递增子序列的最后一个值
	//比如nums=[0, 1, 0, 3, 2, 3],dp=[0,1,2,3]
	for i := 0; i < n; i++ {
		left, right := 0, max
		for left < right {
			mid := (right + left) >> 1
			if dp[mid] < nums[i] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		dp[left] = nums[i]
		if max == right {
			max++
		}
	}

	return max
}

func main() {
	println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}

//leetcode submit region end(Prohibit modification and deletion)
