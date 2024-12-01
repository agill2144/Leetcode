func checkIfExist(arr []int) bool {
    set := map[int]bool{}
    for i := 0; i < len(arr); i++ {
        if arr[i] % 2 == 0 && set[arr[i]/2] {return true}
        if set[arr[i]*2] {return true}
        set[arr[i]] = true
    }
    return false
}