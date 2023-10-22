/*
    approach: linear scan
    - square root of a number ( square! , i.e a^2 )
    - so we need to find such $a whose square == x
    - we can start from 1 and go all the upto x itself
    - if we exceed x, we can return false    
    eg1; x = 25; return 5; because 5^2 = 25
    eg2; x = 28; return false; because 5^2 = 25, 6^2 = 36, which is greater than x,
    
    time = o(x)
    space = o(1)
*/
func isPerfectSquare(num int) bool {
    for i := 0; i <= num; i++ {
        if i*i == num {return true}
        if i*i > num {return false}
    }
    return false
}