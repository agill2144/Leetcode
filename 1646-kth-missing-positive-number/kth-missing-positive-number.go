func findKthPositive(arr []int, k int) int {
    val := 1
    missing := 0
    ptr := 0
    for ptr < len(arr) {
        if arr[ptr] != val {
            missing++
            if missing == k {return val}
            val++
        } else {
            ptr++
            val++
        }
    }
    return (val-1) + (k-missing)

}

// k = 7
// [1,2,3,4]
// val = 5
// ptr = 4
// return val + (k-val)