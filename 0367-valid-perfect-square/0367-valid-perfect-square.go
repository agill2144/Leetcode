/*
    approach: binary search on answers!
    - we know the sqrt is between 1 and x
    - THIS IS A SORTED RANGE
    - which means we can easily binary search on these answers
    
    *BINARY SEARCH ON ANSWERS IS A COMMON PATTER*
    - Identify if there is a range we can use
    - ITS a RANGE, ranges are sorted
    - then binary search on that range
    
    *The best way to identify the range, is to see it in brute force
*/
func isPerfectSquare(num int) bool {
    left := 1
    right := num

    for left <= right {
        mid := left + (right-left)/2
        if mid*mid == num {return true}
        if mid*mid > num {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    // could not find such $a whose square equals to num ( no such $a^2 == num )
    // therefore return false
    return false
}

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
// func isPerfectSquare(num int) bool {
//     for i := 0; i <= num; i++ {
//         if i*i == num {return true}
//         if i*i > num {return false}
//     }
//     return false
// }