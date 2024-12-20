func buildArray(target []int, n int) []string {
    res := []string{}
    t := 0
    top := -1
    for i := 1; i <= n && t < len(target); i++ {
        res = append(res, "Push")
        top = i
        if top < target[t] {
            res = append(res, "Pop")
        } else if top == target[t] {
            t++
        } 
    }
    return res
}