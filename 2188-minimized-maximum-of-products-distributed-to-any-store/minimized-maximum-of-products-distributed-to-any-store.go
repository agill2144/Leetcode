// minimize the maximum
// mid = atMax
// if mid works , search left because we want smallest such mid
// when does mid not work ? when stores needed exceed stores given
// meaning our quantity per store (mid) is too small, increase it the load per store ( left = mid+1 )

func minimizedMaximum(n int, quantities []int) int {
    left := 1
    right := slices.Max(quantities)
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < len(quantities); i++ {
            count += int(math.Ceil(float64(quantities[i])/float64(mid)))
        }
        if count > n {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}