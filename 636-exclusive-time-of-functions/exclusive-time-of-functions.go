func exclusiveTime(n int, logs []string) []int {
    out := make([]int, n)
    st := [][]int{} // id, start, childTime

    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        logId, _ := strconv.Atoi(log[0])
        logType := log[1]
        logTime, _ := strconv.Atoi(log[2])
        if logType == "start" {
            st = append(st, []int{logId,logTime,0})
        } else {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            id, start, cTime := top[0],top[1],top[2]
            out[id] += (logTime-start+1) - cTime
            if len(st) > 0 {
                st[len(st)-1][2] += logTime-start+1
            }
        }
    }


    return out
}