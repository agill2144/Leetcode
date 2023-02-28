/*
    approach: dfs
    - because the graphs are disconnected and there are multiple separate connected components, we need to run dfs on all nodes!
    
    - we are simply counting the number of graphs
    - However these are undirected graph
    - meaning we will need to add edges both ways ( the input does not have that )
    - we will maintain a visited set globally to keep track of already explored nodes
        - this will help us avoid same node exploration because of undirected graph
    - One caveat is since we are counting number of components, this means running 1 dfs on a starting node
        will likely not traverse the rest of the disconnected components.
    - There we need to run dfs on ALL nodes
    - which means we need to know the range of nodes
    - and the questions says: You have a graph of n nodes
    - its from 0 to n-1
    - therefore dfs from 0 to n(excluding n)
    - we need a adjList because edges input array is not enough to look up a node neighbors...
    
    time:  o(v+e)
    space: o(v+e) + o(v) for recursive stack
*/
// func countComponents(n int, edges [][]int) int {
//     // create adjList from edges list
//     adjList := map[int][]int{}
//     // time : o(v+e)
//     // space: o(v+e)
//     for i := 0; i < len(edges); i++ {
//         // 2 way connection because this is undirected graph!
//         adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
//         adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])        
//     }
    
//     // visited set to keep track of visited nodes
//     visited := map[int]struct{}{}
    
    
//     // worst case, recursion takes o(v) space because there is only 1 deep graph
//     var dfs func(node int)
//     dfs = func(node int) {
//         // base
//         // if node is already visited, return and exit early
//         if _, ok := visited[node]; ok {return}
        
//         // logic
//         // visit this node
//         visited[node] = struct{}{}
//         // explore its neighbors
//         for _, edge := range adjList[node] {
//             dfs(edge)
//         }
//     }
    
//     // now run dfs on all the nodes
//     count := 0
//     // time: o(v)
//     for i := 0; i < n; i++ {
//         if _, ok := visited[i]; !ok {
//             count++
//             dfs(i)
//         }
//     }
//     return count
// }


// bfs 
// because the graphs are disconnected and there are multiple separate connected components, we need to run bfs on all nodes!
func countComponents(n int, edges [][]int) int {
    // create adjList from edges list
    adjList := map[int][]int{}
    // time : o(v+e)
    // space: o(v+e)
    for i := 0; i < len(edges); i++ {
        // 2 way connection because this is undirected graph!
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
        adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])        
    }
    
    // visited set to keep track of visited nodes
    visited := map[int]struct{}{}
    count := 0
    q := []int{}
    for i := 0; i < n; i++ {
        if _, ok := visited[i]; !ok {
            visited[i] = struct{}{}
            count++

            q = append(q, i)
            for len(q) != 0 {
                dq := q[0]; q = q[1:]
                for _, child := range adjList[dq] {
                    if _, ok := visited[child]; !ok {
                        visited[child] = struct{}{}
                        q = append(q, child)
                    }
                }
            }
            
            
        }
    }
    return count
}