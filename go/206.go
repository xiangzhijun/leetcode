//[206]反转链表
//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//
//
// 示例 3：
//
//
//输入：head = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000
//
//
//
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
// Related Topics 递归 链表 👍 2885 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

package main

import "encoding/json"

type ListNode struct {
	Val  int
	Next *ListNode
}

//迭代
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head
	head, newHead.Next = head.Next, nil
	for head != nil {
		next := head.Next
		head.Next, newHead = newHead, head
		head = next
	}

	return newHead
}

//递归
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	newHead := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return newHead
//}

func test(list []int) {
	if len(list) == 0 {
		return
	}
	head := &ListNode{Val: list[0]}
	node := head
	for i := 1; i < len(list); i++ {
		newNode := ListNode{Val: list[i]}
		node.Next = &newNode
		node = &newNode
	}

	bytes, _ := json.Marshal(list)
	println(string(bytes))

	newHead := reverseList(head)
	result := make([]int, 0)
	for newHead != nil {
		result = append(result, newHead.Val)
		newHead = newHead.Next
	}
	bytes, _ = json.Marshal(result)
	println(string(bytes))
}

func main() {
	test([]int{1, 2, 3, 4, 5})
}

//leetcode submit region end(Prohibit modification and deletion)
