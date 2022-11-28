# 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
#
#  说明：本题中，我们将空字符串定义为有效的回文串。
#
#
#
#  示例 1:
#
#
# 输入: "A man, a plan, a canal: Panama"
# 输出: true
# 解释："amanaplanacanalpanama" 是回文串
#
#
#  示例 2:
#
#
# 输入: "race a car"
# 输出: false
# 解释："raceacar" 不是回文串
#
#
#
#
#  提示：
#
#
#  1 <= s.length <= 2 * 10⁵
#  字符串 s 由 ASCII 字符组成
#
#
#  Related Topics 双指针 字符串 👍 563 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """

        def is_letter_or_num(c):
            return '0' <= c <= '9' or 'A' <= c <= 'Z' or 'a' <= c <= 'z'

        i = 0
        j = len(s) - 1
        while j > i:
            if not is_letter_or_num(s[i]):
                i += 1
            elif not is_letter_or_num(s[j]):
                j -= 1
            else:
                if s[i].lower() != s[j].lower():
                    return False
                i += 1
                j -= 1

        return True


# print(Solution().isPalindrome('A man, a plan, a canal: Panama'))
# print(Solution().isPalindrome('0P'))

# leetcode submit region end(Prohibit modification and deletion)
