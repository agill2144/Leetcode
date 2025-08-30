/*
    binary exponentiation: repeatedly square the base, and half the exp
    - when exp is odd, save the current base to multiply LATER (because PEMDAS = exp first and then rem base terms)
    - until exp > 0
    tc = o(logn)
    sc = o(1)

    another approach: binary expo using dfs
    tc = o(logn)
    sc = o(logn) - recursive stack
*/
// func myPow(x float64, n int) float64 {
//     if n < 0 {
//         x = 1/x
//         n *= -1
//     }
//     var res float64 = 1.0
//     for n != 0 {
//         if n % 2 == 0 {
//             x *= x
//             n /= 2
//         } else {
//             // save the current modified base ( which is also the answer created so far to multiply later )
//             // when n = 3
//             // x = has become some updated value
//             // we need to save the current base to be multiplied later
//             // we are collecting all the base in res 
//             res *= x
//             // keep x going as is
//             // when n = 1
//             // this condition will run again and thereby saving our updated x and multiplying the saved bases in res
//             n--
//         }
//     }
//     return res
// }


func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1/x
        n *= -1
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