func sumSubarrayMins(nums []int) int {
    mod := 1000000007
    n := len(nums)
    total := 0
    st := []int{} // idx
    for i := 0; i < n; i++ {
        curr := nums[i]
        for len(st) != 0 && curr < nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i
            nsl := -1
            if len(st) != 0 {
                nsl = st[len(st)-1]
            }
            leftCount := top-nsl
            rightCount := nsr-top
            totalSubarrays := leftCount * rightCount
            total = (total + (totalSubarrays * nums[top])) % mod
        }
        st = append(st, i)
    }

    // if we still have elements in st;
    // it means no ith element became the nsr
    // therefore nsr = n
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsr := n
        nsl := -1
        if len(st) != 0 {
            nsl = st[len(st)-1]
        }
        leftCount := top-nsl
        rightCount := nsr-top
        totalSubarrays := leftCount * rightCount
        total = (total + (totalSubarrays * nums[top])) % mod
    }

    return total
}