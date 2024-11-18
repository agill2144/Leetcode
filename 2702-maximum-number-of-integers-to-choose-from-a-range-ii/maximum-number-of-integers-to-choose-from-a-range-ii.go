func maxCount(banned []int, n int, maxSum int64) int {
    set := map[int64]bool{}
    for i := 0; i < len(banned); i++ {set[int64(banned[i])] = true}
    left := 1
    right := n
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        sum := int64(mid * (mid+1)/2)
        // why this count ? 
        // we want to know number of numbers we were able to pick
        // we dont want the the mid number itself
        // sum = sum of all number from 1 to mid
        // now we need to remove banned numbers from the sum
        // therefore we are removing the number of digits that made the sum from 1 to mid
        // to keep track of num of numbers that are part of the sum , we use count
        // then we loop over the banned set to remove numbers from the sum that are banned 
        // in doing this, number of digits also go down, therefore count--
        count := mid
        for k,_ := range set {if k <= int64(mid) {sum -= k; count--}}
        if  sum <= maxSum {
            ans = count
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}