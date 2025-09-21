// numOfOptionsPerNode ^ maxDepthOfRecursion
// numOfOptionsPerNode = 2 ( choose, not-choose )
// maxDepthOfRecursion, X = min(k, n)
// tc = 2^X * k
// the extra *k is coming when we are creating a new copy of path to save into output
// sc = o(X) + o(k) when creating paths in recursion
func combinationSum3(k int, n int) [][]int {
    out := [][]int{}
    var dfs func(start int, sum int, path []int)
    dfs = func(start, sum int, path []int) {
        // base
        if len(path) > k {return}
        if sum >= n {
            if sum == n && len(path) == k {
                newP := make([]int, k)
                copy(newP, path)
                out = append(out, newP)
            }
            return
        }

        // logic
        for i := start; i <= 9; i++ {
            sum += i
            path = append(path, i)
            dfs(i+1, sum, path)
            path = path[:len(path)-1]
            sum -= i
        }
    }
    dfs(1,0,nil)
    return out
}