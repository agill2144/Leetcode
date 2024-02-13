func findPairs(nums []int, k int) int {
    set := map[[2]int]struct{}{}
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            diff := abs(nums[j]-nums[i])
            if diff == k {
                tmp := [2]int{nums[i], nums[j]}
                if nums[j] < nums[i] {
                    tmp = [2]int{nums[j], nums[i]}
                }
                if _, ok := set[tmp]; !ok {
                    set[tmp] = struct{}{}
                }
            }
        }
    }
    return len(set)
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}