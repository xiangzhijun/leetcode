package org.zhijun.leetcode;


//[309]最佳买卖股票时机含冷冻期
//给定一个整数数组
// prices，其中第 prices[i] 表示第 i 天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
//
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
//
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
//
//
//
// 示例 1:
//
//
//输入: prices = [1,2,3,0,2]
//输出: 3
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
//
// 示例 2:
//
//
//输入: prices = [1]
//输出: 0
//
//
//
//
// 提示：
//
//
// 1 <= prices.length <= 5000
// 0 <= prices[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 1482 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
class J309 {
    public int maxProfit(int[] prices) {
        //动态规划
        //每天有四种操作方案，买、卖、无成本空闲、有成本空闲
        //当天买：前一天必须是无成本空闲
        //当天买：前一天是买或者有成本空闲
        //当天无成本空闲：前一天必须是无成本空闲
        //当天有成本空闲：前一天必须是买货有成本空闲
        //取最后一天四种操作中的最大值即可
        int buy = 0, sell = 0, freeze = 0, idleBuy = 0, buyCost = prices[0], idleCost = -1;
        int currBuy, currSell, currFreeze, currIdleBuy, currBuyCost, currIdleCost;
        for (int i = 1; i < prices.length; i++) {
//            System.out.printf("[%d]: buy :%d sell:%d free:%d idle:%d %d %d%n", prices[i - 1], buy, sell, freeze, idleBuy, buyCost, idleCost);
            //buy
            currBuy = freeze;
            currBuyCost = prices[i];

            //sell
            if (idleCost >= 0) {
                currSell = Integer.max(buy + prices[i] - buyCost, idleBuy + prices[i] - idleCost);
            } else {
                currSell = buy + prices[i] - buyCost;
            }

            //freeze
            currFreeze = Integer.max(sell, freeze);

            //idleBuy
            if (idleCost >= 0 && (idleBuy - idleCost) >= (buy - buyCost)) {
                currIdleBuy = idleBuy;
                currIdleCost = idleCost;
            } else {
                currIdleBuy = buy;
                currIdleCost = buyCost;
            }


            buy = currBuy;
            sell = currSell;
            freeze = currFreeze;
            idleBuy = currIdleBuy;
            buyCost = currBuyCost;
            idleCost = currIdleCost;
//            System.out.printf("[%d]: buy :%d sell:%d free:%d idle:%d %d %d%n", prices[i], buy, sell, freeze, idleBuy, buyCost, idleCost);
//            System.out.println();
        }

        return Integer.max(Integer.max(buy, sell), Integer.max(freeze, idleBuy));
    }

    public static void main(String[] args) {
        J309 solution = new J309();
        System.out.println(solution.maxProfit(new int[]{1, 2, 3, 0, 2}));
        System.out.println(solution.maxProfit(new int[]{6, 1, 3, 2, 4, 7}));
        System.out.println(solution.maxProfit(new int[]{2, 1, 2, 1, 0, 1, 2}));
    }
}
//leetcode submit region end(Prohibit modification and deletion)
