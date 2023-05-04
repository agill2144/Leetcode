func predictPartyVictory(senate string) string {
    rq := []int{}
    dq := []int{}
    for i := 0; i < len(senate); i++ {
        if senate[i] == 'R' {
            rq = append(rq, i)
        } else {
            dq = append(dq, i)
        }
    }
    for len(rq) > 0 && len(dq) > 0 {
        d := dq[0]; dq = dq[1:]
        r := rq[0]; rq = rq[1:]
        if d < r {
            dq = append(dq, d+len(senate))
        } else {
            rq = append(rq, r+len(senate))
        }
    }
    if len(rq) == 0 {return "Dire"}
    return "Radiant"
    
}

