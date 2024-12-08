func maximizeSweetness(sweetness []int, k int) int {
    left := 1
    right := 0
    for i := 0; i < len(sweetness); i++ {
        right += sweetness[i]
    }
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        atMin := mid
        count := 0
        rSum := 0
        for i := 0; i < len(sweetness); i++ {
            rSum += sweetness[i]
            if rSum >= atMin {
                count++
                rSum = 0
            }
        }
        // when does atMin or mid not work?
        // when we have < k+1 chunks
        if count < k+1 {
            // decrease per chunk sweetness level
            right = mid-1            
        } else {
            ans = mid
            // we want max such ans
            left = mid+1
        }
    }
    return ans
}