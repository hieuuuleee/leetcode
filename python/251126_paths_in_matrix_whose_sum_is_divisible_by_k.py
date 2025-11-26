from collections import defaultdict
from typing import List

class Solution:
    def numberOfPaths(self, grid: List[List[int]], k: int) -> int:
        m = len(grid)
        n = len(grid[0])
        dp = [[defaultdict(int) for _ in range(n+1)] for _ in range(m+1)]
        dp[1][1][grid[0][0]%k] += 1

        for r in range(m):
            for c in range(n):
                if r or c:
                    for remain in range(k):
                        cur_remain = (remain + grid[r][c]) % k
                        dp[r+1][c+1][cur_remain] += (dp[r][c+1][remain] + dp[r+1][c][remain])
        return dp[m][n][0]

if __name__ == "__main__":
    solution = Solution()
    print(solution.numberOfPaths([[5,2,4],[3,0,5],[0,7,2]], 3))
