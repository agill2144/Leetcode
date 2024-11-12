func subarraySum(nums []int, k int) int {
    freq := map[int]int{}
    rSum := 0
    total := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        // when sum from idx 0 to i is k
        // we do not have to remove anything from this rSum to make it equal to k
        // well, technically, we have to remove 0 from rSum to make it equal to k
        // but we do not have 0 as rSum in our freq map, therefore handling of this case
        // we have 2 options;
        // 1. add 0 sum to freq at the very beginning {0:1}
        // 2. handle when rSum == k without adding 0 sum in the freq map
        if rSum == k {total++}
        total += freq[rSum - k]
        freq[rSum]++
    }
    return total
}