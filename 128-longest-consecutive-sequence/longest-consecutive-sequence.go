func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    longest := 0
    for val, _ := range set {

        _, prevExists := set[val-1]
        if prevExists {continue}

        _, ok := set[val]
        start := val
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