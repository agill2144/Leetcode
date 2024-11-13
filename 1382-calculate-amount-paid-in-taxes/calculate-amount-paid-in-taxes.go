func calculateTax(brackets [][]int, income int) float64 {
    // when there is no income, there is nothing to tax on
    if income == 0 {return 0.0}
    taxes := 0.0
    for i := 0; i < len(brackets) && income > 0; i++ {
        /*
            - we have to evaluate each tax bracket ( while our income > 0 )
            - taxRate = difference between current taxRate - prevTaxRate
            - taxRate = 3 and income could be 7
            - then tax is calculated using min value between the two * taxPercent in this bracket
            - taxedAmount = min(taxRate, income) * (bracket[i][1]/100)
            - then this taxedAmount is removed from income
            - then we go to next tax bracket, and repeat the same things
        */
        taxRate := brackets[i][0]
        if i > 0 {taxRate = brackets[i][0] - brackets[i-1][0]}
        taxRate = min(taxRate, income)
        taxes += (float64(taxRate) * (float64(brackets[i][1])/100.0))
        income -= taxRate
    }
    return taxes
}

