func destCity(paths [][]string) string {
    outdegrees := map[string]int{}
    for i := 0; i < len(paths); i++ {
        src := paths[i][0]
        outdegrees[src]++
    }
    for i := 0; i < len(paths); i++ {
        a, b := paths[i][0], paths[i][1]
        if val, ok := outdegrees[a]; !ok || val == 0 {return a}
        if val, ok := outdegrees[b]; !ok || val == 0 {return b}
    }
    return ""
}