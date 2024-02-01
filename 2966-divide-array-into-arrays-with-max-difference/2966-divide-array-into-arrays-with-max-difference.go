func divideArray(nums []int, k int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    for i := 0; i <= len(nums)-3; i+=3 {
        first := nums[i]
        second := nums[i+1]
        third := nums[i+2]
        if third - first > k {
            return nil
        }
        out = append(out, []int{first,second,third})
    }
    return out
}

