package org.zhijun.leetcode;
//[967]连续差相同的数字
//返回所有长度为 n 且满足其每两个连续位上的数字之间的差的绝对值为 k 的 非负整数 。
//
// 请注意，除了 数字 0 本身之外，答案中的每个数字都 不能 有前导零。例如，01 有一个前导零，所以是无效的；但 0 是有效的。
//
// 你可以按 任何顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：n = 3, k = 7
//输出：[181,292,707,818,929]
//解释：注意，070 不是一个有效的数字，因为它有前导零。
//
//
// 示例 2：
//
//
//输入：n = 2, k = 1
//输出：[10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]
//
// 示例 3：
//
//
//输入：n = 2, k = 0
//输出：[11,22,33,44,55,66,77,88,99]
//
//
// 示例 4：
//
//
//输入：n = 2, k = 2
//输出：[13,20,24,31,35,42,46,53,57,64,68,75,79,86,97]
//
//
//
//
// 提示：
//
//
// 2 <= n <= 9
// 0 <= k <= 9
//
//
// Related Topics 广度优先搜索 回溯 👍 90 👎 0


import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

//leetcode submit region begin(Prohibit modification and deletion)
public class J967 {
    public int[] numsSameConsecDiff(int n, int k) {
        List<Integer> rstList = new ArrayList<>();
        for (int i = 1; i < 10; i++) {
            if (i - k > 0 || k - i >= 0) {
                bfs(n - 1, k, i, rstList);
            }
        }

        int[] rst = new int[rstList.size()];
        for (int i = 0; i < rstList.size(); i++) {
            rst[i] = rstList.get(i);
        }
        return rst;
    }

    public void bfs(int n, int k, int num, List<Integer> rst) {
        if (n == 0) {
            rst.add(num);
            return;
        }
        int pre = num % 10;
        if (pre + k <= 9 && k > 0) {
            bfs(n - 1, k, num * 10 + pre + k, rst);
        }
        if (pre - k >= 0) {
            bfs(n - 1, k, num * 10 + pre - k, rst);
        }
    }

    public static void main(String[] args) {
        J967 solution = new J967();
        System.out.println(Arrays.toString(solution.numsSameConsecDiff(3, 7)));
        System.out.println(Arrays.toString(solution.numsSameConsecDiff(2, 1)));
        System.out.println(Arrays.toString(solution.numsSameConsecDiff(2, 0)));
        System.out.println(Arrays.toString(solution.numsSameConsecDiff(2, 2)));

    }

}
//leetcode submit region end(Prohibit modification and deletion)

