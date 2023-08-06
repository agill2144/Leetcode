func findCenter(edges [][]int) int {
    n := len(edges)
    // whichever node has n-1 edges, is the center
    freq := map[int]int{}
    for i := 0; i < len(edges); i++ {
        u,v := edges[i][0], edges[i][1]
        freq[u]++
        freq[v]++
    }
    // fmt.Println(freq)
    for node, countEdges := range freq {
        if countEdges == n {return node}
    }
    return -1
}