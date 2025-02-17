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
    // why -1 ?
    // because we counted a empty path when we first start
    // i.e when dfs first starts, our path == ""
    // which is not valid, therefore to remove that extra +1 count, we do a -1 at the end
    return count-1
}