func check(nums []int) bool {
    n := len(nums)
    idxs := map[int][]int{}
    minNum := 101
    for i := 0; i < n; i++ {
        idxs[nums[i]] = append(idxs[nums[i]], i)
        minNum = min(minNum, nums[i])
    }

    for _, idx := range idxs[minNum] {
        ptr := idx+1
        for ptr % n != idx {
            if nums[ptr%n] >= nums[(ptr-1)%n] {
                ptr++
            } else {
                break
            }
        }
        if ptr % n == idx {return true}
    }
    return false
}