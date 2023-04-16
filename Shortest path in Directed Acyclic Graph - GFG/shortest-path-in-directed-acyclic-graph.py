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
            if src not in adjList:
                adjList[src] = []
            adjList[src].append([dest, dist])
        
        src = 0
        out = [math.inf] * n
        out[src] = 0
        q = [[0, 0]]
        while q:
            node, dist = q.pop(0)
            for nei in adjList.get(node, []):
                neiNode = nei[0]
                neiDist = nei[1]
                newDist = dist + neiDist
                if newDist < out[neiNode]:
                    out[neiNode] = newDist
                    q.append([neiNode, newDist])
        
        # if there were unreachable nodes
        # update their min dist to -1 as it denotes not-reachable.
        for i in range(n):
            if out[i] == math.inf:
                out[i] = -1
        
        return out
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