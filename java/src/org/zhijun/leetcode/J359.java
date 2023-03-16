package org.zhijun.leetcode;
//[369]至少有K个重复字符串的最长子串
//给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
//
//
//
// 示例 1：
//
//
//输入：s = "aaabb", k = 3
//输出：3
//解释：最长子串为 "aaa" ，其中 'a' 重复了 3 次。
//
//
// 示例 2：
//
//
//输入：s = "ababbc", k = 2
//输出：5
//解释：最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文字母组成
// 1 <= k <= 10⁵
//
//
// Related Topics 哈希表 字符串 分治 滑动窗口 👍 796 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class J359 {
    public int longestSubstring(String s, int k) {
        return dfs(s, 0, s.length() - 1, k);
    }

    public int dfs(String s, int left, int right, int k) {
        int[] count = new int[26];
        for (int i = left; i <= right; i++) {
            count[s.charAt(i) - 'a']++;
        }

        char split = 0;
        for (int i = 0; i <26; i++) {
            if (count[i] > 0 && count[i] < k) {
                split = (char) (i + 'a');
                break;
            }
        }
        if (split == 0) {
            return right - left + 1;
        }

        int ret = 0;
        int i = left;
        while (i <= right) {
            while (i <= right && s.charAt(i) == split) {
                i++;
            }
            if (i > right) {
                break;
            }
            int start = i;
            while (i <= right && s.charAt(i) != split) {
                i++;
            }
            ret = Math.max(ret, dfs(s, start, i - 1, k));
        }

        return ret;
    }

    public static void main(String[] args) {
        J359 solution = new J359();
        System.out.println(solution.longestSubstring("aaabb", 3));
        System.out.println(solution.longestSubstring("ababbc", 2));
        System.out.println(solution.longestSubstring("aacbbddee", 3));

    }
}
//leetcode submit region end(Prohibit modification and deletion)
