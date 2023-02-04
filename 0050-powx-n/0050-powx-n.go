/*
    approach: brute force 
    - loop n times and multiply x...
    - if n < 0 , then x = 1/x and n = -n
    - for loop 0 to n(exlcuding n and multiply x each time to a result var)
    - return result
    time: o(n) - we loop n times
    space: o(1)
    - Outcome: TLE
*/

// brute force
// o(n) time and o(1) space
// func myPow(x float64, n int) float64 {
//     if n < 0 {
//         x = 1/x
//         n *= -1
//     }
//     result := 1.00
//     for i := 0; i < n; i++ {
//         result *= x
//     }
//     return result
// }


/*
    approach: logN recursive
    - if n is negative, make it positive and change the x to 1/x, then we will
    - half the exponent in each recursive call
    - once we reach the leaf of our recursion tree ( the base case when exponent will be 0)
        - then return 1.00
    - once we go back from leaf to a parent call
    - we will check if n(exponent)  is an even number
        - if yes, return result * result
        - if not, return result * result * x (base)
            - why? 
            - because even exponent(n) was perfectly divided in half, odd exponent miss multiplying the base by 1 therefore the extra  x $x(base)
    - How is this optimizing the brute force????
        - Example: 2^10
        - Our brute force loop runs 10 times
        - Whereas the binary search approach runs 5 times - therefore optimized
    
*/


func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1/x
        n *= -1
    }
    var dfs func(base float64, exp int) float64
    dfs = func(base float64, exp int) float64{
        // base
        if exp == 0 {
            return 1
        }
        // logic
        result := dfs(base, exp/2)
        if exp % 2 == 0 {
            return result*result
        } else{
            return result*result*base
        }
    }
    return dfs(x, n)
}