//[295]数据流的中位数
//中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
//
//
// 例如 arr = [2,3,4] 的中位数是 3 。
// 例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
//
//
// 实现 MedianFinder 类:
//
//
// MedianFinder() 初始化 MedianFinder 对象。
// void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
// double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10⁻⁵ 以内的答案将被接受。
//
//
// 示例 1：
//
//
//输入
//["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
//[[], [1], [2], [], [3], []]
//输出
//[null, null, null, 1.5, null, 2.0]
//
//解释
//MedianFinder medianFinder = new MedianFinder();
//medianFinder.addNum(1);    // arr = [1]
//medianFinder.addNum(2);    // arr = [1, 2]
//medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
//medianFinder.addNum(3);    // arr[1, 2, 3]
//medianFinder.findMedian(); // return 2.0
//
// 提示:
//
//
// -10⁵ <= num <= 10⁵
// 在调用 findMedian 之前，数据结构中至少有一个元素
// 最多 5 * 10⁴ 次调用 addNum 和 findMedian
//
//
// Related Topics 设计 双指针 数据流 排序 堆（优先队列） 👍 791 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

type Node struct {
	value int
	pre   *Node
	next  *Node
}

type MedianFinder struct {
	size int
	head *Node
	mid  *Node
}

func Constructor() MedianFinder {
	return MedianFinder{head: &Node{}}
}

func (this *MedianFinder) AddNum(num int) {
	insertIndex := 0
	cur := this.head
	for cur.next != nil {
		if num <= cur.next.value {
			break
		}
		cur = cur.next
		insertIndex++
	}
	node := &Node{
		value: num,
		pre:   cur,
		next:  cur.next,
	}
	if cur.next != nil {
		cur.next.pre = node
	}
	cur.next = node

	if this.mid == nil {
		this.mid = node
	} else {
		midIndex := this.size/2 + this.size%2 - 1
		if insertIndex <= midIndex && this.size%2 == 1 {
			this.mid = this.mid.pre
		} else if insertIndex > midIndex && this.size%2 == 0 {
			this.mid = this.mid.next
		}
	}
	this.size++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size%2 == 0 {
		return float64(this.mid.value+this.mid.next.value) / 2
	}
	return float64(this.mid.value)
}

func main() {
	medianFinder := Constructor()
	//medianFinder.AddNum(1)             // arr = [1]
	//medianFinder.AddNum(2)             // arr = [1, 2]
	//println(medianFinder.FindMedian()) // 返回 1.5 ((1 + 2) / 2)
	//medianFinder.AddNum(3)             // arr[1, 2, 3]
	//println(medianFinder.FindMedian()) // return 2.0

	medianFinder.AddNum(-1)
	println(medianFinder.FindMedian())
	medianFinder.AddNum(-2)
	println(medianFinder.FindMedian())
	medianFinder.AddNum(-3)
	println(medianFinder.FindMedian())
	medianFinder.AddNum(-4)
	println(medianFinder.FindMedian())
	medianFinder.AddNum(-5)
	println(medianFinder.FindMedian())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//leetcode submit region end(Prohibit modification and deletion)
