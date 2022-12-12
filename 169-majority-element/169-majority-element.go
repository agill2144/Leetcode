
// hashmap
func majorityElement(nums []int) int {
    freqMap := map[int]int{}
    for _, ele := range nums {
        freqMap[ele]++
    }
    majorityCount := 0
    majority := 0
    for k,v := range freqMap {
        if v > majorityCount {
            majorityCount = v
            majority = k
        }
    }
    return majority
}