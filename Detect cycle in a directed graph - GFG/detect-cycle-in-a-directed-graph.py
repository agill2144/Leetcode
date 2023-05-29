#User function Template for python3


class Solution:
    
    #Function to detect cycle in a directed graph.
    def isCyclic(self, v, adj):
        # code here
        indegrees = [0] * v
        for i in range(len(adj)):
            for child in adj[i]:
                indegrees[child] += 1
        
        visited = [False] * v
        
        q = []
        for i in range(len(indegrees)):
            if indegrees[i] == 0:
                q.append(i)
                visited[i] = True
        
        count = 0
        while len(q) != 0:
            dq = q[0]
            q = q[1:]
            count += 1
            for nei in adj[dq]:
                indegrees[nei] -= 1
                if indegrees[nei] == 0 and not visited[nei]:
                    q.append(nei)
        
        if count == v:
            return False
        
        return True


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