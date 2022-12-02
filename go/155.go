//[155]最小栈
//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
//
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
//
//
//
//
// 示例 1:
//
//
//输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// -2³¹ <= val <= 2³¹ - 1
// pop、top 和 getMin 操作总是在 非空栈 上调用
// push, pop, top, and getMin最多被调用 3 * 10⁴ 次
//
//
// Related Topics 栈 设计 👍 1462 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)

type Node struct {
	value int
	next  *Node
}

type MinStack struct {
	//采用链表构建栈，方便入栈和出栈
	//数据栈
	dataStack *Node
	//最小栈，栈顶值就是数据栈里的最小值
	minStack *Node
}

func Constructor() MinStack {
	return MinStack{
		dataStack: nil,
		minStack:  nil,
	}
}

func (this *MinStack) Push(val int) {
	node := &Node{value: val, next: this.dataStack}
	this.dataStack = node
	//注意是<=不是<，因为相等的数如果不重复入栈的话，pop会乱序
	if this.minStack == nil || val <= this.minStack.value {
		minNode := &Node{value: val, next: this.minStack}
		this.minStack = minNode
	}
}

func (this *MinStack) Pop() {
	if this.dataStack != nil {
		//数据栈pop时需要同步pop最小栈
		if this.minStack.value == this.dataStack.value {
			this.minStack = this.minStack.next
		}
		this.dataStack = this.dataStack.next
	}
}

func (this *MinStack) Top() int {
	if this.dataStack != nil {
		return this.dataStack.value
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.minStack != nil {
		return this.minStack.value
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
func main() {

}
