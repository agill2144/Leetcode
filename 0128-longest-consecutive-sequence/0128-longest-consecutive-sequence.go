func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    out := 0
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        prev := num-1
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