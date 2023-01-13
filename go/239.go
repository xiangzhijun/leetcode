//[239]滑动窗口最大值
//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 2067 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "encoding/json"

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	rst := make([]int, n-k+1)
	//单调递减队列
	queue := make([]int, n)
	head, length, tail := 0, 0, 0
	for i := 0; i < n; i++ {
		//先删除队首不在[i,i+k]之间的值
		for length > 0 && i-queue[head] > k-1 {
			head = (head + 1) % n
			length--
		}
		//如果当前队尾值小于nums[i]，则删除队尾值，循环此操作
		tail = (head + length - 1) % n
		for length > 0 && nums[queue[tail]] <= nums[i] {
			length--
			tail = (head + length - 1) % n
		}
		//插入i到队尾
		length++
		tail = (head + length - 1) % n
		queue[tail] = i

		if i >= k-1 {
			//队首值即为当前窗口最大值
			rst[i-k+1] = nums[queue[head]]
		}
	}

	return rst
}

func test(nums []int, k int) {
	rst := maxSlidingWindow(nums, k)
	bytes, _ := json.Marshal(rst)
	println(string(bytes))
}

func main() {
	test([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}

//leetcode submit region end(Prohibit modification and deletion)
