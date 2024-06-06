func simplifyPath(path string) string {
    pathList := strings.Split(path, "/")
    st := []string{}
    for i := 0; i < len(pathList); i++ {
        curr := pathList[i]
        if len(curr) == 0 {continue}
        if curr == ".." {
            if len(st) != 0 {
                st = st[:len(st)-1]
            }
        } else if curr == "." {
            continue
        } else {
            st = append(st, curr)
        }
    }
    tmp := strings.Join(st, "/")
    return "/" + tmp
}