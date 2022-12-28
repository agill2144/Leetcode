// time: o(n)
// space: o(1)
func maxSubArrayLen(nums []int, k int) int {
    idxMap := map[int]int{0:-1}
    rSum := 0 // x
    out := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        // x is running sum so far
        // y is a prev subarray sum
        // z is a subarray between y and x
        // [1,1,1,1,1,1,1,1,1]
        //  ------y-----x
        //         --z--
        // if we have x and y, to find z we can do;
        // z = x-y
        // here we are given z instead ( i.e k ), and we need to find whether we have seen a y before
        // y = x-z // z is k
        // if we have seen a y before, then it means sum of elements from y idx+1 till x idx == k
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