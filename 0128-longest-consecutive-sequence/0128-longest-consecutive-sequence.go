func longestConsecutive(nums []int) int {
    if nums == nil || len(nums) == 0 {return 0}
    set := map[int]struct{}{}
    for i := 0 ; i < len(nums); i++ {
        set[nums[i]]=struct{}{}
    }
    maxConsec := 0
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        // only proceed with this num if this is a starting of a seq
        // i.e no number before this exists in nums array
        // however if the prev number of this number exists, 
        // skip because we would want to start from that prev num
        _, ok := set[num-1]
        if ok {continue}
        
        count := 1
        for true {
            _, found := set[num+1]
            if !found {break}
            count++
            num++
        }
        if count > maxConsec {maxConsec = count}
    }
    return maxConsec
}