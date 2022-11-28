# 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
#
#  回文串 是正着读和反着读都一样的字符串。
#
#
#
#  示例 1：
#
#
# 输入：s = "aab"
# 输出：[["a","a","b"],["aa","b"]]
#
#
#  示例 2：
#
#
# 输入：s = "a"
# 输出：[["a"]]
#
#
#
#
#  提示：
#
#
#  1 <= s.length <= 16
#  s 仅由小写英文字母组成
#
#
#  Related Topics 字符串 动态规划 回溯 👍 1243 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def partition(self, s):
        """
        :type s: str
        :rtype: List[List[str]]
        """
        n = len(s)
        p = [[True] * n for _ in range(n)]
        # 需要从后往前遍历，这样才能保证i~j之间的都是遍历过的
        for i in range(n - 1, -1, -1):
            for j in range(i + 1, n):
                p[i][j] = s[i] == s[j] and p[i + 1][j - 1]

        result = []
        temp = []

        def backtrace(i):
            if i == n:
                result.append(temp[:])
            else:
                for j in range(i, n):
                    if p[i][j]:
                        temp.append(s[i:j + 1])
                        backtrace(j + 1)
                        temp.pop()

        backtrace(0)
        return result

# print(Solution().partition('abbab'))

# leetcode submit region end(Prohibit modification and deletion)
