package org.zhijun.leetcode;


//[448]找到所有数组中消失的数字
//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数
//字，并以数组的形式返回结果。
//
//
//
// 示例 1：
//
//
//输入：nums = [4,3,2,7,8,2,3,1]
//输出：[5,6]
//
//
// 示例 2：
//
//
//输入：nums = [1,1]
//输出：[2]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 10⁵
// 1 <= nums[i] <= n
//
//
// 进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
//
// Related Topics 数组 哈希表 👍 1212 👎 0


import java.util.ArrayList;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
class J448 {
    public List<Integer> findDisappearedNumbers(int[] nums) {
        int n = nums.length;
        int v = 0, temp = 0;
        for (int i = 0; i < n; i++) {
            v = nums[i];
            while (nums[v - 1] != v) {
                temp = nums[v - 1];
                nums[v - 1] = v;
                v = temp;
            }
        }
        List<Integer> result = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (nums[i] != i + 1) {
                result.add(i + 1);
            }
        }
        return result;
    }

    public static void main(String[] args) {

    }
}
//leetcode submit region end(Prohibit modification and deletion)
