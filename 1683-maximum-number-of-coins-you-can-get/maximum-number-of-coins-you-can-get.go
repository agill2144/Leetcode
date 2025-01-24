func maxCoins(piles []int) int {
    sort.Ints(piles)
    total := 0
    n := len(piles)
    myPtr := n-2
    bobs := 0
    for bobs < myPtr {
        total += piles[myPtr]
        myPtr -= 2
        bobs++
    }
    return total
}
