#User function Template for python3
import math
class Solution:

    def shortestPath(self, edges, n, m, src):
        adjList = {}
        for s, dest in edges:
            adjList.setdefault(s, []).append(dest)
            adjList.setdefault(dest, []).append(s)
    
        out = [math.inf] * n
        out[src] = 0
        q = [[src, -1, 0]]  # < node, prev, dist >
    
        while q:
            node, prev, currDist = q.pop(0)
            for nei in adjList.get(node, []):
                # if nei == prev:
                #     continue
                newDist = currDist + 1
                if out[nei] == math.inf or newDist < out[nei]:
                    out[nei] = newDist
                    q.append([nei, node, newDist])
        
        for idx, ele in enumerate(out):
            if ele == math.inf:
                out[idx] = -1
    
        return out
    


#{ 
 # Driver Code Starts
#Initial Template for Python 3

if __name__ == '__main__':
    tc = int(input())
    while tc > 0:
        n, m = map(int, input().strip().split())
        edges=[]
        for i in range(m):
            li=list(map(int,input().split()))
            edges.append(li)
        src=int(input())
        ob = Solution()
        ans = ob.shortestPath(edges,n,m,src)
        for i in ans:
            print(i,end=" ")
        print()
        tc -= 1
# } Driver Code Ends