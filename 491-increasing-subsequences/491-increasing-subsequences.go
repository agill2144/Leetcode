func findSubsequences(nums []int) [][]int {
    var result [][]int
    var backtrack func(start int, paths []int)
    backtrack = func(start int, paths []int) {
        // base
        if len(paths) >= 2 {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
        }
        
        // logic
        set := map[int]struct{}{}
        for i := start; i < len(nums); i++ {
            _, seen := set[nums[i]]
            if (len(paths) == 0 || nums[i] >= paths[len(paths)-1]) && (!seen){
                // action
                set[nums[i]] = struct{}{}
                paths = append(paths, nums[i])
                // recurse
                backtrack(i+1, paths)
                // backtrack
                paths = paths[:len(paths)-1]
            }
        }
    }
    backtrack(0, nil)
    return result
}