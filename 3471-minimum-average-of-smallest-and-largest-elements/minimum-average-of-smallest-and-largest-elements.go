func minimumAverage(nums []int) float64 {
    sort.Ints(nums)
    left := 0
    right := len(nums)-1
    minAvg := math.MaxFloat64
    for left < right {
        avg := (float64(nums[left]) + float64(nums[right])) / 2.0
        minAvg = min(minAvg, avg)
        left++
        right--
    }
    return minAvg
}

func min(x, y float64) float64{
    if x < y {return x}
    return y
}