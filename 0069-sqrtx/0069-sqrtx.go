/*
    approach: linear scan
    - square root of a number ( square! , i.e a^2 )
    - so we need to find such $a whose square == x
    - we can start from 1 and go all the upto x itself
    - and x is not a perfect square, we want the left most possible ans
    
    eg1; x = 25; return 5; because 5^2 = 25
    eg2; x = 28; return 5; because 5^2 = 25, 6^2 = 26, which is greater than x, and we want the leftmost possible ans
    
    time = o(x)
    space = o(1)
*/
func mySqrt(x int) int {
    ans := 0
    for i := 1; i <= x; i++ {
        if i*i <= x {
            ans = i
        } else {
            break
        }
    }
    return ans
}