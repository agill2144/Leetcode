func arraySign(nums []int) int {
    countNegatives := 0
    seenZero := false
    for i := 0; i < len(nums); i++ {
        if nums[i] < 0 {
            countNegatives++
        }
        if nums[i] == 0{seenZero = true}
    }
    if seenZero {return 0}
    if countNegatives % 2 == 0 {
        return 1
    }
    return -1
}