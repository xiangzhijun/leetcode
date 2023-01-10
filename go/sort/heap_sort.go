package main

import "encoding/json"

/**
堆排序
*/
func heapSort(arr []int) {
	n := len(arr)
	//构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, n)
	}

	//交换堆顶和堆尾元素，重新调整堆结构
	for j := n - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		adjustHeap(arr, 0, j)
	}
}

//调整堆
func adjustHeap(heap []int, index int, length int) {
	temp := heap[index]
	for k := index*2 + 1; k < length; k = k*2 + 1 {
		//找到当前节点的最大子节点
		if k+1 < length && heap[k] < heap[k+1] {
			k++
		}

		if heap[k] > temp {
			//如果最大子节点大于当前节点，将该子节点复制给当前节点，当前节点下沉
			heap[index] = heap[k]
			index = k
		} else {
			break
		}
	}
	//将要调整的值放到最终位置
	heap[index] = temp
}

func test(arr []int) {
	heapSort(arr)
	bytes, _ := json.Marshal(arr)
	println(string(bytes))
}

func main() {
	test([]int{6, 5, 4, 7, 2, 1})
}
