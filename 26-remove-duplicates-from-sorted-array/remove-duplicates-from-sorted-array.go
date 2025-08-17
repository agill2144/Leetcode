func removeDuplicates(nums []int) int {
    s := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        if i == 0 {s++; continue}
        if nums[i] != nums[i-1] {
            nums[s] = nums[i]
            s++
        }
    }
    return s
}

/*
              s
    0,1,2,3,4,2,2,3,3,4
                        i
*/