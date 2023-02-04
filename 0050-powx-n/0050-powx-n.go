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
        result := dfs(base, exp/2)
        if exp % 2 == 0 {
            return result * result
        } else {
            return result * result * base
        }
    }
    return dfs(x, n)
}