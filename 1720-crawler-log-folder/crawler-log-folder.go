func minOperations(logs []string) int {
    st := []string{}
    for i := 0; i < len(logs); i++ {
        if logs[i] == "../" {
            if len(st) != 0 {
                st = st[:len(st)-1]
            }
        } else if logs[i] == "./" {
            continue
        } else {
            st = append(st, logs[i])
        }
    }
    return len(st)
}