# 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
#
#  请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
#
#
#
#  示例 1：
#
#
# 输入：nums = [100,4,200,1,3,2]
# 输出：4
# 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
#
#  示例 2：
#
#
# 输入：nums = [0,3,7,2,5,8,4,6,0,1]
# 输出：9
#
#
#
#
#  提示：
#
#
#  0 <= nums.length <= 10⁵
#  -10⁹ <= nums[i] <= 10⁹
#
#
#  Related Topics 并查集 数组 哈希表 👍 1367 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def longestConsecutive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        max_len = 0
        num_set = set(nums)
        for n in nums:
            if n + max_len in num_set and n - 1 not in num_set:
                max_num = n + 1
                while max_num in num_set:
                    max_num += 1
                max_len = max(max_len, max_num - n)
        return max_len


# print(Solution().longestConsecutive([100, 4, 200, 1, 3, 2]))
# print(Solution().longestConsecutive([9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6]))
# leetcode submit region end(Prohibit modification and deletion)
