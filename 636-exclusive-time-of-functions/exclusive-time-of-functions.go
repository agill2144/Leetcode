func exclusiveTime(n int, logs []string) []int {
    st := [][]int{} // funcID, startTime
    out := make([]int, n)
    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        id, _ := strconv.Atoi(log[0])
        logType := log[1]
        time, _ := strconv.Atoi(log[2])
        if logType == "start" {
            if len(st) > 0 {
                out[st[len(st)-1][0]] += (time - st[len(st)-1][1])
            }
            st = append(st, []int{id,time})
        } else {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top[0]] += (time-top[1]+1)
            if len(st) > 0 {
                st[len(st)-1][1] = time+1
            }
        }
    }
    return out
}