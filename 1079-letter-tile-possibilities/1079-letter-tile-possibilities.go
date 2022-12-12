func numTilePossibilities(tiles string) int {
    if len(tiles) <= 1 {return len(tiles)}
    
    freqMap := map[byte]int{}
    for i := 0; i < len(tiles); i++ { freqMap[tiles[i]]++ }
    
    count := 0
    var dfs func() 
    dfs = func() {
        // base
       count++
        
        // logic
        for k, v := range freqMap {
            if v == 0 {continue}
            // action
            freqMap[k]--
            // recurse
            dfs()
            // backtrack
            freqMap[k]++
        }
    }
    dfs()
    return count-1
    
}