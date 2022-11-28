# 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列
#  beginWord -> s1 -> s2 -> ... -> sk：
#
#
#  每一对相邻的单词只差一个字母。
#
#  对于 1 <= i <= k 时，每个
#  si 都在
#  wordList 中。注意， beginWord 不需要在
#  wordList 中。
#
#  sk == endWord
#
#
#  给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列
# 中的 单词数目 。如果不存在这样的转换序列，返回 0 。
#
#  示例 1：
#
#
# 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot",
# "log","cog"]
# 输出：5
# 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
#
#
#  示例 2：
#
#
# 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot",
# "log"]
# 输出：0
# 解释：endWord "cog" 不在字典中，所以无法进行转换。
#
#
#
#  提示：
#
#
#  1 <= beginWord.length <= 10
#  endWord.length == beginWord.length
#  1 <= wordList.length <= 5000
#  wordList[i].length == beginWord.length
#  beginWord、endWord 和 wordList[i] 由小写英文字母组成
#  beginWord != endWord
#  wordList 中的所有字符串 互不相同
#
#
#  Related Topics 广度优先搜索 哈希表 字符串 👍 1099 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
# class Solution(object):
#     def ladderLength(self, beginWord, endWord, wordList):
#         """
#         :type beginWord: str
#         :type endWord: str
#         :type wordList: List[str]
#         :rtype: int
#         """
#
#         used_set = set()
#         used_set.add(beginWord)
#         self.temp_len = 1
#         self.min_len = float('inf')
#         match_len = len(beginWord)
#
#         def compare(word_1, word_2):
#             count = 0
#             for i in range(match_len):
#                 if word_1[i] != word_2[i]:
#                     count += 1
#             return count <= 1
#
#         def bfs(begin_w):
#             for w in wordList:
#                 if w in used_set:
#                     continue
#
#                 if compare(begin_w, w):
#                     if w == endWord:
#                         self.min_len = min(self.min_len, self.temp_len + 1)
#                         return True
#
#                     used_set.add(w)
#                     self.temp_len += 1
#
#                     bfs(w)
#
#                     used_set.remove(w)
#                     self.temp_len -= 1
#
#             return False
#
#         bfs(beginWord)
#         return self.min_len if self.min_len < float('inf') else 0

class Solution(object):
    def __init__(self):
        self.id_num = 0

    def ladderLength(self, beginWord, endWord, wordList):
        """
        :type beginWord: str
        :type endWord: str
        :type wordList: List[str]
        :rtype: int
        """

        word_id_map = {}
        edges = {}

        def get_word_id(w):
            if w not in word_id_map:
                word_id_map[w] = self.id_num
                self.id_num += 1
            return word_id_map[w]

        def add_edge(w):
            id = get_word_id(w)
            if id not in edges: edges[id] = []
            chars = list(w)
            for i in range(len(chars)):
                c = chars[i]
                chars[i] = '*'
                new_word = ''.join(chars)
                new_id = get_word_id(new_word)
                edges[id].append(new_id)
                if new_id not in edges: edges[new_id] = []
                edges[new_id].append(id)
                chars[i] = c

        add_edge(beginWord)
        for w in wordList:
            add_edge(w)
        if endWord not in word_id_map:
            return 0

        begin_id, end_id = word_id_map[beginWord], word_id_map[endWord]
        max_float = float('inf')
        distance = [max_float] * self.id_num
        queue = [begin_id]
        distance[begin_id] = 0
        while queue:
            id = queue.pop(0)
            if id == end_id:
                return distance[id] // 2 + 1

            for edge_id in edges[id]:
                if distance[edge_id] == max_float:
                    distance[edge_id] = distance[id] + 1
                    queue.append(edge_id)

        return 0


# print(Solution().ladderLength(beginWord="hit", endWord="cog", wordList=["hot", "dot", "dog", "lot", "log", "cog"]))
# print(Solution().ladderLength(beginWord="lost", endWord="miss",
#                               wordList=["most", "mist", "miss", "lost", "fist", "fish"]))
print(Solution().ladderLength(beginWord="leet", endWord="code",
                              wordList=["lest", "leet", "lose", "code", "lode", "robe", "lost"]))

# leetcode submit region end(Prohibit modification and deletion)
