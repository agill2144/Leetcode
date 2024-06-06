func validateStackSequences(pushed []int, popped []int) bool {
    if len(pushed) != len(popped) {return false}

    st := []int{}
    p := 0
    for i := 0; i < len(pushed); i++ {
        st = append(st, pushed[i])
        for len(st) != 0 && st[len(st)-1] == popped[p] {
            p++
            st = st[:len(st)-1]
        }
    }
    return len(st) == 0 && p == len(popped)
}