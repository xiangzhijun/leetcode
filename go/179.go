//[179]最大数
//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//
// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//
//
// 示例 1：
//
//
//输入：nums = [10,2]
//输出："210"
//
// 示例 2：
//
//
//输入：nums = [3,30,34,5,9]
//输出："9534330"
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 10⁹
//
//
// Related Topics 贪心 字符串 排序 👍 1055 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"math"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	compare := func(i, j int) bool {
		nI, nJ := nums[i], nums[j]
		if nI == 0 {
			return false
		}
		if nJ == 0 {
			return true
		}
		sI := int(math.Pow(10, float64(int(math.Log10(float64(nI)))+1)))
		sJ := int(math.Pow(10, float64(int(math.Log10(float64(nJ)))+1)))
		return sI*nJ+nI < sJ*nI+nJ
	}

	sort.Slice(nums, compare)
	index, n := 0, len(nums)-1
	for index = 0; index < n; index++ {
		if nums[index] != 0 {
			break
		}
	}
	nums = nums[index:]
	result := ""
	for _, n := range nums {
		result += strconv.Itoa(n)
	}
	return result
}

func main() {
	println(largestNumber([]int{3, 30, 34, 5, 9}))
	println(largestNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	println(largestNumber([]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
	println(largestNumber([]int{0, 0}))
}

//leetcode submit region end(Prohibit modification and deletion)
