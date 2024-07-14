
func sumSubarrayMins(arr []int) int {
    n := len(arr)
    mod := 1000000007
    total := 0
    st := []int{} // idx
    for i := 0; i < n; i++ {
        for len(st) != 0 && arr[i] < arr[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i
            nsl := -1
            if len(st) != 0 {nsl = st[len(st)-1]}
            leftCount := top-nsl
            rightCount := nsr-top
            totalSubArrs := leftCount*rightCount
            total = (total + (totalSubArrs*arr[top])) % mod 
        }
        st = append(st, i)
    }
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsr := n
        nsl := -1
        if len(st) != 0 {nsl = st[len(st)-1]}
        leftCount := top-nsl
        rightCount := nsr-top
        totalSubArrs := leftCount*rightCount
        total = (total + (totalSubArrs*arr[top])) % mod
    }
    return total
}

// func sumSubarrayMins(arr []int) int {
//     n := len(arr)
//     total := 0
//     for i := 0; i < n; i++ {
//         minVal := math.MaxInt64
//         for j := i; j < n; j++ {
//             minVal = min(minVal, arr[j])
//             total += minVal

//         }
//     }
//     return total
// }