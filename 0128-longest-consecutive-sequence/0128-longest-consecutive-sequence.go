func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    out := 0
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        prev := num-1
        // only proceed with this num if this is a starting of a seq
        // i.e no number before this exists in nums array
        // however if the prev number of this number exists, 
        // skip because we would want to start from that prev num
        if _, ok := set[prev]; ok {
            continue
        }
        size := 1
        for true {
            if _, ok := set[num+1]; ok {
                size++
                num++
            } else {
                break
            }
        }
        if size > out {out = size}
    }
    return out
}