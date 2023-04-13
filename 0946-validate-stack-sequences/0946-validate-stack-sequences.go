func validateStackSequences(pushed []int, popped []int) bool {
    st := []int{}
    poppedPtr := 0
    for i := 0; i < len(pushed); i++ {
        curr := pushed[i]
        st = append(st, curr)
        for len(st) != 0 && st[len(st)-1] == popped[poppedPtr] {
            st = st[:len(st)-1]
            poppedPtr++
        }
    }
    return poppedPtr == len(popped)
}