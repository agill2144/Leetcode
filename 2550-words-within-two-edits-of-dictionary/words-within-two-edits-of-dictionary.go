func twoEditWords(queries []string, dictionary []string) []string {
    out := []string{}
    for i := 0; i < len(queries); i++ {
        src := queries[i]
        for j := 0; j < len(dictionary); j++ {
            target := dictionary[j]
            if src == target {out = append(out, src); break;}
            
            s,t := 0,0
            count := 0
            for s < len(src) && t < len(target) {
                if src[s] != target[t] {
                    count++
                }
                s++; t++
            }
            if count <= 2 {
                out = append(out, src)
                break
            }
        }
    }
    return out
}
