func numTilePossibilities(tiles string) int {
    if len(tiles) <= 1 {return len(tiles)}
    freq := map[byte]int{}
    for i := 0; i < len(tiles); i++ {freq[tiles[i]]++}
    count := 0
    var dfs func()
    dfs = func() {
        // base

        // logic
        count++
        for k, v := range freq {
            if v == 0 {continue}
            // action
            freq[k]--
            // recurse
            dfs()
            // backtrack
            freq[k]++
        }
        
    }
    dfs()
    return count-1
}