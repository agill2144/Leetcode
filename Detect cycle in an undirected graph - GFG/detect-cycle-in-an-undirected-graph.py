from typing import List

class Solution:
    # Function to detect cycle in an undirected graph.
    def isCycle(self, V: int, adj: List[List[int]]) -> bool:
        visited = [False] * V
        q = []  # <node, prev>

        for i in range(V):
            if not visited[i]:
                visited[i] = True
                q.append([i, -1])
                while len(q) != 0:
                    dq = q[0]
                    q = q[1:]
                    currNode = dq[0]
                    prev = dq[1]

                    for nei in adj[currNode]:
                        if nei == prev:
                            continue
                        if visited[nei]:
                            return True

                        visited[nei] = True
                        q.append([nei, currNode])

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