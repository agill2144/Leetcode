func sortArrayByParityII(nums []int) []int {
    i := 0
    n := len(nums)
    e := 0
    o := n-1
    for i < n && e < n && o >= 0 {
        numIsEven := nums[i] % 2 == 0
        idxIsEven := i % 2 == 0
        numIsOdd := !numIsEven
        idxIsOdd := !idxIsEven
        if numIsEven {
            if idxIsEven {
                i++
            } else {
                nums[i], nums[e] = nums[e], nums[i]
                e+=2
            }
        } else if numIsOdd {
            if idxIsOdd {
                i++
            } else {
                nums[i], nums[o] = nums[o], nums[i]
                o -= 2
            }
        }
    }
    return nums
}