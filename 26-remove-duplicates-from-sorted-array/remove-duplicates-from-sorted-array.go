func removeDuplicates(nums []int) int {
    slow := 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] {
            nums[slow] = nums[i]
            slow++
        }
    }
    return slow
}

/*
              s
    0,1,2,3,4,2,2,3,3,4
                        i
*/