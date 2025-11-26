from collections import deque
from typing import List

INF = 1000000001

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        union = {}
        exist = set([])
        for n in nums:
            if n in exist:
                continue
            else:
                exist.add(n)
                union[n] = [INF,INF]
            if n-1 in exist:
                union[n-1][1] = n
                union[n][0] = n-1
            if n+1 in exist:
                union[n+1][0] = n
                union[n][1] = n+1
        
        check = set([])
        res = 0
        for n in nums:
            if n in check:
                continue
            que = deque()
            que.append(n)
            cur_len = 0
            while que:
                cur = que.popleft()
                if cur in check:
                    continue
                check.add(cur)
                cur_len += 1
                if union[cur][0]!=INF:
                    que.append(union[cur][0])
                if union[cur][1]!=INF:
                    que.append(union[cur][1])
            res = max(res, cur_len)
        
        return res

    # def longestConsecutive(self, nums: List[int]) -> int:
    #     num_set = set(nums)
    #     longest = 0

    #     for n in num_set:
    #         if n - 1 not in num_set:
    #             length = 1

    #             while n + length in num_set:
    #                 length += 1
                
    #             longest = max(longest, length)
        
    #     return longest
                

if __name__ == "__main__":
    solution = Solution()
    print(solution.longestConsecutive([100,4,200,1,3,2]))
