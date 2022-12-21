//[212单词搜索II]
//给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
//
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使
//用。
//
//
//
// 示例 1：
//
//
//输入：board = [["o","a","a","n"],
//				["e","t","a","e"],
//				["i","h","k","r"],
//				["i","f","l","v"]],
//words = ["oath","pea","eat","rain"]
//输出：["eat","oath"]
//
//
// 示例 2：
//
//
//输入：board = [["a","b"],["c","d"]], words = ["abcb"]
//输出：[]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] 是一个小写英文字母
// 1 <= words.length <= 3 * 10⁴
// 1 <= words[i].length <= 10
// words[i] 由小写英文字母组成
// words 中的所有字符串互不相同
//
//
// Related Topics 字典树 数组 字符串 回溯 矩阵 👍 731 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

import "encoding/json"

type Trie struct {
	nextNodes [26]*Trie
	word      string
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		c -= 'a'
		if cur.nextNodes[c] == nil {
			cur.nextNodes[c] = &Trie{}
		}
		cur = cur.nextNodes[c]
	}
	cur.word = word
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	for _, w := range words {
		root.Insert(w)
	}

	findMap := make(map[string]bool, 0)
	m, n := len(board), len(board[0])
	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		c := board[x][y]
		if c == '0' {
			return
		}

		node = node.nextNodes[c-'a']
		if node == nil {
			return
		}
		if node.word != "" {
			findMap[node.word] = true
		}

		board[x][y] = '0'
		dfs(node, x, y+1)
		dfs(node, x, y-1)
		dfs(node, x+1, y)
		dfs(node, x-1, y)
		board[x][y] = c
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(root, i, j)
		}
	}

	result := make([]string, 0, len(findMap))
	for w, _ := range findMap {
		result = append(result, w)
	}

	return result
}

func main() {
	//board := [][]byte{
	//	{'o', 'a', 'a', 'n'},
	//	{'e', 't', 'a', 'e'},
	//	{'i', 'h', 'k', 'r'},
	//	{'i', 'f', 'l', 'v'},
	//}
	//words := []string{"oath", "pea", "eat", "rain"}
	board := [][]byte{{'a', 'a'}}
	words := []string{"aa"}
	result := findWords(board, words)
	jsonBytes, _ := json.Marshal(result)
	println(string(jsonBytes))
}

//leetcode submit region end(Prohibit modification and deletion)
