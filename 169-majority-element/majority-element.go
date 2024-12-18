func majorityElement(nums []int) int {
    n := len(nums)
    freq := map[int]int{}
    for i := 0; i < n; i++ {
        freq[nums[i]]++
        if freq[nums[i]] > n/2 {return nums[i]}
    }
    return -1
}