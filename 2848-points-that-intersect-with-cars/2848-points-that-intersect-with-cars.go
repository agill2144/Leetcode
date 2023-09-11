func numberOfPoints(nums [][]int) int {
    if len(nums) == 0{return -1}
    if len(nums) == 1 {return nums[0][1]-nums[0][0]+1}
    
    sort.SliceStable(nums, func(i, j int)bool{
        return nums[i][0] < nums[j][0]
    })
    merged := [][]int{}
    merged = append(merged, nums[0])
    for i := 1; i < len(nums); i++ {
        if nums[i][0] <= merged[len(merged)-1][1] {
            merged[len(merged)-1][0] = min(nums[i][0], merged[len(merged)-1][0])
            merged[len(merged)-1][1] = max(nums[i][1], merged[len(merged)-1][1])
        } else {
            merged = append(merged, nums[i])
        }
    }
    total := 0
    for i := 0; i < len(merged); i++ {
        total += merged[i][1]-merged[i][0]+1
    }
    return total
}

func min(x, y int) int {
    if x < y {return x}
    return y
}

func max(x, y int) int {
    if x > y {return x}
    return y
}