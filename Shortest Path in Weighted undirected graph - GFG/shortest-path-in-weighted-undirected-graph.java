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




class Solution {
    public static List<Integer> shortestPath(int n, int m, int[][] edges) {
        List<List<int[]>> adjList = new ArrayList<>(); // { parent : [[node, dist]] }
        for (int i = 0; i <= n; i++) {
            adjList.add(new ArrayList<>());
        }
        
        for (int i = 0; i < edges.length; i++) {
            int u = edges[i][0];
            int v = edges[i][1];
            int w = edges[i][2];
            adjList.get(u).add(new int[]{v, w});
            adjList.get(v).add(new int[]{u, w});
        }
        
        int start = 1;
        int end = n;
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);

        dist[start] = 0;
        PriorityQueue<PQNode> pq = new PriorityQueue<>();
        pq.offer(new PQNode(new ArrayList<>(Arrays.asList(start)), 0));
        
        while (!pq.isEmpty()) {
            PQNode dq = pq.poll();
            List<Integer> path = dq.path;
            int currNode = path.get(path.size() - 1);
            int currDist = dq.dist;
            
            if (currNode == end) {
                return path;
            }
            
            for (int[] nei : adjList.get(currNode)) {
                int adjNode = nei[0];
                int adjNodeDist = nei[1] + currDist;
                if (adjNodeDist < dist[adjNode]) {
                    List<Integer> newPath = new ArrayList<>(path);
                    newPath.add(adjNode);
                    pq.offer(new PQNode(newPath, adjNodeDist));
                    dist[adjNode] = adjNodeDist;
                }
            }
        }
        
        return new ArrayList<>(Arrays.asList(-1));
    }
    
    static class PQNode implements Comparable<PQNode> {
        List<Integer> path;
        int dist;
        
        public PQNode(List<Integer> path, int dist) {
            this.path = path;
            this.dist = dist;
        }
        
        @Override
        public int compareTo(PQNode other) {
            return Integer.compare(this.dist, other.dist);
        }
    }
}