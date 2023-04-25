package org.zhijun.leetcode;

//[416]分割等和子集
//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 1731 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class J416 {
    public boolean canPartition(int[] nums) {
        if (nums.length < 2) return false;
        int sum = 0, max = 0;
        for (int num : nums) {
            sum += num;
            max = Integer.max(max, num);
        }
        if (sum % 2 != 0) return false;
        int half = sum / 2;
        if (max > half) return false;

        boolean[] dp = new boolean[half + 1];
        dp[0] = true;
        for (int num : nums) {
            for (int j = half; j >= num; j--) {
                dp[j] |= dp[j - num];
            }
        }

        return dp[half];
    }

    public static void main(String[] args) {
        J416 solution = new J416();
        System.out.println(solution.canPartition(new int[]{1, 5, 11, 5}));
        System.out.println(solution.canPartition(new int[]{1, 2, 3, 5}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
