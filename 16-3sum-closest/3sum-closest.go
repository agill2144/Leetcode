func threeSumClosest(nums []int, target int) int {
    closestSum := math.MaxInt64
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        left := i+1
        right := len(nums)-1
        for left < right {
            total := nums[i]+nums[left] + nums[right]
            if total == target {return total}
            if abs(target-total) < abs(target-closestSum) {closestSum = total}
            if total > target {
                right--
            } else {
                left++
            }
        }
    }
    return closestSum
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}