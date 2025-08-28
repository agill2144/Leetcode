// x^n can be rewritten as : (x^n/2) * (x^n/2)
// we can recursively solve for one of the expressions
// and multiply an extra base term when n is odd
func myPow(x float64, n int) float64 {
    if n < 0 {x = 1/x; n*=-1}
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