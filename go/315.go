//[315]计算右侧小于当前元素的个数
//给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是 nums[i] 右侧小于
//nums[i] 的元素的数量。
//
//
//
// 示例 1：
//
//
//输入：nums = [5,2,6,1]
//输出：[2,1,1,0]
//解释：
//5 的右侧有 2 个更小的元素 (2 和 1)
//2 的右侧仅有 1 个更小的元素 (1)
//6 的右侧有 1 个更小的元素 (1)
//1 的右侧有 0 个更小的元素
//
//
// 示例 2：
//
//
//输入：nums = [-1]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：nums = [-1,-1]
//输出：[0,0]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 932 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "encoding/json"

func countSmaller(nums []int) []int {
	var (
		n         = len(nums)
		index     = make([]int, n, n) //保存nums[i]在nums中的index，和nums一起排序
		temp      = make([]int, n, n) //归并排序合并过程中保存merge num的临时数组
		tempIndex = make([]int, n, n) //归并排序合并过程中保存merge num index的临时数组
		result    = make([]int, n, n) //保存结果
		mergeSort func(left, right int)
	)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	//归并排序
	mergeSort = func(left, right int) {
		if left >= right {
			return
		}

		mid := (left + right) >> 1
		mergeSort(left, mid)
		mergeSort(mid+1, right)

		i, j, p := left, mid+1, left
		for i <= mid && j <= right {
			if nums[i] <= nums[j] {
				temp[p] = nums[i]
				tempIndex[p] = index[i]
				result[index[i]] += j - mid - 1
				i++
				p++
			} else {
				temp[p] = nums[j]
				tempIndex[p] = index[j]
				j++
				p++
			}
		}

		for i <= mid {
			temp[p] = nums[i]
			tempIndex[p] = index[i]
			result[index[i]] += j - mid - 1
			i++
			p++
		}
		for j <= right {
			temp[p] = nums[j]
			tempIndex[p] = index[j]
			j++
			p++
		}

		for k := left; k <= right; k++ {
			nums[k] = temp[k]
			index[k] = tempIndex[k]
		}
	}

	mergeSort(0, n-1)
	return result
}

func test(nums []int) {
	rst := countSmaller(nums)
	numsBytes, _ := json.Marshal(nums)
	rstBytes, _ := json.Marshal(rst)
	println(string(numsBytes))
	println(string(rstBytes))
	println()
}

func main() {
	test([]int{5, 2, 6, 1})
}

//leetcode submit region end(Prohibit modification and deletion)
