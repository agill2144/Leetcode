func findBuildings(heights []int) []int {
    st := []int{}
    for i := len(heights)-1; i >= 0; i-- {
        if len(st) == 0{
            st = append(st, i)
        } else {
            peekBuildIdx := st[len(st)-1]
            peekBuildH := heights[peekBuildIdx]
            if heights[i] > peekBuildH {
                st = append(st, i) 
            }
        }
    }
    for i := 0; i < len(st)/2; i++ {
        st[i], st[len(st)-i-1] = st[len(st)-1-i], st[i]
    }
    return st
}