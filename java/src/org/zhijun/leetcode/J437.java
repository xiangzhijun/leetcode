package org.zhijun.leetcode;


//[437]路径总和III
//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
//输出：3
//解释：和等于 8 的路径有 3 条，如图所示。
//
//
// 示例 2：
//
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：3
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [0,1000]
//
// -10⁹ <= Node.val <= 10⁹
// -1000 <= targetSum <= 1000
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 1593 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 * int val;
 * TreeNode left;
 * TreeNode right;
 * TreeNode() {}
 * TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) {
 * this.val = val;
 * this.left = left;
 * this.right = right;
 * }
 * }
 */
class J437 {
    public int pathSum(TreeNode root, int targetSum) {
        //前缀和，保存的是所有root到当前节点路径中所有节点的路径和
        Map<Long, Integer> prefix = new HashMap<>();
        prefix.put(0L, 1);
        return dfs(root, targetSum, 0, prefix);
    }

    public int dfs(TreeNode node, int targetSum, long curr, Map<Long, Integer> prefix) {
        if (node == null) {
            return 0;
        }
        curr += node.val;
        int count = prefix.getOrDefault(curr - targetSum, 0);
        prefix.put(curr, prefix.getOrDefault(curr, 0) + 1);
        count += dfs(node.left, targetSum, curr, prefix);
        count += dfs(node.right, targetSum, curr, prefix);
        prefix.put(curr, prefix.get(curr) - 1);

        return count;
    }

    static class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
    }
}
//leetcode submit region end(Prohibit modification and deletion)
