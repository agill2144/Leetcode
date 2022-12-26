func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    out := 0
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        if _, ok := set[num-1]; ok {
            continue
        }
        
        start := num
        count := 0
        for true {
            start++
            count++
            if _, ok := set[start]; !ok {break}
        }
        out = max(out, count)
    }
    
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}