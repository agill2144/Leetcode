func canBeEqual(target []int, arr []int) bool {
    tFreq := map[int]int{}
    for i := 0; i < len(target); i++ {
        tFreq[target[i]]++
    }
    for i := 0; i < len(arr); i++ {
        if tFreq[arr[i]] == 0 {return false}
        tFreq[arr[i]]--
        if tFreq[arr[i]] == 0 {delete(tFreq, arr[i])}
    }
    return len(tFreq) == 0 
}