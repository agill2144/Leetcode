func findKthPositive(arr []int, k int) int {
    missing := []int{}
    counter := 1
    i := 0
    for i < len(arr) {
        if arr[i] == counter {
            i++
            counter++
        } else {
            missing = append(missing, counter)
            counter++
        }
    }
    if k-1 <= len(missing)-1 {return missing[k-1]}
    return arr[len(arr)-1]+k-len(missing)
}