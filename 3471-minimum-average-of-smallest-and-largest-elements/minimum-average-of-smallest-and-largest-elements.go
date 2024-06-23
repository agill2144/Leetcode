func minimumAverage(nums []int) float64 {
    sort.Ints(nums)
    left := 0
    right := len(nums)-1
    minAvg := math.MaxFloat64
    for left < right {
        avg := (float64(nums[left]) + float64(nums[right])) / 2.0
        minAvg = min(minAvg, avg)
        // which ptr to move away from ?
        // if we move away from left ( smaller val )
        // the sum will be bigger, therefore avg will be bigger
        // if we move away from right ( smaller val )
        // the sum will be smaller, therefore avg will be smaller
        // so move away from right ?
        // HOWEVER!
        // the question states; "Remove the smallest element, minElement, and the largest element maxElement, from nums"
        // so once we have seen/used left and right vals, we can no longer use them
        // therefore moving away from both!
        left++
        right--
    }
    return minAvg
}

func min(x, y float64) float64{
    if x < y {return x}
    return y
}