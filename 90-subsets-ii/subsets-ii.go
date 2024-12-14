func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base

        // logic
        newL := make([]int, len(path))
        copy(newL, path)
        out = append(out, newL)

        for i := start; i < len(nums); i++ {
            // action : choose this number
            path = append(path, nums[i])
            // recurse
            dfs(i+1, path)
            // backtrack : remove this number
            // as-in, do not choose this num
            skip := path[len(path)-1]
            path = path[:len(path)-1]
            // if next num is same as the one we just removed, skip it!

            for i+1 < len(nums) && nums[i+1] == skip {i++}

        } 
    }
    dfs(0, nil)
    return out
}