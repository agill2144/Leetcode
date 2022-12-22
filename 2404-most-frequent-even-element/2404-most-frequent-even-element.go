func mostFrequentEven(nums []int) int {
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] % 2 == 0 { freqMap[nums[i]]++ }
    }
    if len(freqMap) == 0 {return -1}
    
    count := -1
    num := -1
    for k, v := range freqMap {
        if count == -1 {
            count = v
            num = k
        } else if (v > count) || (v >= count && k < num) {
            count = v
            num = k
        }
    }
    return num
}