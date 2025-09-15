func myPow(x float64, n int) float64 {
    if n < 0 {
        n *= -1
        x = 1/x
    }
    var res float64 = 1.0
    for n > 0 {
        if n % 2 == 0 {
            x *= x
            n /= 2
        } else {
            // we have ran into a odd exp
            // save the current state of base we need multiply in res
            // ( recursion did this when we stack unfolded - it kept the base as is )
            // ( here we are mutating the base as we are going )
            // ( hence save the current base state as a snapshot in res var )
            // why ?
            // because PEMDAS
            // solve exp first and then multiply, hence save the current base 
            res *= x
            n--
        }
    }
    return res
}
// func myPow(x float64, n int) float64 {
//     if n < 0 {
//         n *= -1
//         x = 1/x
//     }
//     var dfs func(base float64, exp int) float64
//     dfs = func(base float64, exp int) float64 {
//         // base
//         if exp <= 1 {
//             if exp == 1 {return base}
//             return 1.0
//         }

//         // logic
//         res := dfs(base*base, exp/2)
//         if exp % 2 != 0 {
//             res *= base
//         }
//         return res
//     }
//     return dfs(x, n)
// }