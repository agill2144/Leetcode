func myPow(x float64, n int) float64 {
    // x^n = x^n/2 * x^n/2 = (x*x)^n/2
    if n < 0 {
        n *= -1
        x = 1/x
    }
    var ans float64 = 1.0
    for n != 0 {
        if n % 2 == 0 {
            x *= x
            n /= 2
        } else {
            // save the extra base term and we will multiply later
            // because we need to continue exp calculation first,
            // because PEMDAS
            // x = 2, n = 3
            // x*x = 2*2 = 4 ; n = 2
            // 4^2 = 16
            // which is not correct because 2^3 = 8
            // and this can be fixed by
            // saving current base to be multiplied later
            // save current base (2) - extra term to be multiplied later
            // and x stays as 2, and n becomes 2 (n--)
            // 2^2 = 4, and we multiply the remaining term we have saved (2)
            // 4*2 = 8, this is correct
            // hence we need to save the current base extra term to be multiplied later!
            ans *= x
            n--
        }
    }
    return ans
}
// func myPow(x float64, n int) float64 {
//     // x^n = x^n/2 * x^n/2 = (x*x)^n/2
//     if n < 0 {
//         n *= -1
//         x = 1/x
//     }
//     var dfs func(base float64, exp int) float64
//     dfs = func(base float64, exp int) float64 {
//         // base
//         if exp <= 1 {
//             if exp == 1 {return base}
//             return 1
//         }

//         // logic
//         ans := dfs(base*base, exp/2)
//         if exp % 2 != 0 {
//             ans *= base
//         }
//         return ans
//     }
//     return dfs(x, n)
// }