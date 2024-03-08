func maxFrequencyElements(nums []int) int {
    freq := map[int]int{}
    maxFreq := -1
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
        maxFreq = max(maxFreq, freq[nums[i]])
    }
    total := 0
    for _, v := range freq {
        if v == maxFreq {
            total += maxFreq
        }
    }
    return total
}