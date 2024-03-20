func maxLength(ribbons []int, k int) int {
    sort.Ints(ribbons)
    start := 1
    end := ribbons[len(ribbons)-1]
    left := start
    right := end
    ans := 0
    for left <= right {
        mid := left + (right-left)/2

        count := 0
        for j := 0; j < len(ribbons); j++ {
            count += (ribbons[j]/mid)
        }
        // when does mid work ?
        // if we were able to collect atleast k ribbons
        if count >= k {
            ans = mid
            // but keep searching right, since we want max such ans
            left = mid+1
        } else {
            // our desired size is too big, search left and our desired size of a ribbon smaller
            right = mid-1
        }
    }
    return ans
}