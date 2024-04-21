import heapq
import sys

class Solution:
    def findAnswer(self, n: int, edges: List[List[int]]) -> List[bool]:
        edgeIdx = {}
        adjList = {i: [] for i in range(n)}
        
        for i in range(len(edges)):
            u, v, w = edges[i][0], edges[i][1], edges[i][2]
            adjList[u].append([v, w])
            adjList[v].append([u, w])
            edgeIdx[f"{u}-{v}"] = i
            edgeIdx[f"{v}-{u}"] = i
        
        dist = [self.Node(i, sys.maxsize, "") for i in range(n)]
        dist[0] = self.Node(0, 0, "0")
        dest = n - 1
        shortestDistToDest = sys.maxsize
        out = [False] * len(edges)
        pq = [dist[0]]
        
        while pq:
            dq = heapq.heappop(pq)
            currNode = dq.val
            currDist = dq.dist
            currPath = dq.path
            
            if currNode == dest:
                if currDist < shortestDistToDest:
                    shortestDistToDest = currDist
                    out = [False] * len(edges)
                    splitPath = currPath.split("-")
                    for i in range(len(splitPath)):
                        if i + 1 > len(splitPath) - 1:
                            continue
                        pair = f"{splitPath[i]}-{splitPath[i+1]}"
                        pair2 = f"{splitPath[i+1]}-{splitPath[i]}"
                        if pair in edgeIdx:
                            out[edgeIdx[pair]] = True
                        if pair2 in edgeIdx:
                            out[edgeIdx[pair2]] = True
                elif currDist == shortestDistToDest:
                    splitPath = currPath.split("-")
                    for i in range(len(splitPath)):
                        if i + 1 == len(splitPath):
                            continue
                        pair = f"{splitPath[i]}-{splitPath[i+1]}"
                        pair2 = f"{splitPath[i+1]}-{splitPath[i]}"
                        if pair in edgeIdx:
                            out[edgeIdx[pair]] = True
                        if pair2 in edgeIdx:
                            out[edgeIdx[pair2]] = True
            
            if currDist > dist[currNode].dist:
                continue
            
            for nei in adjList[currNode]:
                neiNode, neiDist = nei[0], currDist + nei[1]
                neiPath = currPath + "-" + str(neiNode)
                neiTypedNode = self.Node(neiNode, neiDist, neiPath)
                if neiDist < dist[neiNode].dist:
                    heapq.heappush(pq, neiTypedNode)
                    dist[neiNode] = neiTypedNode
                elif neiDist == dist[neiNode].dist and neiPath != dist[neiNode].path:
                    heapq.heappush(pq, neiTypedNode)
                    dist[neiNode] = neiTypedNode
        
        return out
    
    class Node:
        def __init__(self, val, dist, path):
            self.val = val
            self.dist = dist
            self.path = path
        
        def __lt__(self, other):
            return self.dist < other.dist