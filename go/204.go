//[204]计数质数
//给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
//
//
//
// 示例 1：
//
//
//输入：n = 10
//输出：4
//解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
//
//
// 示例 2：
//
//
//输入：n = 0
//输出：0
//
//
// 示例 3：
//
//
//输入：n = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= n <= 5 * 10⁶
//
//
// Related Topics 数组 数学 枚举 数论 👍 986 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func countPrimes(n int) int {
	flags := make([]bool, n, n)
	count := 0
	for i := 2; i < n; i++ {
		if !flags[i] {
			count++
			if i*i < n {
				//i是质数，那边i的整倍数一定是合数
				for j := i * i; j < n; j += i {
					flags[j] = true
				}
			}
		}
	}
	return count
}

func main() {
	println(countPrimes(10))
}

//leetcode submit region end(Prohibit modification and deletion)
