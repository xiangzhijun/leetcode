//[166]分数到小数
//给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以 字符串形式返回小数 。
//
// 如果小数部分为循环小数，则将循环的部分括在括号内。
//
// 如果存在多个答案，只需返回 任意一个 。
//
// 对于所有给定的输入，保证 答案字符串的长度小于 10⁴ 。
//
//
//
// 示例 1：
//
//
//输入：numerator = 1, denominator = 2
//输出："0.5"
//
//
// 示例 2：
//
//
//输入：numerator = 2, denominator = 1
//输出："2"
//
//
// 示例 3：
//
//
//输入：numerator = 4, denominator = 333
//输出："0.(012)"
//
//
//
//
// 提示：
//
//
// -2³¹ <= numerator, denominator <= 2³¹ - 1
// denominator != 0
//
//
// Related Topics 哈希表 数学 字符串 👍 424 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	numerator64, denominator64 := int64(numerator), int64(denominator)
	//能整除，直接返回结果
	if numerator64%denominator64 == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	temp := ""
	indexMap := make(map[int64]int)
	//判断结果是否为负数
	if numerator64*denominator64 < 0 {
		temp += "-"
	}

	//绝对值函数
	abs := func(num int64) int64 {
		if num < 0 {
			return num * -1
		}
		return num
	}
	numerator64, denominator64 = abs(numerator64), abs(denominator64)
	//商整数部分
	temp += strconv.FormatInt(numerator64/denominator64, 10) + "."
	numerator64 %= denominator64
	for numerator64 != 0 {
		//保存当前小数第一次出现的位置
		indexMap[numerator64] = len(temp)
		//余数补0后继续除
		numerator64 *= 10
		temp += strconv.FormatInt(numerator64/denominator64, 10)
		numerator64 %= denominator64

		index, ok := indexMap[numerator64]
		//当前余数出现过，代表结果是无限循环小数，直接返回
		if ok {
			return fmt.Sprintf("%s(%s)", temp[:index], temp[index:])
		}
	}
	//返回有限小数
	return temp
}

func main() {
	println(fractionToDecimal(1, 2))
	println(fractionToDecimal(2, 1))
	println(fractionToDecimal(4, 333))
}

//leetcode submit region end(Prohibit modification and deletion)
