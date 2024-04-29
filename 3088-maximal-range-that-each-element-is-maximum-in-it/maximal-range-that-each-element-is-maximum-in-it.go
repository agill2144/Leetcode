// ngr and ngl ; process top implementation
func maximumLengthOfRanges(nums []int) []int {
    st := []int{} // holds indicies to be processed
    n := len(nums)
    out := make([]int, n)
    for i := 0; i < n; i++ {
        curr := nums[i]
        for len(st) != 0 && curr > nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            ngr := i
            ngl := -1
            if len(st) != 0 {
                ngl = st[len(st)-1]
            }
            ngr-=1
            ngl+=1
            size := ngr-ngl+1
            out[top] = size
        }
        // push ith element to be processed
        st = append(st, i)
    }
    
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        ngr := n-1
        ngl := -1
        if len(st) != 0 {
            ngl = st[len(st)-1]
        }
        size := ngr-ngl
        out[top] = size
    }
    return out
}