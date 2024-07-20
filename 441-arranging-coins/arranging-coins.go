// we can never have more tha n steps at all
// from 1 to n, this is a sorted range
// can we use binary search to evaluate?
// that is, we are using binary search on our possible answer space
// how can we check if mid step works?
// if n = 5 and mid = 3
// how many coins would we need to fill all 3 steps ? 
// 1+2+3 = 6 ; which is also n(n+1)/2 = sum of all numbers upto n (including n)
// 3(3+1)/2 = 3*4/2 = 12/2 = 6
// and we only have 5 coins, therefore mid is too big, search left
func arrangeCoins(n int) int {
    if n <= 0 {return 0}
    left := 0
    right := n
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        coinsNeeded := (mid * (mid+1))/2
        if coinsNeeded > n {
            right = mid-1
        } else {
            ans = mid
            left = mid+1
        }
    }
    return ans
}
// brute force: simulate it
// func arrangeCoins(n int) int {
//     if n <= 0 {return 0}
//     step := 0
//     for n > 0 {
//         if n >= step+1 {
//             n-=(step+1)
//             step++
//         } else {
//             break
//         }
//     }
//     return step
// }