func numWaterBottles(numBottles int, numExchange int) int {
    if numExchange > numBottles {
        return numBottles
    }
    total := 0
    fulls := numBottles
    emptys := 0

    for fulls > 0 {
        aboutToDrink := min(numExchange, fulls)
        fulls -= aboutToDrink
        emptys += aboutToDrink
        total += aboutToDrink
        if emptys == numExchange {
            fulls++
            emptys = 0
        }
    }
    return total
}