package org.zhijun.leetcode;


//[394]字符串解码
//给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
//
//
// 示例 1：
//
//
//输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//
//
// 示例 2：
//
//
//输入：s = "3[a2[c]]"
//输出："accaccacc"
//
//
// 示例 3：
//
//
//输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
//
//
// 示例 4：
//
//
//输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 30
//
// s 由小写英文字母、数字和方括号
// '[]' 组成
// s 保证是一个 有效 的输入。
// s 中所有整数的取值范围为
// [1, 300]
//
//
// Related Topics 栈 递归 字符串 👍 1467 👎 0


import java.util.Stack;

//leetcode submit region begin(Prohibit modification and deletion)
class J394 {
    public String decodeString(String s) {
        Stack<String> stack = new Stack<>();
        int i = 0;
        while (i < s.length()) {
            char v = s.charAt(i);
            if (v == '[') {
                stack.push("[");
                i++;
            } else if (v == ']') {
                StringBuilder builder = new StringBuilder();
                while (true) {
                    String temp = stack.pop();
                    if (temp.charAt(temp.length() - 1) == '[') {
                        int count = Integer.parseInt(temp.substring(0, temp.length() - 1));
                        StringBuilder repeat = new StringBuilder();
                        for (int j = 0; j < count; j++) {
                            repeat.append(builder);
                        }
                        stack.push(repeat.toString());
                        break;
                    }
                    builder.insert(0, temp);
                }
                i++;
            } else if (v >= '0' && v <= '9') {
                int j = i + 1;
                while (s.charAt(j) != '[') {
                    j++;
                }
                stack.push(s.substring(i, j + 1));
                i = j + 1;
            } else {
                int j = i + 1;
                while (j < s.length() && s.charAt(j) >= 'a' && s.charAt(j) <= 'z') {
                    j++;
                }
                stack.push(s.substring(i, j));
                i = j;
//                stack.push(s.substring(i, i + 1));
//                i++;
            }
        }

        StringBuilder resultBuilder = new StringBuilder();
        while (!stack.empty()) {
            resultBuilder.insert(0, stack.pop());
        }
        return resultBuilder.toString();
    }

    public static void main(String[] args) {
        J394 solution = new J394();
//        assert "aaabcbc".equals(solution.decodeString("3[a]2[bc]"));
        System.out.println(solution.decodeString("2[abc]3[cd]ef"));
    }
}
//leetcode submit region end(Prohibit modification and deletion)

