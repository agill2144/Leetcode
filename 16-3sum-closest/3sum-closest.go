func threeSumClosest(nums []int, target int) int {
    closestSum := math.MaxInt64
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        left := i+1
        right := len(nums)-1
        for left < right {
            sum := nums[i]+nums[left] + nums[right]
            if sum == target {return sum}
            if abs(target-sum) < abs(target-closestSum) {closestSum = sum}
            if sum > target {
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