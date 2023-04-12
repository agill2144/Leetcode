func simplifyPath(path string) string {
    if len(path) <= 1 {return path}
    pathSplit := strings.Split(path, "/")
    st := []string{}
    for _, ele := range pathSplit {
        if ele == "" || ele == "." {continue}
        if ele == ".." {
            if len(st) != 0 {st = st[:len(st)-1]}
            continue
        }
        
        st = append(st, ele)
    }
    
    out := new(strings.Builder)
    out.WriteString("/")
    for i := 0; i < len(st); i++ {
        out.WriteString(st[i])
        if i != len(st)-1 {
            out.WriteString("/")
        }
    }
    
    return out.String()
}