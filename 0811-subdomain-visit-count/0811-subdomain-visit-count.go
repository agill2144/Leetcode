func subdomainVisits(cpdomains []string) []string {
    freqMap := map[string]int{}
    for i := 0; i < len(cpdomains); i++ {
        pair := strings.Split(cpdomains[i], " ")
        count, _ := strconv.Atoi(pair[0])
        fqdn := pair[1]
        freqMap[fqdn] += count
        d2 := strings.SplitN(fqdn, ".", 2)
        freqMap[d2[1]]+=count
        if strings.Contains(d2[1], ".") {
            d3 := strings.Split(d2[1], ".")
            freqMap[d3[1]]+=count            
        }
        
    }
    out := []string{}
    for k, v := range freqMap {
        out = append(out, fmt.Sprintf("%v %v",v,k))
    }
    return out
}