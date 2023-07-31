from typing import List

from typing import List

class Solution:
    # Function to detect cycle in an undirected graph.
    def isCycle(self, V: int, adj: List[List[int]]) -> bool:
        visited = [False] * V

        def dfs(node, prev):
            # Base case
            if visited[node]:
                return False

            # Mark the current node as visited
            visited[node] = True

            # Explore the neighbors (adjacent nodes)
            for nei in adj[node]:
                if nei == prev:
                    continue
                if not dfs(nei, node):
                    return False

            return True

        # Perform DFS for each unvisited node
        for i in range(V):
            if not visited[i]:
                if not dfs(i, -1):
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