// recursion/backtracking, try every possible split
func maxUniqueSplit(s string) int {
    max := 0
    set := map[string]struct{}{}
    var dfs func(start int)
    dfs = func(start int) {
        // base
        if start == len(s) {
            if len(set) > max {
                max = len(set)
            }
            return
        }
        
        // logic
        for i := start; i < len(s); i++ {
            subStr := string(s[start:i+1])
            if _, ok := set[subStr]; !ok {
                // action
                set[subStr] = struct{}{}
                // recurse
                dfs(i+1)
                // backtrack
                delete(set, subStr)
            }
        }
    }
    dfs(0)
    return max
}