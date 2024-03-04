func permuteUnique(nums []int) [][]int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }
    out := [][]int{}
    var dfs func(path []int)
    dfs = func(path []int) {
        // base
        if len(path) == len(nums) {
            newp := make([]int, len(path))
            copy(newp, path)
            out = append(out, newp)
            return
        }
        

        // logic
        for k,v := range freq {
            if v == 0 {continue}
            path = append(path, k)
            freq[k]--
            dfs(path)
            freq[k]++
            path = path[:len(path)-1]
        }
    }
    dfs(nil)
    return out
}