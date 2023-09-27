func rearrangeArray(nums []int) []int {
    out := make([]int, len(nums))
    pPtr := 0 // positive ptr
    nPtr := 1 // negative ptr
    for i := 0; i < len(nums); i++{
        // collect all negatives
        if nums[i] < 0 {
            out[nPtr] = nums[i]
            nPtr+=2
        } else {
            out[pPtr] = nums[i]
            pPtr += 2
        }
    }
    return out
}   