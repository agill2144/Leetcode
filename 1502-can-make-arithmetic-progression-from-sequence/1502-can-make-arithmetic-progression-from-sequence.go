func canMakeArithmeticProgression(arr []int) bool {
    sort.Ints(arr)
    diff := arr[1] - arr[0]
    for i := 2; i < len(arr); i++ {
        if arr[i-1] + diff != arr[i] {return false}
    }
    return true
}