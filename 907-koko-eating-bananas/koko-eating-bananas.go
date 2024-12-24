func minEatingSpeed(piles []int, h int) int {
    right := 0
    left := 1
    for i := 0; i < len(piles); i++ {
        right += piles[i]
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        speed := mid
        hours := 0
        for j := 0; j < len(piles); j++ {
            hours += int(math.Ceil(float64(piles[j])/float64(speed)))
        }
        // when does mid not work?
        // when hours taken were more than allowed h
        // meaning we need increase the amount eaten per hours ( i.e left = mid+1 )
        if hours > h {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}