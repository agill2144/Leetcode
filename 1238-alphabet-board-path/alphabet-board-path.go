func alphabetBoardPath(target string) string {
    m := 5
    res := new(strings.Builder)
    r, c := 0, 0
    for i := 0; i < len(target); i++ {
        char := target[i]
        if char == 'z' {
            destRow := m
            destCol := 0
            dir, incr := colDir(c, destCol)
            for c != destCol {
                res.WriteByte(dir)
                c += incr
            }
            dir, incr = rowDir(r, destRow)
            for r != destRow {
                res.WriteByte(dir)
                r += incr
            }
            res.WriteByte('!')
        } else {
            destRow := int(char - 'a') / m
            destCol := int(char - 'a') % m
            dir, incr := rowDir(r,destRow)
            for r != destRow {
                res.WriteByte(dir)
                r += incr
            }
            dir, incr = colDir(c, destCol)
            for c != destCol {
                res.WriteByte(dir)
                c += incr
            }
            res.WriteByte('!')
        }
        
    }
    return res.String()
}


func colDir(c int, targetC int) (byte, int) {
    if targetC > c {
        return 'R',1    
    }
    return 'L', -1
}

func rowDir(r int, targetR int) (byte, int) {
    if targetR > r {
        return 'D', 1
    }
    return 'U', -1
}