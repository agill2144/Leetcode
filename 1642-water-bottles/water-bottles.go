func numWaterBottles(numBottles int, numExchange int) int {
    if numExchange > numBottles {
        return numBottles
    }
    total := 0
    fulls := numBottles
    emptys := 0

    for fulls > 0 {
        fulls--
        emptys++
        total++
        if emptys == numExchange {
            fulls++
            emptys = 0
        }
    }
    return total
}