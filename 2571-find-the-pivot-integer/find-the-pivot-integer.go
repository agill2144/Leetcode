func pivotInteger(n int) int {
    total := n*(n+1)/2
    leftSum := 0
    x := 1
    for x <= n {
        leftSum += x
        rightSum := total-leftSum+x
        if rightSum == leftSum {return x}
        x++
    }
    return -1
}