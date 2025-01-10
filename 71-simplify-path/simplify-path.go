func simplifyPath(path string) string {
    dirs := strings.Split(path, "/")
    st := []string{}
    for i := 0; i < len(dirs); i++ {
        if dirs[i] == ".." {
            if len(st) != 0 {st = st[:len(st)-1]}
            continue
        }
        if dirs[i] == "." || dirs[i] == "" {continue}
        st = append(st, dirs[i])
    }
    res := &strings.Builder{}
    res.WriteString("/")
    for i := 0; i < len(st); i++ {
        res.WriteString(st[i])
        if i != len(st)-1 {res.WriteString("/")}
    }
    return res.String()
}