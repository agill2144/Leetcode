func arrayChange(nums []int, operations [][]int) []int {
    idxMap := map[int]int{}
    for i := 0; i < len(nums); i++ {
        idxMap[nums[i]] = i
    }
    
    for i := 0; i < len(operations); i++ {
        oldNum := operations[i][0]
        newNum := operations[i][1]
        idx := idxMap[oldNum]
        nums[idx] = newNum
        idxMap[newNum] = idx
        delete(idxMap, oldNum)
    }
    return nums
}