func maximalNetworkRank(n int, roads [][]int) int {
    adjList := map[int]*set{}
    for i := 0; i < len(roads); i++ {
        u,v := roads[i][0], roads[i][1]
        
        _, ok := adjList[u]
        if !ok {adjList[u] = newSet()}
        _, ok = adjList[v]
        if !ok {adjList[v] = newSet()}

        adjList[u].add(v)
        adjList[v].add(u)
    }
    
    
    maxCount := 0
    
    // why are we forming all nodes here ?
    // Look at example3
    // "Notice that all the cities do not have to be connected"
    // this means we can pick 2 nodes far away from each other , we are looking for
    // which 2 nodes produce the max number of edges , therefore trying all
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            // if j == i {continue}
            u, v := i, j
            uAdjNodes := adjList[u]
            if uAdjNodes == nil {uAdjNodes=newSet()}
            vAdjNodes := adjList[v]
            if vAdjNodes == nil {vAdjNodes=newSet()}
            total := len(uAdjNodes.items) + len(vAdjNodes.items)
            if uAdjNodes.contains(v) || vAdjNodes.contains(u) {
                total--
            }
            if total > maxCount {maxCount = total}        
        }
    }
    return maxCount
}


type set struct{
    items map[int]struct{}
}

func newSet() *set {
    return &set{items:map[int]struct{}{}}
}
func (s *set) add(x int) {
    s.items[x] = struct{}{}
}
func (s *set) contains(x int) bool {
    _, ok := s.items[x]
    return ok
}