// time: o(n)
// space: o(1)
func leastInterval(tasks []byte, n int) int {
    if n == 0 {return len(tasks)}
    maxCount := 0 // total number times maxFreq was seen
    maxFreq := 0
    m := map[byte]int{}
    for i := 0; i < len(tasks); i++ {
        m[tasks[i]]++
        if m[tasks[i]] > maxFreq {
            maxFreq = m[tasks[i]]
        }
    }
    for _, val := range m {
        if val == maxFreq {maxCount++}
    }
    totalTasks := len(tasks)
    partitions := maxFreq-1
    avail := partitions * (n - (maxCount-1))
    pending := totalTasks - (maxCount*maxFreq)
    idle := int(math.Max(0.0, float64(avail)-float64(pending)))
    return totalTasks+idle
}