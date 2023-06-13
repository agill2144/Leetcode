func networkDelayTime(times [][]int, n int, k int) int {
    adjList := map[int][][]int{}
    for i := 0; i < len(times); i++ {
        src := times[i][0]
        dest := times[i][1]
        time := times[i][2]
        adjList[src] = append(adjList[src], []int{dest, time})
    }
    t := make([]int, n+1)
    for i := 1; i < len(t); i++ {
        t[i] = math.MaxInt64
    }
    visited := make([]bool, n+1)
    t[k] = 0
    visited[k] = true
    q := [][]int{{k,0}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        node := dq[0]
        time := dq[1]
        for _, nei := range adjList[node] {
            adjNode := nei[0]
            adjTime := nei[1]
            newTime := time+adjTime
            if !visited[adjNode] || newTime < t[adjNode] {
                visited[adjNode] = true
                t[adjNode] = newTime
                q = append(q, []int{adjNode, newTime})
            } 
        }
    }
    out := -1
    for i := 1; i < len(t); i++ {
        if t[i] == math.MaxInt64 {return -1}
        if t[i] > out {
            out = t[i]
        }
    }
    return out
}

// This is like find the shortest path to all nodes from a given node
// Top sort , and then finding edge distances ( "relaxing edges" )
// The above for this does not work because the above ONLY WORKS FOR DAG! ( i.e no cycles )
// this question has cycles...
// func networkDelayTime(times [][]int, n int, k int) int {
//     adjList := map[int][][]int{}
//     for i := 0; i < len(times); i++ {
//         src := times[i][0]
//         dest := times[i][1]
//         time := times[i][2]
//         adjList[src] = append(adjList[src], []int{dest, time})
//     }
    
//     visited := make([]bool, n+1)
//     st := []int{}
//     var dfs func(node int)
//     dfs = func(node int) {
//         // base
//         if visited[node] {return}
        
//         // logic
//         visited[node] = true
//         for _, nei := range adjList[node] {
//             dfs(nei[0])
//         }
        
//         st = append(st, node)
//     }
//     for i := 1; i <= n; i++ {
//         if !visited[i] {dfs(i)}
//     }
    
//     for st[len(st)-1] != k {
//       st = st[:len(st)-1]
//     }
    
//     dist := make([]int, n+1)
//     for i := 0; i < len(dist); i++ {dist[i] = math.MaxInt64}
//     dist[k] = 0
    
//     for len(st) != 0 {
//         top := st[len(st)-1]
//         st = st[:len(st)-1]
//         topDist := dist[top]
//         if topDist == math.MaxInt64 {topDist = 0}
//         adjNodes := adjList[top]
//         for _, adjNode := range adjNodes {
//             if topDist + adjNode[1] < dist[adjNode[0]] {
//                 dist[adjNode[0]] = topDist + adjNode[1]
//             }
//         }
//     }
//     fmt.Println(dist)
//     out := 0
//     for i := 1; i <= n; i++ {
//         if i == k {continue}
//         if dist[i] > out {out = dist[i]}
//     }    
//     if out == math.MaxInt64 {return -1}
//     return out
// }