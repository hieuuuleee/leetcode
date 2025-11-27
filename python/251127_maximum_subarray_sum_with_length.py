from typing import List
INF = 10**9 + 1
class Solution:
    def maxSubarraySum(self, nums: List[int], k: int) -> int:
        window = 0
        n = len(nums)
        dp = [0 for i in range(n)]
        for i in range(k):
            window += nums[i]
        
        dp[k-1] = window
        res = window
        for i in range(k,n):
            window = window-nums[i-k]+nums[i]
            dp[i] = max(window, window + dp[i-k])
            res = max(res, dp[i])
            print(window, dp[i], res)
        
        print(dp)

        return res

if __name__ == "__main__":
    solution = Solution()
    print(solution.maxSubarraySum([-1,-2,-3,-4,-5], 4))
