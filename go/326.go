//[326]3 的幂
//给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
//
// 整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3ˣ
//
//
//
// 示例 1：
//
//
//输入：n = 27
//输出：true
//
//
// 示例 2：
//
//
//输入：n = 0
//输出：false
//
//
// 示例 3：
//
//
//输入：n = 9
//输出：true
//
//
// 示例 4：
//
//
//输入：n = 45
//输出：false
//
//
//
//
// 提示：
//
//
// -2³¹ <= n <= 2³¹ - 1
//
//
//
//
// 进阶：你能不使用循环或者递归来完成本题吗？
//
// Related Topics 递归 数学 👍 289 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func isPowerOfThree(n int) bool {
	//temp := 1
	//for temp < n {
	//	temp *= 3
	//}
	//return temp == n

	//最大的 3 的幂为 3^19=1162261467。我们只需要判断 n 是否是 3^19 的约数即可
	return n > 0 && 1162261467%n == 0
}

func printByte(n int32) {
	for i := 31; i >= 0; i-- {
		print((n >> i) & 1)
	}

	println()
}

func main() {
	//println(isPowerOfThree(81))
	//println(isPowerOfThree(100))
	printByte(1)
	printByte(3)
	printByte(3 * 3)
	printByte(3 * 3 * 3)
	printByte(3 * 3 * 3 * 3)

}

//leetcode submit region end(Prohibit modification and deletion)
