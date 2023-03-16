//[384]打乱数组
//给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。
//
// 实现 Solution class:
//
//
// Solution(int[] nums) 使用整数数组 nums 初始化对象
// int[] reset() 重设数组到它的初始状态并返回
// int[] shuffle() 返回数组随机打乱后的结果
//
//
//
//
// 示例 1：
//
//
//输入
//["Solution", "shuffle", "reset", "shuffle"]
//[[[1, 2, 3]], [], [], []]
//输出
//[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
//
//解释
//Solution solution = new Solution([1, 2, 3]);
//solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回 [3,
//1, 2]
//solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
//solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 50
// -10⁶ <= nums[i] <= 10⁶
// nums 中的所有元素都是 唯一的
// 最多可以调用 10⁴ 次 reset 和 shuffle
//
//
// Related Topics 数组 数学 随机化 👍 329 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "math/rand"

type Solution struct {
	nums    []int
	shuffle []int
	length  int
}

func Constructor(nums []int) Solution {
	shuffle := make([]int, len(nums))
	copy(shuffle, nums)
	return Solution{nums: nums, shuffle: shuffle, length: len(nums)}
}

func (this *Solution) Reset() []int {
	copy(this.shuffle, this.nums)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	randIndex := 0
	for i := range this.shuffle {
		randIndex = i + rand.Intn(this.length-i)
		this.shuffle[i], this.shuffle[randIndex] = this.shuffle[randIndex], this.shuffle[i]
	}
	return this.shuffle
}

func main() {

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
//leetcode submit region end(Prohibit modification and deletion)
