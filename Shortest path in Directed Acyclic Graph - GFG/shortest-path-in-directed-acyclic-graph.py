import heapq
import math
from typing import List

class Solution:
    def shortestPath(self, n: int, m: int, edges: List[List[int]]) -> List[int]:
        adjList = {}
        for i in range(len(edges)):
            u, v, w = edges[i][0], edges[i][1], edges[i][2]
            if u not in adjList:
                adjList[u] = []
            adjList[u].append([v, w])
        
        src = 0
        dist = [math.inf] * n
        dist[src] = 0
        pq = [(0, 0)]  # min heap
        
        while pq:
            currNode, currDist = heapq.heappop(pq)
            
            if currDist > dist[currNode]:
                continue
            
            if currNode in adjList:
                for nei in adjList[currNode]:
                    adjNode, adjDist = nei[0], currDist + nei[1]
                    if adjDist < dist[adjNode]:
                        dist[adjNode] = adjDist
                        heapq.heappush(pq, (adjNode, adjDist))
        
        for i in range(len(dist)):
            if dist[i] == math.inf:
                dist[i] = -1
        
        return dist



#{ 
 # Driver Code Starts
#Initial Template for Python 3

from typing import List




class IntMatrix:
    def __init__(self) -> None:
        pass
    def Input(self,n,m):
        matrix=[]
        #matrix input
        for _ in range(n):
            matrix.append([int(i) for i in input().strip().split()])
        return matrix
    def Print(self,arr):
        for i in arr:
            for j in i:
                print(j,end=" ")
            print()



class IntArray:
    def __init__(self) -> None:
        pass
    def Input(self,n):
        arr=[int(i) for i in input().strip().split()]#array input
        return arr
    def Print(self,arr):
        for i in arr:
            print(i,end=" ")
        print()


if __name__=="__main__":
    t = int(input())
    for _ in range(t):
        
        n,m=map(int,input().split())
        
        
        edges=IntMatrix().Input(m, 3)
        
        obj = Solution()
        res = obj.shortestPath(n, m, edges)
        
        IntArray().Print(res)
# } Driver Code Ends