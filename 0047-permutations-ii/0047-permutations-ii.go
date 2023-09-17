func permuteUnique(nums []int) [][]int {
    // same problem as combination sum with dupes
    // group the dupes together, exhaust them until we no longer can use them, 
    // only then move forward to next set of elements
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }
    deduped := [][]int{}
    for k,v := range freq {
        deduped = append(deduped, []int{k,v})
    }
    out := [][]int{}
    var dfs func(path []int)
    dfs = func(path []int) {
        // base
        if len(path) == len(nums) {
            newPath := make([]int, len(path))
            copy(newPath, path)
            out = append(out, newPath)
            return
        }
        
        // logic
        for i := 0; i < len(deduped); i++ {
            if deduped[i][1] == 0 {continue}
            path = append(path, deduped[i][0])
            deduped[i][1]--
            dfs(path)
            deduped[i][1]++
            path = path[:len(path)-1]
        }
    }
    dfs(nil)
    return out
}