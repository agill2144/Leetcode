func minimizeMax(nums []int, p int) int {
    if p > len(nums)/2 {return 0}
    // such that the maximum difference amongst all the pairs is minimized
    sort.Ints(nums)
    left := 0
    right := nums[len(nums)-1]
    ans := 0
    for left <= right {
        atMax := left + (right-left)/2

        count := 0
        for i := 1; i < len(nums); i++ {
            curr := nums[i]
            prev := nums[i-1]
            diff := curr-prev
            if diff <= atMax {
                count++
                i++
            }
        }
        if count >= p {
            ans = atMax
            right = atMax-1
        } else {
            left = atMax+1
        }
    }
    return ans
}