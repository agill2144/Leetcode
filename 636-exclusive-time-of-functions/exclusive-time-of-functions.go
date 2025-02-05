func exclusiveTime(n int, logs []string) []int {
    st := []int{} // funcID, startTime
    out := make([]int, n)
    prevTime := 0
    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        id, _ := strconv.Atoi(log[0])
        logType := log[1]
        time, _ := strconv.Atoi(log[2])
        if logType == "start" {
            if len(st) > 0 {
                out[st[len(st)-1]] += (time - prevTime)
            }
            prevTime = time
            st = append(st, id)
        } else {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top] += (time-prevTime+1)
            prevTime = time+1
        }
    }
    return out
}