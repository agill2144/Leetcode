func simplifyPath(path string) string {
    pathList := strings.Split(path, "/")
    st := []string{}
    for i := 0; i < len(pathList); i++ {
        if pathList[i] == ".." {
            if len(st) != 0 {
                st = st[:len(st)-1]
            }
        } else if pathList[i] == "." || pathList[i] == ""{
            continue
        } else {
            st = append(st, pathList[i])
        }
    }
    out := new(strings.Builder)
    out.WriteString("/")
    for i := 0; i < len(st); i++ {
        out.WriteString(st[i])
        if i != len(st)-1 {out.WriteString("/")}
    }
    return out.String()
}