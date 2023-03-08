//[328]奇偶链表
//给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。
//
// 第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
//
// 请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
//
// 你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
//
//
//
// 示例 1:
//
//
//
//
//输入: head = [1,2,3,4,5]
//输出: [1,3,5,2,4]
//
// 示例 2:
//
//
//
//
//输入: head = [2,1,3,5,6,4,7]
//输出: [2,3,6,7,1,5,4]
//
//
//
// 提示:
//
//
// n == 链表中的节点数
// 0 <= n <= 10⁴
// -10⁶ <= Node.val <= 10⁶
//
//
// Related Topics 链表 👍 683 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	oddHead, evenHead := &ListNode{}, &ListNode{}
	oddCur, evenCur, temp := oddHead, evenHead, head
	for count, cur := 1, head; cur != nil; count++ {
		temp, cur.Next, cur = cur, nil, cur.Next

		if count%2 == 0 {
			evenCur.Next, evenCur = temp, temp
		} else {
			oddCur.Next, oddCur = temp, temp
		}
	}
	oddCur.Next = evenHead.Next
	return oddHead.Next
}

func main() {
	head := &ListNode{}
	cur := head
	for i := 1; i <= 5; i++ {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}
	oddEvenList(head.Next)
}

//leetcode submit region end(Prohibit modification and deletion)
