func countKDifference(nums []int, k int) int {
    idxMap := map[int][]int{}
    for i := 0; i < len(nums); i++ {
        idxMap[nums[i]] = append(idxMap[nums[i]], i)
    }
    
    count := 0
    for i := 0; i < len(nums); i++ {
        toRemove1 := nums[i] - k
        for _, ele := range idxMap[toRemove1] {
            if ele > i {count++}
        }
        toRemove2 := nums[i] + k
        if toRemove1 == toRemove2 {continue}
        for _, ele := range idxMap[toRemove2] {
            if ele > i {count++}
        }
    }
    return count
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}