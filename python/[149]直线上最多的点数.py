# 给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。
#
#
#
#  示例 1：
#
#
# 输入：points = [[1,1],[2,2],[3,3]]
# 输出：3
#
#
#  示例 2：
#
#
# 输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
# 输出：4
#
#
#
#
#  提示：
#
#
#  1 <= points.length <= 300
#  points[i].length == 2
#  -10⁴ <= xi, yi <= 10⁴
#  points 中的所有点 互不相同
#
#
#  Related Topics 几何 数组 哈希表 数学 👍 435 👎 0


# leetcode submit region begin(Prohibit modification and deletion)
class Solution(object):
    def maxPoints(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """
        n = len(points)
        max_len = 1
        for i in range(n):
            w_b_map = {}
            for j in range(i + 1, n):
                p1, p2 = points[i], points[j]
                if p1[0] == p2[0]:
                    w = float('inf')
                    b = p1[0]
                else:
                    w = float(p1[1] - p2[1]) / float(p1[0] - p2[0])
                    b = p1[1] - w * p1[0]
                if w in w_b_map:
                    if b in w_b_map[w]:
                        w_b_map[w][b] += 1
                    else:
                        w_b_map[w][b] = 2
                else:
                    w_b_map[w] = {b: 2}
                max_len = max(max_len, w_b_map[w][b])

        return max_len


# print(Solution().maxPoints([[1, 1], [3, 2], [5, 3], [4, 1], [2, 3], [1, 4]]))
# print(Solution().maxPoints([[0, 0], [4, 5], [7, 8], [8, 9], [5, 6], [3, 4], [1, 1]]))
# leetcode submit region end(Prohibit modification and deletion)
