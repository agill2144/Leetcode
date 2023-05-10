func areSentencesSimilarTwo(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {

    if len(sentence1) != len(sentence2) {return false}
    
    adjList := map[string][]string{}
    for i := 0; i < len(similarPairs); i++ {
        w1 := similarPairs[i][0]
        w2 := similarPairs[i][1]
        adjList[w1] = append(adjList[w1], w2)
        adjList[w2] = append(adjList[w2], w1)
    }
    
    visited := map[string]struct{}{}
    var dfs func(node, prev, target string) bool
    dfs = func(node, prev, target string) bool{
        // base
        if node == target {return true}
        if _, ok := visited[node]; ok {return false}
        
        
        // logic
        visited[node] = struct{}{}
        for _, nei := range adjList[node] {
            if nei == prev {continue}
            if dfs(nei, node, target) {return true}
        }
        return false
    }
    
    for i := 0; i < len(sentence1); i++ {
        src := sentence1[i]
        dest := sentence2[i]
        if src == dest {continue}
        
        visited = map[string]struct{}{}
        
        // both words should be part of our graph
        // if either one of them is not, there is no path from src->dest
        // therefore we can return false and exit early
        _, ok := adjList[src]
        if !ok {return false}
        _, ok = adjList[dest]
        if !ok {return false}
        
        if !dfs(src, "", dest) {return false}
    }
    return true
    
}