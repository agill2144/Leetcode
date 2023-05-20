/*
Given a Directed Acyclic Graph of N vertices from 0 to N-1 and a 2D Integer array(or vector) edges[ ][ ] of length M, where there is a directed edge from edge[i][0] to edge[i][1] with a distance of edge[i][2] for all i, 0<=i

Find the shortest path from src(0) vertex to all the vertices and if it is impossible to reach any vertex, then return -1 for that vertex.

 

Example:

Input:
n = 6, m= 7
edge=[[0,1,2],[0,4,1],[4,5,4]
,[4,2,2],[1,2,3],[2,3,6],[5,3,1]]

Output:
0 2 3 6 1 5


GFG: https://practice.geeksforgeeks.org/problems/shortest-path-in-undirected-graph/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=direct-acyclic-graph


*/

func main() {
    fmt.Println(shortestPath(6,7,[][]int{{0,1,2},{0,4,1},{4,5,4},{4,2,2},{1,2,3},{2,3,6},{5,3,1}}))
}

/*
    DFS / BFS could lead close to o(v+e)^2 time complexity when there are many edges with different weights
    and we may have to process each node all over again because its parent found a shorter path
    
    The problem with bfs / dfs is we find a shorter dist and we continue exploring its adj nodes ALL over again
    BFS / DFS does not try to minimize short distance from the start
    
    
    Topological sort helps because topSort returns us the strict order of nodes to be processed again
    Such that if there are other edges to other nodes that do lead to shorter distance, their parents will have
    minimized their distances already. Why because, the strict order of top sort, we will ensure each node has minimized dist
    from the start, such that when we get to its child, its parent dist will already by minimized
    
    
    top sort can be done in 2 ways
    1. bfs with indegrees ( indegrees with 0 as src node )
    2. dfs with all nodes as src node 
    
*/




func shortestPath(n, m int, edges [][]int) []int {
    adjList := map[int][][]int{}
    for i := 0; i < len(edges); i++ {
        src := edges[i][0]
        dest := edges[i][1]
        wt := edges[i][2]
        adjList[src] = append(adjList[src], []int{dest, wt})
    }
    
    visited := make([]bool, n)
    var dfs func(node int)
    st := []int{}
    // this is a DAG, no need to check for cycles
    dfs = func(node int) {
        // base
        if visited[node] {return}
        
        // logic
        visited[node] = true
        for _, nei := range adjList[node] {
            dfs(nei[0])
        }
        st = append(st, node)
    }
    
    for i := 0; i < n; i++ {
        if !visited[i] {
            dfs(i)
        }
    }
    src := 0
    for len(st) != 0 && st[len(st)-1] != src {
        st = st[:len(st)-1]
    }
    
    dist := make([]int, n)
    for i := 0; i < n; i++ {dist[i] = math.MaxInt64}
    if len(st) == 0 {return dist}
    
    dist[src] = 0
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        currDist := dist[top]
        for _, neiList := range adjList[top] {
            nei := neiList[0]
            neiDist := neiList[1]
            newDist := currDist + neiDist
            if newDist < dist[nei] {
                dist[nei] = newDist
            }
        }
    }
    return dist

}


// dfs
// func shortestPath(n, m int, edges [][]int) []int {
//     adjList := map[int][][]int{}
//     for i := 0; i < len(edges); i++ {
//         src := edges[i][0]
//         dest := edges[i][1]
//         wt := edges[i][2]
//         adjList[src] = append(adjList[src], []int{dest, wt})
//     }
    
//     visited := make([]bool, n)
//     src := 0
//     dist := make([]int, n)
//     for i := 0; i < n; i++ {
//         dist[i] = math.MaxInt64
//     }
    
//     var dfs func(node int, currWt int)
//     dfs = func(node, currWt int) {
//         // base
//         if visited[node] && currWt >= dist[node] {
//             return
//         }
        
//         // logic
//         dist[node] = currWt
//         visited[node] = true
//         for _, nei := range adjList[node] {
//             dfs(nei[0], currWt+nei[1])
//         }
//     }
//     dfs(src, 0)
//     for i := 0; i < n; i++ {
//         if dist[i] == math.MaxInt64 {dist[i] = -1}
//     }
//     return dist    
// }

// func shortestPath(n, m int, edges [][]int) []int {
//     adjList := map[int][][]int{}
//     for i := 0; i < len(edges); i++ {
//         src := edges[i][0]
//         dest := edges[i][1]
//         wt := edges[i][2]
//         adjList[src] = append(adjList[src], []int{dest, wt})
//     }
//     src := 0
//     out := make([]int, n)
//     for i := 0; i < n; i++ {
//         out[i] = math.MaxInt64
//     }
//     q := [][]int{{src,0}}
//     out[src] = 0
//     for len(q) != 0 {
        
//         dq := q[0]
//         q = q[1:]
//         node := dq[0]
//         dist := dq[1]

//         for _, nei := range adjList[node] {
//             newNode := nei[0]
//             newDist := dist+nei[1]
//             if newDist < out[newNode] {
//                 out[newNode] = newDist
//                 q = append(q, []int{newNode, newDist})
//             }
//         }
//     }
//     for i := 0; i < n; i++ {
//         if out[i] == math.MaxInt64 {out[i] = -1}
//     }
    
//     return out
// }
