func find132pattern(nums []int) bool {
    n := len(nums)
    numsK := math.MinInt64
    numsJ := math.MinInt64
    st := []int{} // idx
    for i := n-1; i >= 0; i-- {
        if numsK != math.MinInt64 && nums[i] < numsK && numsK < numsJ {return true}

        for len(st) != 0 && nums[i] > nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            numsJ = nums[i]
            numsK = nums[top]
        }
        st = append(st, i)
    }
    return false
}

// func find132pattern(nums []int) bool {
//     n := len(nums)
//     numsK := math.MinInt64
//     numsJ := math.MinInt64
//     for i := n-1; i >= 0; i-- {
//         if nums[i] > numsJ {
//             numsK = numsJ
//             numsJ = nums[i]
//         } else if numsK != math.MinInt64 && nums[i] < numsK {return true}
//     }
//     return false
// }