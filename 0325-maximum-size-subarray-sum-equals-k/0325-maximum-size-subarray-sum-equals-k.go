
func maxSubArrayLen(nums []int, k int) int {
    idxMap := map[int]int{0:-1}
    rSum := 0 // x
    out := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        // z = x-y
        // here we are given z instead ( i.e k ), we need to find whether we have seen a y before
        // y = x-z // z is k
        y := rSum-k
        idx, ok := idxMap[y]
        if ok {
            if i-idx > out {
                out = i-idx
            }
        }
        if _, ok := idxMap[rSum]; !ok {idxMap[rSum]=i}
    }
    return out
}