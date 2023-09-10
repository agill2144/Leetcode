// 2 pass, get a list a factors
// then just run backtracking
func getFactors(n int) [][]int {
    if n == 1 {return nil}
    factors := []int{}
    for i := 2; i < n-1; i++ {
        if n % i == 0 {
            factors = append(factors, i)
        }
    }
    out := [][]int{}
    var dfs func(start, prod int, path []int)
    dfs = func(start, prod int, path []int) {
        // base
        if prod >= n {
            if prod == n {
                newL := make([]int, len(path))
                copy(newL, path)
                out = append(out, newL)
            }
            return
        }
        
        // logic
        for i := start; i < len(factors); i++ {
            path = append(path, factors[i])
            dfs(i, prod*factors[i], path)
            path = path[:len(path)-1]
        }
    }
    dfs(0, 1, []int{})
    return out
}