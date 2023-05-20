#User function Template for python3

from typing import List
import math

class Solution:
    def shortestPath(self, n : int, m : int, edges : List[List[int]]) -> List[int]:
        adjList = {}
        for i in range(len(edges)):
            src = edges[i][0]
            dest = edges[i][1]
            wt = edges[i][2]
            if src in adjList:
                adjList[src].append([dest, wt])
            else:
                adjList[src] = [[dest, wt]]
        
        visited = [False] * n
        st = []
        
        def dfs(node):
            if visited[node]:
                return
            
            visited[node] = True
            if node in adjList:
                for nei in adjList[node]:
                    dfs(nei[0])
            
            st.append(node)
        
        for i in range(n):
            if not visited[i]:
                dfs(i)
        
        src = 0
        while len(st) != 0 and st[-1] != src:
            st.pop()
        
        dist = [math.inf] * n
        dist[src] = 0
        while len(st) != 0:
            top = st[-1]
            st.pop()
            currDist = dist[top]
            if top in adjList:
                for neiList in adjList[top]:
                    nei = neiList[0]
                    neiDist = neiList[1]
                    newDist = currDist + neiDist
                    if newDist < dist[nei]:
                        dist[nei] = newDist
        
        
        for idx, ele in enumerate(dist):
            if ele == math.inf:
                dist[idx] = -1
        
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