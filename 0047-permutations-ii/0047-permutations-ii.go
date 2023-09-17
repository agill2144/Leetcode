/*
    approach: dedupe using freq map
    - just like combination sum with dupes,
    - we need to group together the dupes
    - we can further use a index based data structre like an array to control which elment we go to next
    - but this is permutation, we dont care about the order, just generate all
    - therefore we dont need the control of where a recursion starts its for loop on
    - as long as that element has freq avail, add to path, recurse
*/
func permuteUnique(nums []int) [][]int {

    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++    
    }
    out := [][]int{}
    var dfs func(path []int)
    dfs = func(path []int)  {
        // base
        if len(path) == len(nums) {
            deepCp := make([]int, len(path))
            copy(deepCp, path)
            out = append(out, deepCp)
            return
        }
        
        // logic
        for k, v := range freq {
            if v == 0 {continue}
            // action
            path = append(path, k)
            freq[k]--
            // recurse
            dfs(path)
            // backtrack
            freq[k]++
            path = path[:len(path)-1]
        }
    }
    dfs(nil)
    return out
}