#User function Template for python3

from typing import List
import math

class Solution:
    def shortestPath(self, n, m, edges):
        adjList = {}
        for i in range(len(edges)):
            src = edges[i][0]
            dest = edges[i][1]
            wt = edges[i][2]
            if src not in adjList:
                adjList[src] = []
            adjList[src].append([dest, wt])
        
        visited = [False] * n
        src = 0
        dist = [math.inf] * n
        def dfs(node, currWt):
            # base
            if visited[node] and currWt >= dist[node]:
                return
            
            # logic
            dist[node] = currWt
            visited[node] = True
            if node in adjList:
                for nei in adjList[node]:
                    dfs(nei[0], currWt+nei[1])
        
        dfs(src, 0)
        for i in range(n):
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