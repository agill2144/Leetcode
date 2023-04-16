/*
find distance from 0 to each node
- build an adj list, with values being a pair
  - <adjNode, dist>

- so we can start a dfs from node 0 and running dist as 0
- then if the running dist for this node is better, save it
- recurse on adjacent nodes with new dist = curr dist + dist to reach that node

- no need for a visited array, because this is a DAG ( directed acyclic graph )
- if we cannot reach a specific node, from 0, then the min distance for that node will be -1 denoting not-reachable

time: o(v+e)
space: o(v+e)
*/
func shortestPathDAGDFS(n, m int, edges [][]int) []int {
	adjList := map[int][][]int{}
	for i := 0; i < len(edges); i++ {
		src := edges[i][0]
		dest := edges[i][1]
		dist := edges[i][2]
		adjList[src] = append(adjList[src], []int{dest, dist})
	}
	src := 0
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = math.MaxInt64
	}

	var dfs func(node, dist int)
	dfs = func(node, dist int) {
		// base
		if dist < out[node] {
			out[node] = dist
		}

		// logic
		for _, nei := range adjList[node] {
			dfs(nei[0], dist+nei[1])
		}
	}
	dfs(src, 0)

	// if there were unreachable nodes
	// update their min dist to -1 as it denotes not-reachable.
	for i := 0; i < n; i++ {
		if out[i] == math.MaxInt64 {
			out[i] = -1
		}
	}
	return out
}
