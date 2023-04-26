package org.zhijun.leetcode;

//[438]找到字符串中所有字母异位词
//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
//
//
//
// 示例 1:
//
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
//
// 示例 2:
//
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
//
//
//
// 提示:
//
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s 和 p 仅包含小写字母
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 1154 👎 0


import java.util.ArrayList;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class J438 {
    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> result = new ArrayList<>();
        if (p.length() > s.length()) return result;
        int[] pc = new int[26], sc = new int[26];
        for (int i = 0; i < p.length(); i++) {
            pc[p.charAt(i) - 'a']++;
        }

        int size = 0;
        for (int i : pc) {
            if (i > 0) size++;
        }

        int start = 0, end = 0, count = 0;
        while (end < s.length()) {
            if (end - start == p.length()) {
                int startIndex = s.charAt(start) - 'a';
                if (sc[startIndex] > 0) {
                    if (sc[startIndex] == pc[startIndex]) {
                        count--;
                    }
                    sc[startIndex]--;
                }
                start++;
            }

            int endIndex = s.charAt(end) - 'a';
            if (pc[endIndex] > 0) {
                sc[endIndex]++;
                if (sc[endIndex] == pc[endIndex]) {
                    count++;
                }
            }

            if (count == size) {
                result.add(start);
            }
            end++;
        }

        return result;
    }

    public static void main(String[] args) {
        J438 solution = new J438();
        System.out.println(solution.findAnagrams("cbaebabacd", "abc"));
        System.out.println(solution.findAnagrams("abab", "ab"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
