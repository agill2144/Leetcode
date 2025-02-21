func getHint(secret string, guess string) string {
    sFreq := make([]int, 10)
    gFreq := make([]int, 10)
    bulls := 0
    cows := 0
    n := len(secret)
    for i := 0; i < n; i++ {
        sChar := secret[i]
        gChar := guess[i]
        if sChar == gChar {
            bulls++
        } else {
            if sFreq[int(gChar-'0')] > 0 {
                cows++
                sFreq[int(gChar-'0')]--
            } else {
                gFreq[int(gChar-'0')]++
            }

            if gFreq[int(sChar-'0')] > 0 {
                cows++
                gFreq[int(sChar-'0')]--
            } else {
                sFreq[int(sChar-'0')]++
            }
        }
    }
    return fmt.Sprintf("%vA%vB",bulls, cows)
}