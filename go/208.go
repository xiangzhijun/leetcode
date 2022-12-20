//[208]实现 Trie（前缀树)
//Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼
//写检查。
//
// 请你实现 Trie 类：
//
//
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回
//false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否
//则，返回 false 。
//
//
//
//
// 示例：
//
//
//输入
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//输出
//[null, null, true, false, true, null, true]
//
//解释
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // 返回 True
//trie.search("app");     // 返回 False
//trie.startsWith("app"); // 返回 True
//trie.insert("app");
//trie.search("app");     // 返回 True
//
//
//
//
// 提示：
//
//
// 1 <= word.length, prefix.length <= 2000
// word 和 prefix 仅由小写英文字母组成
// insert、search 和 startsWith 调用次数 总计 不超过 3 * 10⁴ 次
//
//
// Related Topics 设计 字典树 哈希表 字符串 👍 1351 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
package main

type Node struct {
	pass     int
	end      int
	nextList [26]*Node
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: &Node{}}
}

func (this *Trie) Insert(word string) {
	this.root.pass++
	cur := this.root
	var index int32 = 0
	for _, c := range word {
		index = c - 'a'
		if cur.nextList[index] == nil {
			cur.nextList[index] = &Node{}
		}
		cur = cur.nextList[index]
		cur.pass++
	}
	cur.end++
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	var index int32 = 0
	for _, c := range word {
		index = c - 'a'
		if cur.nextList[index] == nil {
			return false
		}
		cur = cur.nextList[index]
	}
	return cur.end > 0
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	var index int32 = 0
	for _, c := range prefix {
		index = c - 'a'
		if cur.nextList[index] == nil {
			return false
		}
		cur = cur.nextList[index]
	}
	return cur.pass > 0
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	println(trie.Search("apple"))   // 返回 True
	println(trie.Search("app"))     // 返回 False
	println(trie.StartsWith("app")) // 返回 True
	trie.Insert("app")
	println(trie.Search("app")) // 返回 True
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
