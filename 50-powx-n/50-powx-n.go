func myPow(x float64, n int) float64 {
    if n < 0 {
        n *= -1
        x = 1/x
    }
    var dfs func(base float64, exp int) float64
    dfs = func(base float64, exp int) float64 {
        // base
        if exp == 1 {
            return base
        }
        if exp == 0 {return 1}
        
        // logic
        result := dfs(base, exp/2)
        if exp % 2 == 0 {
            return result * result
        }
        return base * result * result
    }
    return dfs(x, n)
}