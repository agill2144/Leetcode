func maxCoins(piles []int) int {
    sort.Ints(piles)
    bob := 0
    alice := 0
    myCoins := 0
    b, i := 0, len(piles)-1
    for b < i {
        alice += piles[i]
        i--
        myCoins += piles[i]
        i--
        bob += piles[b]
        b++
    }
    return myCoins
}