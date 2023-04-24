package org.zhijun.leetcode;


//[399]除法求值
//给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和
//values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
//
// 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj =
// ? 的结果作为答案。
//
// 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替
//代这个答案。
//
// 注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
//
//
//
// 示例 1：
//
//
//输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"]
//,["b","a"],["a","e"],["a","a"],["x","x"]]
//输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
//解释：
//条件：a / b = 2.0, b / c = 3.0
//问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
//结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
//
//
// 示例 2：
//
//
//输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0],
//queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
//输出：[3.75000,0.40000,5.00000,0.20000]
//
//
// 示例 3：
//
//
//输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],[
//"a","c"],["x","y"]]
//输出：[0.50000,2.00000,-1.00000,-1.00000]
//
//
//
//
// 提示：
//
//
// 1 <= equations.length <= 20
// equations[i].length == 2
// 1 <= Ai.length, Bi.length <= 5
// values.length == equations.length
// 0.0 < values[i] <= 20.0
// 1 <= queries.length <= 20
// queries[i].length == 2
// 1 <= Cj.length, Dj.length <= 5
// Ai, Bi, Cj, Dj 由小写英文字母与数字组成
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 数组 最短路 👍 922 👎 0


import java.util.*;

//leetcode submit region begin(Prohibit modification and deletion)
class J399 {
    public double[] calcEquation(List<List<String>> equations, double[] values, List<List<String>> queries) {
        Map<String, Map<String, Double>> edgeMap = new HashMap<>();
        Map<String, Double> temp = null;
        for (int i = 0; i < equations.size(); i++) {
            List<String> pair = equations.get(i);
            if ((temp = edgeMap.get(pair.get(0))) == null) {
                temp = new HashMap<>();
                edgeMap.put(pair.get(0), temp);
                temp.put(pair.get(0), 1D);
            }
            temp.put(pair.get(1), values[i]);

            if ((temp = edgeMap.get(pair.get(1))) == null) {
                temp = new HashMap<>();
                edgeMap.put(pair.get(1), temp);
                temp.put(pair.get(1), 1D);
            }
            temp.put(pair.get(0), 1 / values[i]);
        }
        Set<String> visited = new HashSet<>();
        double[] result = new double[queries.size()];
        for (int i = 0; i < queries.size(); i++) {
            List<String> query = queries.get(i);
            result[i] = dfs(edgeMap, query.get(0), query.get(1), visited);
            visited.clear();
        }

        return result;
    }

    public double dfs(Map<String, Map<String, Double>> edgeMap, String left, String right, Set<String> visited) {
        if (!edgeMap.containsKey(left) || !edgeMap.containsKey(right) || visited.contains(left)) {
            return -1;
        }
        Map<String, Double> edge = edgeMap.get(left);
        if (edge.containsKey(right)) {
            return edge.get(right);
        }

        visited.add(left);
        for (Map.Entry<String, Double> entry : edge.entrySet()) {
            double v = dfs(edgeMap, entry.getKey(), right, visited);
            if (v > 0) {
                return entry.getValue() * v;
            }
        }
        return -1;
    }


    public static void main(String[] args) {
        J399 solution = new J399();
        List<List<String>> equations = new ArrayList<>();
        equations.add(Arrays.asList("a", "b"));
        equations.add(Arrays.asList("b", "c"));
        double[] values = new double[]{2.0, 3.0};

        List<List<String>> queries = new ArrayList<>();
        queries.add(Arrays.asList("a", "c"));
        queries.add(Arrays.asList("b", "a"));
        queries.add(Arrays.asList("a", "e"));
        queries.add(Arrays.asList("a", "a"));
        queries.add(Arrays.asList("x", "x"));

        System.out.println(Arrays.toString(solution.calcEquation(equations, values, queries)));


    }

}
//leetcode submit region end(Prohibit modification and deletion)
