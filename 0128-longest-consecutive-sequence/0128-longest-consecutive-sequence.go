func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    max := 0
    for i := 0; i < len(nums); i++ {

        curr := nums[i]
        _, ok := set[curr-1]
        if ok {continue} // we have a prev number to start with, skip this one
        
        count := 1
        for true {
            if _, ok := set[curr+1]; ok {
                count++
                curr++
            } else {
                break
            }
        }
        if count > max {max = count}

    }
    return max
}