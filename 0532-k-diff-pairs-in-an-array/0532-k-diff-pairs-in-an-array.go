func findPairs(nums []int, k int) int {

    
    // dedupe it first since nums can have dupes and we cannot count 2,1 and 1,2 as 2 different pairs, they are counted as 1
    
    freqMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freqMap[nums[i]]++
    }
    out := 0
    for num, count := range freqMap {
        if k == 0 {
            if count >= 2 {
                out++
            }
        } else {
            complement := num + k
            if _, ok := freqMap[complement]; ok {out++}
        }
    }
    
    return out
}