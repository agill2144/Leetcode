// 1 <= s.length <= 16
// i see small constraints, jump to the ultimate brute force; try and explore all paths
// i.e dfs + backtracking
// 2 choices at each char ; take it or not take it
// n = len(s)
// time = o(n * 2^n)
// the extra n is when we create substr at each recursive node
// space = o(n) = max depth of recursion tree
func maxUniqueSplit(s string) int {
    n := len(s)
    set := map[string]bool{}
    maxSplits := 0
    var dfs func(start int)
    // ababccc
    dfs = func(start int) {
        // base
        if start == n {
            maxSplits = max(maxSplits, len(set))
            return
        }

        // logic
        for i := start; i < n; i++ {
            subStr := s[start:i+1]
            if !set[subStr] {
                // action
                set[subStr] = true
                // recurse
                dfs(i+1)
                // backtrack
                delete(set, subStr)
            }
        }
    }
    dfs(0)
    return maxSplits
}