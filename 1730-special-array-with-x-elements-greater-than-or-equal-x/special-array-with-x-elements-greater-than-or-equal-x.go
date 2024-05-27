func specialArray(nums []int) int {
    sort.Ints(nums)
    for i := 1; i <= nums[len(nums)-1]; i++ {
        x := i

        j := 0
        for j < len(nums)  {
            if nums[j] < x {j++; continue}
            break
        }
        count := len(nums)-j
        if count == x {return x}
    }
    return -1
}