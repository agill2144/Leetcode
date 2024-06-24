func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    longest := 0
    for i := 0; i < len(nums); i++ {
        _, prevExists := set[nums[i]-1]
        if prevExists {continue}

        _, ok := set[nums[i]]
        start := nums[i]
        size := 0
        for ok {
            size++
            longest = max(longest, size)
            start++
            _, ok = set[start]
        }
    }
    return longest
}