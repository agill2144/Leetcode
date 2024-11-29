func myPow(x float64, n int) float64 {
    // x^n = x^n/2 * x^n/2 = (x*x)^n/2
    if n < 0 {
        n *= -1
        x = 1/x
    }
    var dfs func(base float64, exp int) float64
    dfs = func(base float64, exp int) float64 {
        // base
        if exp <= 1 {
            if exp == 1 {return base}
            return 1
        }

        // logic
        ans := dfs(base*base, exp/2)
        if exp % 2 != 0 {
            ans *= base
        }
        return ans
    }
    return dfs(x, n)
}