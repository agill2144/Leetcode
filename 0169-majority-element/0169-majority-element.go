func majorityElement(nums []int) int {
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freqMap[nums[i]]++
        if val := freqMap[nums[i]]; val > len(nums)/2 {return nums[i]}
    }
    return -1
}