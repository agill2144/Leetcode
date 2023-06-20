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

class Solution {
    public static List<Integer> shortestPath(int n, int m, int[][] edges) {
        Map<Integer, List<int[]>> adjList = new HashMap<>();
        for (int i = 0; i < edges.length; i++) {
            int u = edges[i][0], v = edges[i][1], w = edges[i][2];
            adjList.putIfAbsent(u, new ArrayList<>());
            adjList.putIfAbsent(v, new ArrayList<>());
            adjList.get(u).add(new int[]{v, w});
            adjList.get(v).add(new int[]{u, w});
        }
        
        int start = 1;
        int end = n;
        
        int[] dist = new int[n + 1];
        int[] parent = new int[n + 1];
        
        Arrays.fill(dist, Integer.MAX_VALUE);
        Arrays.fill(parent, -1);
        
        dist[start] = 0;
        parent[start] = start;
        
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[1] - b[1]);
        pq.offer(new int[]{start, 0});
        
        while (!pq.isEmpty()) {
            int[] dq = pq.poll();
            int node = dq[0];
            int weight = dq[1];
            
            for (int[] nei : adjList.getOrDefault(node, new ArrayList<>())) {
                int adjNode = nei[0];
                int adjNodeWeight = nei[1] + weight;
                
                if (adjNodeWeight < dist[adjNode]) {
                    dist[adjNode] = adjNodeWeight;
                    parent[adjNode] = node;
                    pq.offer(new int[]{adjNode, adjNodeWeight});
                }
            }
        }
        
        if (dist[end] == Integer.MAX_VALUE) {
            return Arrays.asList(-1);
        }
        
        List<Integer> path = new ArrayList<>();
        int ptr = end;
        
        while (ptr != parent[ptr]) {
            path.add(ptr);
            ptr = parent[ptr];
        }
        
        path.add(start);
        reverseList(path);
        
        return path;
    }
    
    private static void reverseList(List<Integer> list) {
        int left = 0;
        int right = list.size() - 1;
        
        while (left < right) {
            int temp = list.get(left);
            list.set(left, list.get(right));
            list.set(right, temp);
            
            left++;
            right--;
        }
    }
}