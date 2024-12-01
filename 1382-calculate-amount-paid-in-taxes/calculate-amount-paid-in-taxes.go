func calculateTax(brackets [][]int, income int) float64 {
    var taxes float64
    for i := 0; i < len(brackets); i++ {
        upper, taxRate := brackets[i][0], float64(brackets[i][1])/100.0
        if i-1 >= 0 {upper = brackets[i][0]-brackets[i-1][0]}
        incomeToBeTaxed := min(upper, income)
        taxes += float64(incomeToBeTaxed) * (taxRate)
        income -= incomeToBeTaxed
    }
    return taxes
}