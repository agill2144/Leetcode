func numberOfBeams(bank []string) int {
    total := 0
    prev := 0
    m := len(bank)
    n := len(bank[0])
    
    for i := 0; i < m; i++ {
        
        curr := 0
        for j := 0; j < n; j++ {
            if bank[i][j] == '1' {curr++}
        }
        
        // row is done
        if prev != 0 {
            total += (curr*prev)   
        }
        // the only time when we do not want to promote curr number of lasers to prev
        // is when curr row had none
        if curr != 0 {prev = curr}
    }
    return total
}