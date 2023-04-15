#User function Template for python3

from typing import List
import math

class Solution:
    def shortestPath(self, n : int, m : int, edges : List[List[int]]) -> List[int]:
        adjList = {}
        for i in range(len(edges)):
            src = edges[i][0]
            dest = edges[i][1]
            dist = edges[i][2]
            if src in adjList:
                adjList[src].append([dest, dist])
            else:
                adjList[src] = [[dest, dist]]
        out = [0]*n
    
        visited = {0:0}
        def dfs(node, dest, dist):
            # base
            if node == dest:
                return dist
            if node in visited:
                if visited[node] < dist:
                    return math.inf
            # logic
            minDist = math.inf
            visited[node] = dist
            for nei in adjList.get(node, []):
                minDist = min(dfs(nei[0], dest, dist+nei[1]), minDist)
            return minDist
    
        for i in range(n):
            out[i] = dfs(0, i, 0)
            if out[i] == math.inf:
                out[i] = -1
        return out
    
    def min(x, y):
        if x < y:
            return x
        return y
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