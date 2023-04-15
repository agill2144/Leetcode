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
    
        def bfs(dest):
            visited = {0:0}
            # <node, dist>
            minDist = math.inf
            q = [[0,0]]
            while len(q) != 0:
                dq = q[0]
                q = q[1:]
                node = dq[0]
                dist = dq[1]
                if node == dest:
                    if dist < minDist:
                        minDist = dist
                for nei in adjList.get(node, []):
                    visitedDist = visited.get(nei[0])
                    newDist = nei[1] + dist
                    if visitedDist is None or newDist < visitedDist:
                        visited[nei[0]] = newDist
                        q.append([nei[0], newDist])
            if minDist == math.inf:
                return -1
            return minDist
    
        for i in range(n):
            out[i] = bfs(i)
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