//[371]两整数之和
//给你两个整数 a 和 b ，不使用 运算符 + 和 - ，计算并返回两整数之和。
//
//
//
// 示例 1：
//
//
//输入：a = 1, b = 2
//输出：3
//
//
// 示例 2：
//
//
//输入：a = 2, b = 3
//输出：5
//
//
//
//
// 提示：
//
//
// -1000 <= a, b <= 1000
//
//
// Related Topics 位运算 数学 👍 667 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func getSum(a int, b int) int {
	for b != 0 {
		c := uint(a&b) << 1
		a ^= b
		b = int(c)
	}
	return a
}
func main() {
	println(getSum(1, 2))
	println(getSum(2, 3))
	println(getSum(34, 56))
}

//leetcode submit region end(Prohibit modification and deletion)
