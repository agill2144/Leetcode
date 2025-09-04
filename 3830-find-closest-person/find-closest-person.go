func findClosest(x int, y int, z int) int {
    xz := abs(z-x)
    yz := abs(z-y)
    if xz == yz {return 0}
    if xz < yz {return 1}
    return 2
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}