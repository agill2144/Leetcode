func exclusiveTime(n int, logs []string) []int {
    prev := 0
    // we are making use of stack
    // only when a func has started
    // on top of an existing / already-started func
    st := []int{}
    out := make([]int, n)
    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        id, _ := strconv.Atoi(log[0])
        logType := log[1]
        time, _ := strconv.Atoi(log[2])
        if logType == "start" {
            if len(st) > 0 {
                out[st[len(st)-1]] += (time-prev)
            }
            prev = time
            st = append(st, id)
        } else if logType == "end" {
            st = st[:len(st)-1]
            out[id] += (time-prev+1)
            prev = time+1
        }
    }
    return out
}