func minimumRounds(tasks []int) int {
    freq := map[int]int{}
    for i := 0; i < len(tasks); i++ {
        freq[tasks[i]]++        
    }
    total := 0
    for _, v := range freq {
        if v == 1 {return -1}
        // greedy and always try to complete 3 tasks
        total += (v/3)
        v = v % 3
        if v != 0 {total++}
    }
    return total
}