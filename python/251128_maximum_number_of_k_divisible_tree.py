class Solution:
    def maxKDivisibleComponents(self, n: int, edges: List[List[int]], values: List[int], k: int) -> int:
        graph = defaultdict(list)
        count = defaultdict(int)
        for edge in edges:
            graph[edge[0]].append(edge[1])
            graph[edge[1]].append(edge[0])
            count[edge[0]] += 1
            count[edge[1]] += 1
        
        stack = []
        for i in range(n):
            if count[i] == 1:
                stack.append(i)
        
        res = 1
        check = defaultdict(int)
        while stack:
            cur = stack.pop()
            
            if check[cur] or count[cur] < 1:
                continue

            for nxt in graph[cur]:
                if not check[nxt]:
                    count[nxt] -= 1
                    if count[nxt] == 1:
                        stack.append(nxt)
                    
                    if not values[cur] % k:
                        res += 1
                    else:
                        values[nxt] += values[cur]

            check[cur] = 1

        return res
