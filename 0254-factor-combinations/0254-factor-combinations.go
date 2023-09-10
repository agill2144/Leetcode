// 2 pass, get a list a factors
// then just run backtracking
func getFactors(n int) [][]int {
    if n == 1 {return nil}
    factors := []int{}
    
    // Note: factors should be in the range [2, n - 1].
    // n is not divisible after n/2 
    // eg; n=12, 12/2 = 6; 12 is not divisible by 7 to 12
    // eg; n=40, 40/2 = 20; 40 is not divisible by 21 to 40
    // time = o(n)
    // space = o(n)
    for i := 2; i <= n/2; i++ {
        if n % i == 0 {
            factors = append(factors, i)
        }
    }
    out := [][]int{}
    // time = o(2^n+t)
    // space = o(t) for recursive stack and at a path we allocate o(t) again for deep copy. therefore o(t^2)
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