package org.zhijun.leetcode;


//[51]N 皇后
//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 数组 回溯 👍 1743 👎 0


import java.util.*;

//leetcode submit region begin(Prohibit modification and deletion)
class J51 {
    Set<Integer> column = new HashSet<>();
    Set<Integer> bias1 = new HashSet<>();
    Set<Integer> bias2 = new HashSet<>();

    public List<List<String>> solveNQueens(int n) {
        List<List<String>> result = new ArrayList<>();
        int[] queue = new int[n];

        dfs(queue, 0, n, result);
        return result;
    }

    public void dfs(int[] queue, int row, int n, List<List<String>> result) {
        if (row == n) {
            result.add(createBoard(queue, n));
            return;
        }
        for (int i = 0; i < n; i++) {
            if (column.contains(i)) continue;
            if (bias1.contains(row - i)) continue;
            if (bias2.contains(row + i)) continue;

            column.add(i);
            bias1.add(row - i);
            bias2.add(row + i);
            queue[row] = i;
            dfs(queue, row + 1, n, result);
            queue[row] = -1;
            column.remove(i);
            bias1.remove(row - i);
            bias2.remove(row + i);
        }
    }

    public List<String> createBoard(int[] queue, int n) {
        char[] chars = new char[n];
        Arrays.fill(chars, '.');
        List<String> board = new ArrayList<>(n);
        for (int i = 0; i < n; i++) {
            chars[queue[i]] = 'Q';
            board.add(new String(chars));
            chars[queue[i]] = '.';
        }
        return board;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
