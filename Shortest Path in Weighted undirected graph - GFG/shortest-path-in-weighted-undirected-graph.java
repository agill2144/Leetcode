//{ Driver Code Starts
// Initial Template for Java

import java.util.*;
import java.lang.*;
import java.io.*;
@SuppressWarnings("unchecked") class GFG {
    public static void main(String[] args) throws IOException {
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();
        while (T-- > 0) {
            int n = sc.nextInt();
            int m = sc.nextInt();
            int edges[][] = new int[m][3];
            for (int i = 0; i < m; i++) {
                edges[i][0] = sc.nextInt();
                edges[i][1] = sc.nextInt();
                edges[i][2] = sc.nextInt();
            }
            Solution obj = new Solution();
            List<Integer> ans = obj.shortestPath(n, m, edges);
            for (int e : ans) {
                System.out.print(e + " ");
            }
            System.out.println();
        }
    }
}
// } Driver Code Ends


// User function Template for Java

class Pair{
    int first;
    int second;
    public Pair(int first,int second){
        this.first = first;
        this.second = second;
    }
}
class Solution {
    public static List<Integer> shortestPath(int n, int m, int[][] edges) {
        Map<Integer, List<int[]>> adjList = new HashMap<>();
        for (int i = 0; i < edges.length; i++) {
            int src = edges[i][0];
            int dest = edges[i][1];
            int w = edges[i][2];
            adjList.putIfAbsent(src, new ArrayList<>());
            adjList.putIfAbsent(dest, new ArrayList<>());
            adjList.get(src).add(new int[]{dest, w});
            adjList.get(dest).add(new int[]{src, w});
        }
        
        int start = 1;
        int target = n;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[1] - b[1]);
        pq.offer(new int[]{start, 0});
        int[] dist = new int[n + 1];
        int[] parent = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        Arrays.fill(parent, -1);
        dist[start] = 0;
        parent[start] = start;

        while (!pq.isEmpty()) {
            int[] dq = pq.poll();
            int currNode = dq[0];
            int currW = dq[1];
            if (currNode == target) {
                break;
            }
            
            if (adjList.containsKey(currNode)) {
                for (int[] nei : adjList.get(currNode)) {
                    int neiNode = nei[0];
                    int neiWeight = nei[1] + currW;
                    if (neiWeight < dist[neiNode]) {
                        dist[neiNode] = neiWeight;
                        parent[neiNode] = currNode;
                        pq.offer(new int[]{neiNode, neiWeight});
                    }
                }
                
            }
        }
        List<Integer> path = new ArrayList<>();

        if (dist[target] == Integer.MAX_VALUE) {
            path.add(-1);
            return path;
        }
        
        int ptr = target;
        while (ptr != parent[ptr]) {
            path.add(ptr);
            ptr = parent[ptr];
        }
        path.add(start);
        
        Collections.reverse(path);
        return path;
    }
}
