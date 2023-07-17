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
    static class Node implements Comparable<Node> {
        int node;
        int distance;

        public Node(int node, int distance) {
            this.node = node;
            this.distance = distance;
        }

        @Override
        public int compareTo(Node other) {
            return Integer.compare(this.distance, other.distance);
        }
    }

    public static List<Integer> shortestPath(int n, int m, int edges[][]) {
        List<List<int[]>> adjList = new ArrayList<>();
        for (int i = 0; i <= n; i++) {
            adjList.add(new ArrayList<>());
        }

        for (int i = 0; i < m; i++) {
            int u = edges[i][0];
            int v = edges[i][1];
            int w = edges[i][2];
            adjList.get(u).add(new int[]{v, w});
            adjList.get(v).add(new int[]{u, w});
        }

        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        int[] parent = new int[n + 1];
        int src = 1, dest = n;
        dist[src] = 0;
        parent[src] = src;

        boolean found = false;
        PriorityQueue<Node> minHeap = new PriorityQueue<>();
        minHeap.offer(new Node(src, 0));

        while (!minHeap.isEmpty()) {
            Node node = minHeap.poll();
            int currNode = node.node;
            int currDist = node.distance;

            if (currNode == dest) {
                found = true;
                break;
            }

            for (int[] nei : adjList.get(currNode)) {
                int adjNode = nei[0];
                int adjNodeDist = currDist + nei[1];

                if (adjNodeDist < dist[adjNode]) {
                    dist[adjNode] = adjNodeDist;
                    parent[adjNode] = currNode;
                    minHeap.offer(new Node(adjNode, adjNodeDist));
                }
            }
        }

        if (!found) {
            return Arrays.asList(-1);
        }

        List<Integer> out = new ArrayList<>();
        int ptr = dest;
        while (ptr != parent[ptr]) {
            out.add(ptr);
            ptr = parent[ptr];
        }
        out.add(src);

        int left = 0, right = out.size() - 1;
        while (left < right) {
            int temp = out.get(left);
            out.set(left, out.get(right));
            out.set(right, temp);
            left++;
            right--;
        }

        return out;
    }
}
