#
#  请你设计并实现一个满足
#  LRU (最近最少使用) 缓存 约束的数据结构。
#
#
#
#  实现
#  LRUCache 类：
#
#
#
#
#
#  LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
#  int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
#  void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
# key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
#
#
#
#
#  函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
#
#
#
#  示例：
#
#
# 输入
# ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
# [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
# 输出
# [null, null, null, 1, null, -1, null, -1, 3, 4]
#
# 解释
# LRUCache lRUCache = new LRUCache(2);
# lRUCache.put(1, 1); // 缓存是 {1=1}
# lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
# lRUCache.get(1);    // 返回 1
# lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
# lRUCache.get(2);    // 返回 -1 (未找到)
# lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
# lRUCache.get(1);    // 返回 -1 (未找到)
# lRUCache.get(3);    // 返回 3
# lRUCache.get(4);    // 返回 4
#
#
#
#
#  提示：
#
#
#  1 <= capacity <= 3000
#  0 <= key <= 10000
#  0 <= value <= 10⁵
#  最多调用 2 * 10⁵ 次 get 和 put
#
#
#  Related Topics 设计 哈希表 链表 双向链表 👍 2358 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class LRUCache(object):
    class Node(object):
        def __init__(self, key, val, pre=None, next=None):
            self.key = key
            self.val = val
            self.pre = pre
            self.next = next

    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self._cache_map = {}
        self._capacity = capacity
        self._head = None
        self._tail = None

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key in self._cache_map:
            node = self._cache_map[key]
            self._delete_node(node)
            self._add_node(node)
            return node.val
        else:
            return -1

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: None
        """
        if key in self._cache_map:
            node = self._cache_map[key]
            node.val = value
            self._delete_node(node)
            self._add_node(node)
        else:
            if len(self._cache_map) >= self._capacity:
                self._cache_map.pop(self._head.key)
                self._delete_node(self._head)
            node = self.Node(key, value)
            self._cache_map[key] = node
            self._add_node(node)

    def _add_node(self, node):
        if self._tail is None:
            self._tail = node
            self._head = node
        else:
            self._tail.next = node
            node.pre = self._tail
            self._tail = node

    def _delete_node(self, node):
        if node == self._head:
            self._head = node.next
        else:
            node.pre.next = node.next

        if node == self._tail:
            self._tail = node.pre
        else:
            node.next.pre = node.pre

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)

# obj = LRUCache(2)
# print(obj.put(1, 1))
# print(obj.put(2, 2))
# print(obj.get(1))
# print(obj.put(3, 3))
# print(obj.get(2))
# print(obj.put(4, 4))
# print(obj.get(1))
# print(obj.get(3))
# print(obj.get(4))

# leetcode submit region end(Prohibit modification and deletion)
