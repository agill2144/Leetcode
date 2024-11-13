func calculateTax(brackets [][]int, income int) float64 {
    // when there is no income, there is nothing to tax on
    if income == 0 {return 0.0}
    taxableAmount := min(float64(brackets[0][0]), float64(income))
    var taxes float64 =  taxableAmount * (float64(brackets[0][1])/100.0)
    income -= int(taxableAmount)
    for i := 1; i < len(brackets); i++ {
        taxableAmount = min(float64(income), float64(brackets[i][0])-float64(brackets[i-1][0]))
        taxes += taxableAmount * (float64(brackets[i][1])/100.0)
        income -= int(taxableAmount)
    }
    return taxes
}

func min(x, y float64) float64 {
    if x < y {return x}
    return y
}