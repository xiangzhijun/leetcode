//[217]存在重复元素
//给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,1]
//输出：true
//
// 示例 2：
//
//
//输入：nums = [1,2,3,4]
//输出：false
//
// 示例 3：
//
//
//输入：nums = [1,1,1,3,3,4,3,2,4,2]
//输出：true
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics 数组 哈希表 排序 👍 886 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func containsDuplicate(nums []int) bool {
	filter := make(map[int]bool)
	for _, n := range nums {
		if ok, _ := filter[n]; ok {
			return true
		}
		filter[n] = true
	}
	return false
}

func main() {
	println(containsDuplicate([]int{1, 2, 3, 1}))
	println(containsDuplicate([]int{1, 2, 3, 4}))
}

//leetcode submit region end(Prohibit modification and deletion)
