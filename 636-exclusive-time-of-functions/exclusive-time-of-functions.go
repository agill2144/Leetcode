func exclusiveTime(n int, logs []string) []int {
    st := [][]int{} // id, startTime
    out := make([]int, n)
    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        id, _ := strconv.Atoi(log[0])
        logType := log[1]
        time, _ := strconv.Atoi(log[2])
        if logType == "start" {
            if len(st) > 0 {
                peek := st[len(st)-1]
                out[peek[0]] += (time-peek[1])
            }
            st = append(st, []int{id, time})
        } else if logType == "end" {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top[0]] += (time-top[1]+1)
            // resume parent
            // parent has resumed at a new start time, therefore update 
            // the parent start time to time+1
            // what if there were multiple parents?
            // well, when the top one is popped, processed, it will propagate and update
            // new time to its parent in the stack
            // each popped event is updating its parent of their new resumed start time
            if len(st) > 0 {st[len(st)-1][1] = time+1}
        }
    }
    return out
}