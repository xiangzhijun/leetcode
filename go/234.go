//[234]回文链表
//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,2,1]
//输出：true
//
//
// 示例 2：
//
//
//输入：head = [1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 链表中节点数目在范围[1, 10⁵] 内
// 0 <= Node.val <= 9
//
//
//
//
// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
// Related Topics 栈 递归 链表 双指针 👍 1597 👎 0

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

func isPalindrome(head *ListNode) bool {
	//快慢指针找到中间节点
	//反转后半部分链表
	//比较前半部分和后半部分判断是否回文
	//再反转后半部分链表还原

	var (
		reverseList func(node *ListNode) *ListNode
	)

	//反转链表
	reverseList = func(node *ListNode) *ListNode {
		cur := node
		var newHead *ListNode
		for cur != nil {
			temp := cur
			cur = cur.Next
			newHead, temp.Next = temp, newHead
		}
		return newHead
	}

	//快慢指针寻找中间节点
	slowNode, fastNode := head, head
	for fastNode != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
		if fastNode != nil {
			fastNode = fastNode.Next
		}
	}

	//反转后半部分
	backHead := reverseList(slowNode)

	//比较两个链表是否相同，判断是否回文
	list1, list2 := head, backHead
	rst := true
	for list2 != nil {
		if list1.Val != list2.Val {
			rst = false
			break
		}
		list1, list2 = list1.Next, list2.Next
	}

	//把后半部分反转回来
	reverseList(backHead)

	return rst
}

func test(nums []int) {
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}
		cur.Next = node
		cur = node
	}
	println(isPalindrome(head))
}
func main() {
	test([]int{1, 2, 2, 1})
	test([]int{1, 2, 3, 2, 1})
	test([]int{1, 2, 2, 3, 1})
}

//leetcode submit region end(Prohibit modification and deletion)
