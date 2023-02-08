/*
    approach: brute force
    - multiple x to itself n times
    time: o(n)
    space: o(1)
*/
// func myPow(x float64, n int) float64 {
//     if n < 0 {
//         x = 1/x
//         n *= -1
//     }
//     result := 1.0
//     for i := 0; i < n; i++ {
//         result *= x
//     }
//     return result
// }


/*
*/
func myPow(x float64, n int) float64 {
    if n < 0 {
        n *= -1
        x = 1/x
    }
    var dfs func(base float64, exp int) float64
    dfs = func(base float64, exp int) float64 {
        // base
        if exp == 0 {return 1}
        
        // logic
        tmp := dfs(base, exp/2)
        result := tmp * tmp
        if exp % 2 != 0 {
            result *= base
        }
        return result
    }
    return dfs(x, n)
}