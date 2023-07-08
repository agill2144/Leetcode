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
            int u = edges[i][0];
            int v = edges[i][1];
            int w = edges[i][2];
            adjList.putIfAbsent(u, new ArrayList<>());
            adjList.putIfAbsent(v, new ArrayList<>());
            adjList.get(u).add(new int[] {v, w});
            adjList.get(v).add(new int[] {u, w});
        }

        int src = 1;
        int dest = n;
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[src] = 0;
        PriorityQueue<QueueNode> pq = new PriorityQueue<>();
        pq.offer(new QueueNode(new ArrayList<>(Arrays.asList(src)), 0));

        while (!pq.isEmpty()) {
            QueueNode dq = pq.poll();
            List<Integer> currPath = dq.path;
            int currNode = currPath.get(currPath.size() - 1);
            int currDist = dq.dist;

            if (currNode == dest) {
                return currPath;
            }

            if (currDist > dist[currNode]) {
                continue;
            }

            List<int[]> adjNodes = adjList.get(currNode);
            if (adjNodes != null) {
                for (int[] nei : adjNodes) {
                    int adjNode = nei[0];
                    int adjNodeDist = currDist + nei[1];
    
                    if (adjNodeDist < dist[adjNode]) {
                        dist[adjNode] = adjNodeDist;
                        List<Integer> newPath = new ArrayList<>(currPath);
                        newPath.add(adjNode);
                        pq.offer(new QueueNode(newPath, adjNodeDist));
                    }
                }
            }
        }

        return Arrays.asList(-1);
    }

    static class QueueNode implements Comparable<QueueNode> {
        List<Integer> path;
        int dist;

        public QueueNode(List<Integer> path, int dist) {
            this.path = path;
            this.dist = dist;
        }

        @Override
        public int compareTo(QueueNode other) {
            return Integer.compare(this.dist, other.dist);
        }
    }
}

