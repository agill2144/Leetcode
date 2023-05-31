/*
    approach: BFS graph traversal
    - we have to find the shortest transformation sequence from beginWord to endWord
    - this is like saying find the shortest path from startNode to endNode
    - a single transformation sequence is going from word1 -> word2 with only 1 char difference
    - meaning a node can is a adjNode if the number of different chars is exactly 1
    - that means, we can create our adjList by checking each word with all other words
        - i.e check if a node is a child of a given parent node
    - what if endWord does not even exist in wordList
        - while building graph, we will check if we run into the endWord and flip a flag 
            indicating whether we have seen endWord or not. If the flag is not flipped, this
            means there is no endWord in wordList. i.e no path to endWord. return 0
    - what if there are no adj nodes from beginWord
        - after building the graph, we can check whether we have any adj nodes to beginword
        - if not, return 0, there is no path to traverse from beginWord
    - finally, run a classic bfs, use level order traversal to keep track of number of transformations done so far
        - alternatively, we can enqueue a pair <beginWord, lenSoFar>
            - [ [hit,1] ... [cog,5] ]
    
    time: o(n)^2 + o(n)
    - n = len of words in wordList
    - n^2 to build graph
    - n to do bfs
    
    space: o(n+e) + o(n) + o(n)
    - o(n+e) for graph (n = len of wordList, e = total number of edges)
    - o(n) for queue
    - o(n) for visited set to avoid re-processing of nodes that are/were already in queue
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
    hasEnd := false    
    /*
        there could be dupes in wordList
        - we could toss wordList into a set before creating the graph
        the begin word may / may not be part of wordList
        - putting wordList in a set will help eliminate the same duplicate problem
        
        for now, we are going to continue using the wordList as is with beginWord added
    */ 
    adjList := map[string][]string{}
    wordList = append(wordList, beginWord)
    for i := 0; i < len(wordList); i++ {
        parent := wordList[i]
        if parent == endWord {
            hasEnd = true
            continue // no need to create adjNodes for endWord, we are not going past endWord
        }
        for j := 0; j < len(wordList); j++ {
            if i == j {continue}
            child := wordList[j]
            p, c := 0, 0
            diff := 0
            for p < len(parent) && c < len(child) {
                if parent[p] != child[c] {
                    diff++
                } 
                p++
                c++
            }
            if diff == 1 {
                adjList[parent] = append(adjList[parent], child)
            }
        }
    }
    // if endWord is not in wordList, there is no path to reach dest, return 0
    // if beginWord has no adj nodes, there is no path to traverse, return 0
    if !hasEnd || len(adjList[beginWord]) == 0 {return 0}
    
    visited := map[string]struct{}{beginWord: struct{}{}}
    q := []string{beginWord}
    level := 1
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return level}
            for _, nei := range adjList[dq] {
                if _, ok := visited[nei]; !ok {
                    q = append(q, nei)
                    visited[nei] = struct{}{}
                }
            }
            qSize--
        }
        level++
    }
    return 0
}