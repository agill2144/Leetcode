func myPow(x float64, n int) float64 {
    if n < 0 {
        n *= -1
        x = 1/x
    }
    if n == 0 {return 1.0}
    var leftOverTerms float64 = 1.0
    for n > 1 {
        if n % 2 != 0 {
            // save base to res to multiply left over term later
            // because PEMDAS - exponent first and then multiplcation
            // we may still have exponent left ( n could be 2 for instance )
            // therefore save current base as is to result
            leftOverTerms *= x
            n--
        } else {
            x *= x
            n /= 2
        }
    }
    return leftOverTerms * x
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
//         if exp % 2 != 0 {res *= base}
//         return res
//     }
//     return dfs(x, n)
// }