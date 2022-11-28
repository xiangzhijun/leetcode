# 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
#
#  解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
#
#
#
#  示例 1：
#
#
# 输入：nums = [1,2,3]
# 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
#
#
#  示例 2：
#
#
# 输入：nums = [0]
# 输出：[[],[0]]
#
#
#
#
#  提示：
#
#
#  1 <= nums.length <= 10
#  -10 <= nums[i] <= 10
#  nums 中的所有元素 互不相同
#
#
#  Related Topics 位运算 数组 回溯 👍 1744 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        rst = [[]]
        n = len(nums)
        temp = []

        def loop(index, arr):
            for j in range(index, n):
                arr.append(nums[j])
                rst.append(arr[:])
                loop(j + 1, arr)
                arr.pop(len(arr) - 1)

        loop(0, temp)

        return rst


print(Solution().subsets([1, 2, 3]))

# leetcode submit region end(Prohibit modification and deletion)
