from typing import List
class Solution:
    #Function to detect cycle in an undirected graph.
	def isCycle(self, V: int, adj: List[List[int]]) -> bool:
	    visited = [False]*V
	    
	    def dfs(node: int, prev: int, path: List[bool]) -> bool:
	        # base
	        if path[node]:
	           return False
	        if visited[node]:
	            return True
	        
	        # logic
	        path[node] = True
	        for nei in adj[node]:
	            if nei == prev:
	                continue
	            if not dfs(nei, node, path):
	                return False
	        path[node] = False
	        return True
	   
	    p = [False]*V
	    for i in range(V):
	        if not visited[i]:
	            if not dfs(i,-1, p):
	                return True
	   
	    return False
	   
	       
	       


#{ 
 # Driver Code Starts
if __name__ == '__main__':

	T=int(input())
	for i in range(T):
		V, E = map(int, input().split())
		adj = [[] for i in range(V)]
		for _ in range(E):
			u, v = map(int, input().split())
			adj[u].append(v)
			adj[v].append(u)
		obj = Solution()
		ans = obj.isCycle(V, adj)
		if(ans):
			print("1")
		else:
			print("0")

# } Driver Code Ends