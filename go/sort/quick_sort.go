package main

import "encoding/json"

func quickSort(nums []int) {
	quickSelect(nums, 0, len(nums)-1)
}

func quickSelect(nums []int, l, r int) {
	if l >= r {
		return
	}
	k := partitionLeft2Right(nums, l, r)
	quickSelect(nums, l, k-1)
	quickSelect(nums, k+1, r)
}

func partitionLeft2Right(nums []int, l, r int) int {
	//从左往右遍历选择
	//选取最右边的nums[r]作为“基准数”
	x := nums[r]
	i := l //i表示下一个比x小的数应该填充的位置
	for j := l; j < r; j++ {
		if nums[j] < x {
			//当前数比x小，则将当前数跟i交换，nums[i]大于x或等于当前值nums[j]
			swap(nums, i, j)
			//交换完成后i往后移，表示下一次可交换的位置
			i++
		}
	}
	swap(nums, i, r)
	return i
}
func partitionRight2Left(nums []int, l, r int) int {
	//从右往左遍历选择
	//选取最左边的nums[l]作为“基准数”
	x := nums[l]
	i := r
	for j := r; j > l; j-- {
		if nums[j] > x {
			swap(nums, i, j)
			i--
		}
	}
	swap(nums, i, l)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func testQuickSort(nums []int) {
	bytes, _ := json.Marshal(nums)
	println(string(bytes))
	quickSort(nums)
	bytes, _ = json.Marshal(nums)
	println(string(bytes))
	println()
}

func main() {
	testQuickSort([]int{2, 1, 6, 6, 5, 9, 2, 3, 4})
}
