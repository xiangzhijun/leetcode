package org.zhijun.leetcode;

//[96]不同的二叉搜索树
//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：5
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 19
//
//
// Related Topics 树 二叉搜索树 数学 动态规划 二叉树 👍 2202 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class J96 {
    public int numTrees(int n) {
        int[] dp = new int[n + 1];
        dp[1] = 1;
        int left, right;
        for (int i = 2; i <= n; i++) {
            int sum = 0;
            for (int j = 1; j <= i; j++) {
                left = j - 1 <= 0 ? 0 : dp[j - 1];
                right = j + 1 > i ? 0 : dp[i - j];
                if (left * right == 0) {
                    sum += left + right;
                } else {
                    sum += left * right;
                }
            }

            dp[i] = sum;
        }
        return dp[n];
    }


    public static void main(String[] args) {
        J96 solution = new J96();
        System.out.println(solution.numTrees(3));
        System.out.println(solution.numTrees(1));
    }
}
//leetcode submit region end(Prohibit modification and deletion)

