//[198]打家劫舍
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上
//被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//
//
// 示例 1：
//
//
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//
// 示例 2：
//
//
//输入：[2,7,9,3,1]
//输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400
//
//
// Related Topics 数组 动态规划 👍 2372 👎 0
//[1,2,3,4,5,6]
//[1]			1
//[0,1]			2
//[1,0,1]		4
//[0,1,0,1]		6
//[1,0,1,0,1]	9
//[0,1,0,1,0,1]	12

//[6,5,4,3,2,8,9]
//[1]			6
//[1,0]			6
//[1,0,1]		10
//[1,0,1,0]		10
//[1,0,1,0,1]	12
//[1,0,1,0,0,1] 18
//[1,0,1,0,1,0,1]

//leetcode submit region begin(Prohibit modification and deletion)
package main

func rob(nums []int) int {
	/*
		动态规划
		i表示当前房间,dp[i]表示前i个房间偷取的最大值；
		i-1未偷:
			dp[i]=dp[i-1]+num[i]；
		i-1已偷：
			如果dp[i-2]+num[i]>dp[i-1]，则偷取i-1改为偷取i，dp[i]=dp[i-2]+num[i]；
			否则i不偷，dp[i]=dp[i-1]

	*/

	n := len(nums)
	//动态规划，dp[i]表示num[0:i-1]最大值
	dp := make([]int, n+1, n+1)
	dp[0], dp[1] = 0, nums[0]
	last := true //表示前一个房间是否已偷
	for i := 1; i < n; i++ {
		//前一个房间已偷
		if last {
			//当前房间+前面第二个房间最大值>前一个房间最大值，则替换前一个房间
			if dp[i-1]+nums[i] > dp[i] {
				dp[i+1] = dp[i-1] + nums[i]
				last = true
			} else {
				//当前房间+前面第二个房间最大值>前一个房间最大值，当前房间不偷
				dp[i+1] = dp[i]
				last = false
			}
		} else {
			//前一个房间未偷，直接偷取当前房间
			last = true
			dp[i+1] = dp[i] + nums[i]
		}
	}

	return dp[n]
}

func main() {
	println(rob([]int{2, 3, 2}))                //4
	println(rob([]int{1, 2, 3, 1}))             //4
	println(rob([]int{2, 7, 9, 3, 1}))          //12
	println(rob([]int{1, 2, 3, 4, 5, 6}))       //12
	println(rob([]int{6, 5, 4, 3, 2, 1, 8, 9})) //21
}

//leetcode submit region end(Prohibit modification and deletion)