func numTilePossibilities(tiles string) int {
    if len(tiles) <= 1 {return len(tiles)}
    seqs := map[string]bool{}
    var dfs func(set map[int]bool, path string)
    dfs = func(set map[int]bool, path string){
        // base
        if len(path) > 0 {seqs[path]=true}
        // logic
        for i := 0; i < len(tiles); i++ {
            if !set[i] {
                set[i] = true
                dfs(set, path + string(tiles[i]))
                set[i] = false
            }
        }
    }
    dfs(map[int]bool{},"")
    return len(seqs)
}