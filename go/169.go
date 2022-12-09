//[169]多数元素
//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1：
//
//
//输入：nums = [3,2,3]
//输出：3
//
// 示例 2：
//
//
//输入：nums = [2,2,1,1,1,2,2]
//输出：2
//
//
//
//提示：
//
//
// n == nums.length
// 1 <= n <= 5 * 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
//
// Related Topics 数组 哈希表 分治 计数 排序 👍 1632 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func majorityElement(nums []int) int {
	//Boyer-Moore 投票算法
	candidate := nums[0]
	count := 0
	for _, x := range nums {
		if count == 0 {
			candidate = x
		}
		if candidate == x {
			count += 1
		} else {
			count -= 1
		}
	}
	return candidate
}

func main() {

}

//leetcode submit region end(Prohibit modification and deletion)
