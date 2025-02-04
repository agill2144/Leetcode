func exclusiveTime(n int, logs []string) []int {
    out := make([]int, n)
    st := []int{} // id
    prevStartTime := 0

    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        logId, _ := strconv.Atoi(log[0])
        logType := log[1]
        logTime, _ := strconv.Atoi(log[2])
        if logType == "start" {
            if len(st) > 0 {
                out[st[len(st)-1]] += (logTime - prevStartTime)
            }
            prevStartTime = logTime
            st = append(st, logId)
        } else if logType == "end" {
            topId := st[len(st)-1]
            st = st[:len(st)-1]
            out[topId] += (logTime - prevStartTime+1)
            prevStartTime = logTime + 1
        }
    }


    return out
}