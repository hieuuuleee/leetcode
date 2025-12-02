from typing import List

INF = 10**9+7

class Solution:
    def countTrapezoids(self, points: List[List[int]]) -> int:
        y_set = set([])
        y_cnt = {}
        
        for p in points:
            if not p[1] in y_set:
                y_set.add(p[1])
                y_cnt[p[1]] = 0
            y_cnt[p[1]] += 1

        print(y_set)
        print(y_cnt)

        if len(y_set) < 2:
            return 0
            
        if len(y_set) == 2:
            for y in y_set:
                if y_cnt[y] == 1:
                    return 0
        
        res = 0
        cnt_sequences = 0
        for y in y_set:
            if y_cnt[y] == 1:
                continue
            cur_cnt = int((1 + y_cnt[y] - 1)*(y_cnt[y] - 1)/2) % INF
            res = (res + cnt_sequences * cur_cnt) % INF
            cnt_sequences = (cnt_sequences + cur_cnt) % INF
            print(res)

        return res

if __name__ == "__main__":
    print(Solution().countTrapezoids([[1,0],[2,0],[3,0],[2,2],[3,2]]))
