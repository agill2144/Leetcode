func numberOfBeams(bank []string) int {
    prev := 0
    m := len(bank)
    n := len(bank[0])
    total := 0    
    for i := 0; i < m; i++ {
        curr := 0
        for j := 0; j < n; j++ {
            if bank[i][j] == '1' {
                curr++
            }
        }
        if i != 0 {
            total += (curr*prev)
        }
        if curr == 0 {
            continue
        }
        prev = curr
        
    }
    return total
}