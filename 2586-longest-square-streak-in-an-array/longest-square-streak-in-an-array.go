func longestSquareStreak(nums []int) int {
    set := map[float64]bool{}
    for i := 0; i < len(nums); i++ {
        set[float64(nums[i])] = true
    }
    res := 1
    for i := 0; i < len(nums); i++ {
        sqrt := math.Sqrt(float64(nums[i]))
        if set[sqrt] {continue}

        count := 0
        sq := float64(nums[i])
        for set[sq] {
            count++
            sq *= sq
        }
        res = max(res, count)
    }
    if res < 2 {return -1}
    return res
}