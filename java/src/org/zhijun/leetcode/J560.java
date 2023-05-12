package org.zhijun.leetcode;


//[560]和为K的子数组
//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1], k = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3], k = 3
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁴
// -1000 <= nums[i] <= 1000
// -10⁷ <= k <= 10⁷
//
//
// Related Topics 数组 哈希表 前缀和 👍 1921 👎 0


import java.util.HashMap;
import java.util.Map;

//leetcode submit region begin(Prohibit modification and deletion)
class J560 {
    public int subarraySum(int[] nums, int k) {
        int result = 0, prefix = 0;
        Map<Integer, Integer> prefixMap = new HashMap<>();
        prefixMap.put(0, 1);
        for (int i = 0; i < nums.length; i++) {
            prefix += nums[i];
            result += prefixMap.getOrDefault(prefix - k, 0);
            prefixMap.put(prefix, prefixMap.getOrDefault(prefix, 0) + 1);
        }
        return result;
    }
}
//leetcode submit region end(Prohibit modification and deletion)

