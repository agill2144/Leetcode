#User function Template for python3
import heapq
import math
from typing import List



import math
import heapq
from typing import List

class Solution:
    def shortestPath(self, n: int, m: int, edges: List[List[int]]) -> List[int]:
        adjList = {}  # { parent : [[node, dist]] }
        for i in range(len(edges)):
            u, v, w = edges[i][0], edges[i][1], edges[i][2]
            if u not in adjList:
                adjList[u] = []
            if v not in adjList:
                adjList[v] = []
            adjList[u].append([v, w])
            adjList[v].append([u, w])

        start = 1
        end = n
        dist = [math.inf] * (n + 1)

        dist[start] = 0
        pq = []
        heapq.heappush(pq, (0, [start]))

        while pq:
            dq = heapq.heappop(pq)
            path = dq[1]
            currNode = path[-1]
            currDist = dq[0]

            if currNode == end:
                return path
            
            if currNode in adjList:
                for nei in adjList[currNode]:
                    adjNode = nei[0]
                    adjNodeDist = nei[1] + currDist
                    if adjNodeDist < dist[adjNode]:
                        newPath = path.copy()
                        newPath.append(adjNode)
                        heapq.heappush(pq, (adjNodeDist, newPath))
                        dist[adjNode] = adjNodeDist

        return [-1]

#{ 
 # Driver Code Starts
#Initial Template for Python 3

if __name__ == '__main__': 
    t = int(input ())
    for _ in range (t):
        n, m = list(map(int, input().split()))
        edges = []
        for i in range(m):
            node1, node2, weight = list(map(int, input().split()))
            edges.append([node1, node2, weight])
        obj = Solution()
        ans = obj.shortestPath(n, m, edges)
        for e in ans:
            print(e, end = ' ')
        print()
# } Driver Code Ends