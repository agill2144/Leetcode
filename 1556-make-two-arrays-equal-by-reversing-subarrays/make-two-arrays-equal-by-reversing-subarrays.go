func canBeEqual(target []int, arr []int) bool {
    // only possible if arr[i] exists in target x number of times
    tf := map[int]int{}
    for i := 0; i < len(target); i++ {
        tf[target[i]]++
    }
    for i := 0; i < len(arr); i++ {
        c, ok := tf[arr[i]]
        if ok {
            tf[arr[i]]--
            if c == 1 {delete(tf, arr[i])}
        }
    }
    return len(tf) == 0
}