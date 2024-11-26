func subsetsWithDup(nums []int) [][]int {
    m := map[int]int{}
    for i := 0; i < len(nums); i++ {m[nums[i]]++}
    freq := [][]int{}
    for k, v := range m {freq = append(freq, []int{k,v})}
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base

        newL := make([]int, len(path))
        copy(newL, path)
        out = append(out, newL)

        // logic
        for i := start; i < len(freq); i++ {
            if freq[i][1] == 0 {continue}
            // action
            path = append(path, freq[i][0])
            freq[i][1]--
            // recurse
            dfs(i, path)
            // backtrack
            freq[i][1]++
            path = path[:len(path)-1]
        }
    }
    dfs(0, nil)
    return out
}