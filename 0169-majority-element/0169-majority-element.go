func majorityElement(nums []int) int {
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freqMap[nums[i]]++
    }
    num := -1
    count := -1
    for k, v := range freqMap {
        if v > count {
            count = v
            num = k
        }
    }
    return  num
}