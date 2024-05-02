func sumSubarrayMins(arr []int) int {
    // nsr and nsl ( processing top elements )
    mod := 1000000007
    n := len(arr)
    totalSum := 0
    st := []int{} // will hold indicies
    for i := 0; i < n; i++ {
        for len(st) != 0 && arr[i] <= arr[st[len(st)-1]] {
            idx := st[len(st)-1]
            st = st[:len(st)-1]
            countOnRight := i-idx
            nsl := -1
            if len(st) != 0 {nsl = st[len(st)-1]}
            countOnLeft := idx-nsl
            totalSubArrs := countOnLeft * countOnRight
            totalSum += ((arr[idx] * totalSubArrs) % mod)
        }
        st = append(st,i)
    }
    for len(st) != 0 {
        idx := st[len(st)-1]
        st = st[:len(st)-1]
        countOnRight := n-idx
        nsl := -1
        if len(st)!= 0 {
            nsl = st[len(st)-1]
        }
        countOnLeft := idx-nsl
        totalSubArrs := countOnLeft * countOnRight
        totalSum = (totalSum+(arr[idx] * totalSubArrs)) % mod
    }
    return totalSum
}