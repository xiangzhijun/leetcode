# 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
#
#
#
#  注意：
#
#
#  对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
#  如果 s 中存在这样的子串，我们保证它是唯一的答案。
#
#
#
#
#  示例 1：
#
#
# 输入：s = "ADOBECODEBANC", t = "ABC"
# 输出："BANC"
#
#
#  示例 2：
#
#
# 输入：s = "a", t = "a"
# 输出："a"
#
#
#  示例 3:
#
#
# 输入: s = "a", t = "aa"
# 输出: ""
# 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
# 因此没有符合条件的子字符串，返回空字符串。
#
#
#
#  提示：
#
#
#  1 <= s.length, t.length <= 10⁵
#  s 和 t 由英文字母组成
#
#
#
# 进阶：你能设计一个在
# o(n) 时间内解决此问题的算法吗？
#
#  Related Topics 哈希表 字符串 滑动窗口 👍 2054 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def minWindow(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        m = len(s)
        n = len(t)
        if m < n:
            return ""

        # 统计t中每个字符串出现的次数
        target_dict = {}
        for c in t:
            if c not in target_dict:
                target_dict[c] = 1
            else:
                target_dict[c] += 1
        target_size = len(target_dict)

        completed_count = 0
        found_dict = {c: 0 for c in target_dict}
        start_index, end_index = -1, 999999999999999
        left = 0

        for right in range(m):
            c = s[right]
            if c not in target_dict:
                continue

            found_dict[c] += 1
            if found_dict[c] == target_dict[c]:
                completed_count += 1
                while completed_count >= target_size:
                    left_c = s[left]
                    if left_c in target_dict:
                        if end_index - start_index >= right - left:
                            start_index = left
                            end_index = right
                        found_dict[left_c] -= 1
                        if found_dict[left_c] < target_dict[left_c]:
                            completed_count -= 1

                    left += 1

        if start_index >= 0:
            return s[start_index:end_index + 1]
        return ""

# print(Solution().minWindow('ADOBECODEBANC', 'ABC'))

# leetcode submit region end(Prohibit modification and deletion)
