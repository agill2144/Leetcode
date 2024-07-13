func sumSubarrayMins(arr []int) int {
    total := 0
    st := []int{} // idx
    n := len(arr)
    mod := 1000000007
    for i := 0; i < n; i++ {
        for len(st) != 0 && arr[i] < arr[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i
            nsl := -1
            if len(st) != 0 {
                nsl = st[len(st)-1]
            }
            leftCount := top-nsl
            rightCount := nsr-top
            countSubArrs := leftCount*rightCount
            contr := arr[top] * countSubArrs
            total = (total + contr) % mod
        }
        st = append(st, i)
    }
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsr := n
        nsl := -1
        if len(st) != 0 { nsl = st[len(st)-1] }
        leftCount := top-nsl
        rightCount := nsr-top
        countSubArrs := leftCount*rightCount
        contr := arr[top] * countSubArrs
        total = (total + contr) % mod
    }
    return total
}