func myPow(x float64, n int) float64 {
    if n < 0 {
        n *= -1
        x = 1/x
    }
    var dfs func(base float64, exp int) float64
    dfs = func(base float64, exp int) float64 {
        // base
        if exp == 0 {return 1.0}
        
        // logic
        tmp := dfs(base, exp/2)
        result := tmp*tmp
        if exp % 2 != 0 {
            result *= base
        }
        return result
    }
    return dfs(x, n)
}