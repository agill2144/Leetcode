func myPow(x float64, n int) float64 {
    if n < 0 {
        n *= -1
        x = 1/x
    }
    var dfs func(base float64, exp int) float64
    dfs = func(base float64, exp int) float64 {
        // base
        if exp <= 1 {
            if exp == 1 {return base}
            return 1.0
        }

        // logic
        res := dfs(base, exp/2)
        res *= res
        if exp % 2 != 0 {res *= base}
        return res        
    }
    return dfs(x, n)
}