func findBuildings(heights []int) []int {
    nextGreater := nextGreaterOnRight(heights)
    out := []int{}
    for i := 0; i < len(nextGreater); i++ {
        if nextGreater[i] == -1 {out = append(out, i)}
    }
    return out
}

func nextGreaterOnRight(nums []int) []int {
    // process top
    st := []int{} //idx
    n := len(nums)
    out := make([]int, n)
    for i := 0; i < n; i++ {
        out[i] = -1
    }
    for i := 0; i < n; i++ {
        for len(st) != 0 && nums[i] >= nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top] = i
        }
        st = append(st, i)
    }
    return out
}