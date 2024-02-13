func findPairs(nums []int, k int) int {
    ansSet := map[[2]int]struct{}{}
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        // if curr num is num
        num := nums[i]
        // what do we need to remove from num
        // to make it equal to k ?
        // if num = 10; k = 2
        // we need to remove 8 ( i.e 10-2 = 8 )
        // we can also 10 from 12 to make it equal to k
        // i.e searching for 2 complements num+k and num-k
        c1 := num+k
        if _, ok := set[c1]; ok {
            tmp := [2]int{c1, num}
            if num < c1 {tmp = [2]int{num, c1}}
            ansSet[tmp] = struct{}{}
        }

        c2 := num-k
        if _, ok := set[c2]; ok {
            tmp := [2]int{c2, num}
            if num < c2 {tmp = [2]int{num, c2}}
            ansSet[tmp] = struct{}{}
        }

        set[num] = struct{}{}
    }
    return len(ansSet)
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}