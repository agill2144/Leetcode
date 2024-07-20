// brute force: simulate it
func arrangeCoins(n int) int {
    if n <= 0 {return 0}
    step := 0
    for n > 0 {
        if n >= step+1 {
            n-=(step+1)
            step++
        } else {
            break
        }
    }
    return step
}