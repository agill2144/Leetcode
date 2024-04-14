func findLatestTime(s string) string {
    strSplit := strings.Split(s,":")
    hh := strSplit[0]
    hhFirst := hh[0]
    hhSecond := hh[1]
    if hhFirst == '?' && hhSecond == '?' {
        hhFirst = '1'
        hhSecond = '1'
    } else if hhFirst == '?' {
        if hhSecond == '0' || hhSecond == '1' {
            hhFirst = '1'            
        } else {
            hhFirst = '0'
        }
    } else if hhSecond == '?' {
        if hhFirst == '0' {
            hhSecond = '9'
        } else {
            hhSecond = '1'
        }
    }
    mm := strSplit[1]
    mmFirst := mm[0]
    mmSecond := mm[1]
    // ?9 -> 5
    // 5? -> 9
    if mmFirst == '?' && mmSecond == '?' {
        mmFirst = '5'
        mmSecond = '9'
    } else if mmFirst == '?' {
        mmFirst = '5'
    } else if mmSecond == '?' {
        mmSecond = '9' 
    }
    
    return fmt.Sprintf("%v%v:%v%v", string(hhFirst),string(hhSecond),string(mmFirst),string(mmSecond))
    
}