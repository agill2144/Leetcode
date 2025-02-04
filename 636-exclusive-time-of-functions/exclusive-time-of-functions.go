func exclusiveTime(n int, logs []string) []int {
    curr := 0
    prev := 0
    st := []int{}
    out := make([]int, n)
    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        id, _ := strconv.Atoi(log[0])
        logType := log[1]
        time, _ := strconv.Atoi(log[2])
        curr = time
        if logType == "start" {
            if len(st) > 0 {
                out[st[len(st)-1]] += (curr-prev)
            }
            prev = curr
            st = append(st, id)
        } else if logType == "end" {
            st = st[:len(st)-1]
            out[id] += (curr-prev+1)
            prev = curr+1
        }
    }
    return out
}