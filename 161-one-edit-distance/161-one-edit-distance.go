func isOneEditDistance(s string, t string) bool {
    if s == t {
        return false
    }
    sPtr := 0
    tPtr := 0
    insert := true
    del := true
    update := true

    for sPtr < len(s) && tPtr < len(t) {
        sVal := s[sPtr]
        tVal := t[tPtr]
        
        if sVal == tVal {
            sPtr++; tPtr++
            continue
        } 
        
        // update
        if len(s) == len(t) {
            if update {update = false; sPtr++; tPtr++} else {return false}
            continue
        }
        
        // insert 
        if len(s) < len(t) {
            if insert {insert = false; tPtr++;} else {return false}
            continue
        }
        
        // delete
        if len(s) > len(t) {
            if del {del = false; sPtr++} else {return false}
        }
        
    }
    for sPtr < len(s) {
        if del {
            del = false
            sPtr++
        } else {
            return false
        }
    }
    for tPtr < len(t) {
        if insert {
            insert = false
            tPtr++
        } else {
            return false
        }
    }
    
    return true
    
}