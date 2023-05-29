#User function Template for python3


class Solution:
    
    #Function to detect cycle in a directed graph.
    # def isCyclic(self, v, adj):
    #     # code here
    #     indegrees = [0] * v
    #     for i in range(len(adj)):
    #         for child in adj[i]:
    #             indegrees[child] += 1
        
    #     visited = [False] * v
        
    #     q = []
    #     for i in range(len(indegrees)):
    #         if indegrees[i] == 0:
    #             q.append(i)
    #             visited[i] = True
        
    #     count = 0
    #     while len(q) != 0:
    #         dq = q[0]
    #         q = q[1:]
    #         count += 1
    #         for nei in adj[dq]:
    #             indegrees[nei] -= 1
    #             if indegrees[nei] == 0 and not visited[nei]:
    #                 q.append(nei)
        
    #     if count == v:
    #         return False
        
    #     return True
    
    
    def isCyclic(self, v, adj):
        visited = [False] * v
        
        def dfs(node, path):
            # base
            if path[node]:
                return False
            if visited[node]:
                return True
        
            # logic
            path[node] = True
            visited[node] = True
            for nei in adj[node]:
                if not dfs(nei, path):
                    return False
            path[node] = False
            return True
        
        path = [False] * v
        for i in range(v):
            if not visited[i]:
                if not dfs(i, path):
                    return True
        
        return False



#{ 
 # Driver Code Starts
#Initial Template for Python 3

import sys
sys.setrecursionlimit(10**6)
        
if __name__ == '__main__':
    t = int(input())
    for i in range(t):
        V,E = list(map(int, input().strip().split()))
        adj = [[] for i in range(V)]
        for i in range(E):
            a,b = map(int,input().strip().split())
            adj[a].append(b)
        ob = Solution()
        
        if ob.isCyclic(V, adj):
            print(1)
        else:
            print(0)
# } Driver Code Ends