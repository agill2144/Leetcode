func calculateTax(brackets [][]int, income int) float64 {
    // when there is no income, there is nothing to tax on
    if income == 0 {return 0.0}
    taxableAmount := min(brackets[0][0], income)
    var taxes float64 =  float64(taxableAmount) * (float64(brackets[0][1])/100.0)
    income -= taxableAmount
    for i := 1; i < len(brackets) && income > 0; i++ {
        taxableAmount = min(income, brackets[i][0]-brackets[i-1][0])
        taxes += float64(taxableAmount) * (float64(brackets[i][1])/100.0)
        income -= taxableAmount
    }
    return taxes
}

