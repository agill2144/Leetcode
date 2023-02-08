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
approach:  recursion + binary search like
- 3^10 is same as 3^5 * 3^5
- now we solved for 1 of the 3^5, we dont have to solve for the second the 3^5.
- Therefore we can half the exponent each time until the exponent reaches 0
- anything raised to power of 0 is 1 ( base case )
- however when we are solving for 3^5 ; this is not the same as 3^2 * 3^2 ( we halved the 5 therefore the 2 ).. so we are just missing 1 multiplication of base n.
- Therefore when our exponent is odd, it cannot be perfectly halved and we loose multiplying by 1 more base. Therefore multiply the base 1 more time in this case.
time : o(logn)
space: o(logn) - we made half recursion calls. therefore logn space for recursive stack.
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