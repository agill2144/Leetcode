/*
    we simulated each row 1 by 1
    - starting with row 1
    - then row 2
    - then row 3.. and so on
    - until we no longer had enough coins left
    - this simulation is happening on a range of numbers!
    - 1,2,3,.... x
    - ranges are sorted
    - we can use binary search on answers
    - left = 1
    - right = n
    - mid = some mid step in a staircase
    - mid is a valid answer if number of coins <= coins needed to fill all steps from 1 to mid row
    - number of coins needed to fill all rows from 1 to mid = is same as sum of all numbers from 1 to mid
    - i.e math formula; n(n+1)/2 = (mid * (mid+1)) / 2
    - take the search to left or right accordingly

    time = o(logn)
    space = o(1)
*/
func arrangeCoins(n int) int {
    left := 1
    right := n
    ans := 0
    for left <= right {
        mid := left+(right-left)/2
        desiredCoins := (mid * (mid+1)) / 2
        if desiredCoins <= n {
            ans = mid
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}
// simulation brute force
// func arrangeCoins(n int) int {
//     k := 1
//     for n >= k {
//         if n < k {break}
//         n -= k
//         k++
//     }
//     return k-1
// }