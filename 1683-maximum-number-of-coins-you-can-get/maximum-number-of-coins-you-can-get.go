func maxCoins(piles []int) int {
    sort.Ints(piles)
    n := len(piles)
    alice := n-1
    i := n-2
    bob := 0
    iTotal := 0
    for i > bob{
        iTotal += piles[i]
        alice-=2
        i-=2
        bob++
    }
    return iTotal
}
/*
    [1,2,2,4,7,8]
     b a i  
*/