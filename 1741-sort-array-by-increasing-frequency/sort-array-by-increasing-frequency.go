func frequencySort(nums []int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    sort.Slice(nums, func(i, j int)bool{
        iVal, iFreq := nums[i], freq[nums[i]]
        jVal, jFreq := nums[j], freq[nums[j]]
        if iFreq == jFreq {return iVal > jVal}
        return iFreq < jFreq
    })
    return nums
}