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
        List<int[]>[] adjList = new ArrayList[n + 1];
        for (int i = 0; i <= n; i++) {
            adjList[i] = new ArrayList<>();
        }

        for (int i = 0; i < m; i++) {
            int u = edges[i][0];
            int v = edges[i][1];
            int w = edges[i][2];
            adjList[u].add(new int[]{v, w});
            adjList[v].add(new int[]{u, w});
        }

        int src = 1;
        int dest = n;
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[src] = 0;
        List<Integer> destPath = new ArrayList<>();
        List<QueueNode> q = new ArrayList<>();
        q.add(new QueueNode(new ArrayList<>(Arrays.asList(src)), 0));

        while (!q.isEmpty()) {
            QueueNode dq = q.remove(0);
            List<Integer> currPath = dq.path;
            int currNode = currPath.get(currPath.size() - 1);
            int currDist = dq.distSoFar;

            if (dist[currNode] < currDist) {
                continue;
            }

            if (currNode == dest) {
                if (currDist <= dist[currNode]) {
                    destPath = new ArrayList<>(currPath);
                }
                continue;
            }

            List<int[]> adjNodes = adjList[currNode];
            for (int[] nei : adjNodes) {
                int adjNode = nei[0];
                int adjNodeDist = currDist + nei[1];

                if (adjNodeDist < dist[adjNode]) {
                    dist[adjNode] = adjNodeDist;
                    List<Integer> newPath = new ArrayList<>(currPath);
                    newPath.add(adjNode);
                    q.add(new QueueNode(newPath, adjNodeDist));
                }
            }
        }

        if (destPath.size() == 0) {
            return new ArrayList<>(Arrays.asList(-1));
        }
        return destPath;
    }

    static class QueueNode {
        List<Integer> path;
        int distSoFar;

        public QueueNode(List<Integer> path, int distSoFar) {
            this.path = path;
            this.distSoFar = distSoFar;
        }
    }
}

