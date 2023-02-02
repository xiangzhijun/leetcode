//[171]Excel表序列号
//给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。
//
// 例如：
//
//
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//
//
//
// 示例 1:
//
//
//输入: columnTitle = "A"
//输出: 1
//
//
// 示例 2:
//
//
//输入: columnTitle = "AB"
//输出: 28
//
//
// 示例 3:
//
//
//输入: columnTitle = "ZY"
//输出: 701
//
//
//
// 提示：
//
//
// 1 <= columnTitle.length <= 7
// columnTitle 仅由大写英文组成
// columnTitle 在范围 ["A", "FXSHRXW"] 内
//
//
// Related Topics 数学 字符串 👍 359 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

func titleToNumber(columnTitle string) int {
	num, base := 0, 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		//字符'A'的ascii码是65
		num += (int(columnTitle[i]) - 64) * base
		base *= 26
	}
	return num
}

func main() {
	println(titleToNumber("AB"))
	println(titleToNumber("ZY"))
}

//leetcode submit region end(Prohibit modification and deletion)
