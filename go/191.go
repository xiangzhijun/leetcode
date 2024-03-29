//[191]位1的个数
//编写一个函数，输入是一个无符号整数（以二进制串的形式），返回其二进制表达式中数字位数为 '1' 的个数（也被称为汉明重量）。
//
//
//
// 提示：
//
//
// 请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的
//还是无符号的，其内部的二进制表示形式都是相同的。
// 在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。
//
//
//
//
// 示例 1：
//
//
//输入：00000000000000000000000000001011
//输出：3
//解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。
//
//
// 示例 2：
//
//
//输入：00000000000000000000000010000000
//输出：1
//解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。
//
//
// 示例 3：
//
//
//输入：11111111111111111111111111111101
//输出：31
//解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。
//
//
//
// 提示：
//
//
// 输入必须是长度为 32 的 二进制串 。
//
//
//
//
//
//
//
// 进阶：
//
//
// 如果多次调用这个函数，你将如何优化你的算法？
//
//
// Related Topics 位运算 分治 👍 539 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

//func hammingWeight(num uint32) int {
//	count := 0
//	for num > 0 {
//		//n&(n-1)表示把二进制最低位的1变成0，循环此操作直到n=0，统计变换次数
//		num &= num - 1
//		count++
//	}
//	return count
//}

func hammingWeight(num uint32) int {
	n := num
	//n & 0x55555555：只保留奇数位；(n >> 1) & 0x55555555：偶数位向后移一位跟奇数位对齐
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)  //相邻1位相加，奇数位和偶数位相加
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)  //每相邻2位相加
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)  //每相邻4位相加
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)  //每相邻8位相加
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff) //每相邻16位相加

	return int(n)
}

func printBit(n uint32) {
	for i := 31; i >= 0; i-- {
		print((n >> i) & 1)
	}
	println()
}

func main() {
	println(hammingWeight(43261596))
	println(hammingWeight(4294967293))
}

//leetcode submit region end(Prohibit modification and deletion)
