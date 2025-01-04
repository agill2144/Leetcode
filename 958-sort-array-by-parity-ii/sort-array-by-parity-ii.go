func sortArrayByParityII(nums []int) []int {
    i := 0
    n := len(nums)
    e := 0
    o := 1
    for i < n && e < n && o >= 0 {
        numIsEven := nums[i] % 2 == 0
        idxIsEven := i % 2 == 0
        numIsOdd := !numIsEven
        idxIsOdd := !idxIsEven
        if numIsEven && idxIsOdd {
            nums[e], nums[i] = nums[i], nums[e]
            e+=2
        } else if numIsOdd &&  idxIsEven {
            nums[o], nums[i] = nums[i], nums[o]
            o+=2
        } else {
            i++
        }
    }
    return nums
}