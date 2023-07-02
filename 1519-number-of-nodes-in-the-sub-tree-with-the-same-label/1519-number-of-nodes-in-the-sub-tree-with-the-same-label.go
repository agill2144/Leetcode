func countSubTrees(n int, edges [][]int, labels string) []int {
    adjList := map[int][]int{}
    for i := 0; i < len(edges); i++ {
        adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
        adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
    }

    out := make([]int, n)    
    var dfs func(node int, prev int) map[byte]int
    dfs = func(node int, prev int) map[byte]int {
        // base
        
        
        // logic
        freqMap := map[byte]int{}
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            
            neiFreqMap := dfs(nei, node)
            for k, v := range neiFreqMap {
                freqMap[k] += v
            }
        }
        freqMap[labels[node]]++
        out[node] = freqMap[labels[node]]
        return freqMap
    }
    dfs(0, -1)
    return out
}