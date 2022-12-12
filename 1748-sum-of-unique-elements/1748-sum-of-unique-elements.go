func sumOfUnique(nums []int) int {
    freqMap := map[int]int{}
    for _, num := range nums {
        freqMap[num]++
    }
    sum := 0
    for k , v := range freqMap {
        if v == 1 {
            sum += k
        }
    }
    return sum
}