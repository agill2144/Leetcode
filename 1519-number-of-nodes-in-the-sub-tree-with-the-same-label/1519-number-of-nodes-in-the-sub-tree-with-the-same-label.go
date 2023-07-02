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
            // bottom up, take childs values
            // merge with parent
            neiFreqMap := dfs(nei, node)
            //the time of this merge loop is  = o(26)
            // there are only 26 labels ( a-z )
            for k, v := range neiFreqMap {
                freqMap[k] += v
            }
        }
        // update parent
        freqMap[labels[node]]++
        // save parent count
        out[node] = freqMap[labels[node]]
        return freqMap
    }
    dfs(0, -1)
    return out
}