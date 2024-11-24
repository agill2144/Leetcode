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
        res := dfs(base*base, exp/2)
        if exp % 2 != 0 {res *= base}
        return res
    }
    return dfs(x, n)
}
// func myPow(x float64, n int) float64 {
//     // assuming the underlying system in a 64-bit system
//     // then converting a -2^31 (min int32) to positivie will not cause an int32 overflow
//     // if the underlying system is a 32 bit system, then handle the overflow using int64
//     // LC system is a 64 bit system, therefore making a min int32 to positive will not overflow
//     if n < 0 {
//         n *= -1
//         x = 1/x
//     }
//     // x^n = (x^2)^n/2
//     // 2^8 = (2x2)^4
//     // 4^4 = (4x4)^4/2 = 16^2
//     // 16^2 = (16x16)^2/2 = 256^1
//     // 256^1 = odd exp, save runningProd to res; res *= 256; n-- ; n=0
//     var runningProd float64 = x
//     var res float64 = 1.0
//     for n != 0 {
//         if n % 2 != 0 {
//             res *= runningProd
//             n--
//         }
//         runningProd *= runningProd
//         n /= 2
//     }
//     return res
// }