func maxHeightOfTriangle(red int, blue int) int {
    var helper func(r, b int, bFirst bool)int
    helper = func(r, b int, bFirst bool) int {
        flag := true
        if !bFirst {flag = false}
        height := 1
        for r >= height || b >= height {
            if flag && b >= height {
                b -= height
                flag = false
                height++
            } else if !flag && r >= height {
                r -= height
                flag = true
                height++
            } else {
                break
            }
        }
        return height-1
    }
    return max(helper(red, blue, true), helper(red, blue, false))
}