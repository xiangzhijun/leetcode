//[347]前K个高频元素
//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//
//
// 示例 2:
//
//
//输入: nums = [1], k = 1
//输出: [1]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
//
//
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
//
// Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 1485 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"encoding/json"
	"math/rand"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int, 0)
	c, ok := 0, false
	//统计所有元素出现的次数
	for _, n := range nums {
		c, ok = count[n]
		if ok {
			count[n] = c + 1
		} else {
			count[n] = 1
		}
	}

	n := len(count)
	numsArr := make([]int, 0, n)
	countArr := make([]int, 0, n)
	for num, cu := range count {
		numsArr = append(numsArr, num)
		countArr = append(countArr, cu)
	}
	if n <= k {
		return numsArr
	}

	//用快速选择算法选择topK
	k = n - k
	selectTopK(numsArr, countArr, 0, n-1, k)
	return numsArr[k:]
}

func selectTopK(numsArr, countArr []int, left, right, k int) {
	//random
	randIndex := rand.Intn(right-left+1) + left
	swap(numsArr, countArr, left, randIndex)

	x := countArr[left]
	i := right
	for j := i; j > left; j-- {
		if countArr[j] > x {
			swap(numsArr, countArr, i, j)
			i--
		}
	}
	swap(numsArr, countArr, i, left)

	if i == k {
		return
	} else if i > k {
		selectTopK(numsArr, countArr, left, i-1, k)
	} else {
		selectTopK(numsArr, countArr, i+1, right, k)
	}
}

func swap(numsArr, countArr []int, i, j int) {
	numsArr[i], numsArr[j] = numsArr[j], numsArr[i]
	countArr[i], countArr[j] = countArr[j], countArr[i]
}

func topKFrequentTest(nums []int, k int) {
	bytes, _ := json.Marshal(nums)
	println(string(bytes), k)
	rst := topKFrequent(nums, k)
	bytes, _ = json.Marshal(rst)
	println(string(bytes), "\n")
}

func main() {
	topKFrequentTest([]int{1, 1, 1, 2, 2, 3}, 2)
	topKFrequentTest([]int{-1, -1}, 1)
	topKFrequentTest([]int{4, 1, -1, 2, -1, 2, 3}, 2)
}

//leetcode submit region end(Prohibit modification and deletion)
