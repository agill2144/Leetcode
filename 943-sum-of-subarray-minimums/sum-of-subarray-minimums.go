func sumSubarrayMins(arr []int) int {
    nsr := nsr(arr)
    nsl := nsl(arr)
    mod := 1000000007
    count := 0
    for i := 0; i < len(arr); i++ {
        rw := nsr[i]-i
        lw := i-nsl[i]
        count = (count + (arr[i]*(lw*rw))) % mod
    }
    return count
}


func nsr(nums []int) []int{
    n := len(nums)
    out := make([]int, n)
    st := []int{}
    for i := n-1; i >= 0; i-- {
        out[i] = n
        curr := nums[i]
        for len(st) != 0 && nums[st[len(st)-1]] >= curr {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            out[i] = st[len(st)-1]
        }
        st = append(st, i)
    }
    return out
}


func nsl(nums []int) []int{
    n := len(nums)
    out := make([]int, n)
    st := []int{}
    for i := 0; i < n; i++ {
        out[i] = -1
        curr := nums[i]
        for len(st) != 0 && nums[st[len(st)-1]] > curr {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            out[i] = st[len(st)-1]
        }
        st = append(st, i)
    }
    return out
}