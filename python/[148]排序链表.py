# 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
#
#
#
#
#
#
#  示例 1：
#
#
# 输入：head = [4,2,1,3]
# 输出：[1,2,3,4]
#
#
#  示例 2：
#
#
# 输入：head = [-1,5,3,4,0]
# 输出：[-1,0,3,4,5]
#
#
#  示例 3：
#
#
# 输入：head = []
# 输出：[]
#
#
#
#
#  提示：
#
#
#  链表中节点的数目在范围 [0, 5 * 10⁴] 内
#  -10⁵ <= Node.val <= 10⁵
#
#
#
#
#  进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
#
#  Related Topics 链表 双指针 分治 排序 归并排序 👍 1753 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution(object):
    def sortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """

        sub_len = 1
        n = 0
        curr = head
        while curr:
            n += 1
            curr = curr.next

        new_head = ListNode(0, head)
        while sub_len < n:
            pre_node, curr = new_head, new_head.next
            while curr:
                h1 = curr
                for i in range(1, sub_len):
                    if curr.next:
                        curr = curr.next
                    else:
                        break

                h2 = curr.next
                curr.next = None
                curr = h2
                for i in range(1, sub_len):
                    if curr and curr.next:
                        curr = curr.next
                    else:
                        break
                temp_node = None
                if curr:
                    temp_node = curr.next
                    curr.next = None

                while h1 and h2:
                    if h1.val <= h2.val:
                        pre_node.next = h1
                        h1 = h1.next
                    else:
                        pre_node.next = h2
                        h2 = h2.next
                    pre_node = pre_node.next
                pre_node.next = h1 if h1 else h2
                while pre_node.next: pre_node = pre_node.next

                curr = temp_node
            sub_len <<= 1

        return new_head.next

# def go(values):
#     head = ListNode(0, None)
#     curr = head
#     for v in values:
#         node = ListNode(v, None)
#         curr.next = node
#         curr = node
#     return head.next
#
#
# print((Solution().sortList(go([4, 2, 1, 3]))))
# leetcode submit region end(Prohibit modification and deletion)
