func minimumRounds(tasks []int) int {
    freqMap := map[int]int{}
    for i := 0; i < len(tasks); i++ {
        freqMap[tasks[i]]++
    }
    count := 0
    for _, v := range freqMap{
        if v == 1 {return -1}
        if v % 3 == 0 {count += v / 3; continue}
        count += (v / 3)+1
    }
    return count
}