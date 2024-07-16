func maxOperations(nums []int, k int) int {
    freq := map[int]int{}
    pairs := 0
    for i := 0; i < len(nums); i++ {
        diff := k-nums[i]
        val, ok := freq[diff]
        if ok && val > 0{
            pairs++
            freq[diff]--
            freq[nums[i]]--
        }
        freq[nums[i]]++
    }
    return pairs
}