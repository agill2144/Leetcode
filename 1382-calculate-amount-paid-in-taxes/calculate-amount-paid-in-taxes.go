/*
    simulate tax calculation
    - the way tax calculation works is
    - you are given tax brackets in a sorted order ( sorted by upper value )
    - which states that income is taxed at each bracket ( as long as income > 0 )
    - so we start a loop on brackets from left to right
    - at each bracket, we need to determine our taxable income based on current bracket, prev bracket(if applicable) and our income
    - first lets find the taxableIncome ( curr - prev )
        - if we have a prev tax bracket (i.e i > 0 )
        - then taxableIncome is current upper - prev upper
        - otherwise, if i == 0 , then our taxableIncome is current bracket upper value
    - now lets compare our taxableIncome with our current income
    - if our income is less than taxableIncome, than taxableIncome is $income
    - if our income is more than taxableIncome, than taxableIncome is $effective tax rate
    - that is, taxableIncome = min(income, taxableIncome)
    - than to calculate taxes owed on taxableIncome
        - taxableIncome * current bracket percent / 100.0
        - we can add this to our total taxes owed so far
    - finally, we need to remove taxableIncome from our income before moving to next bracket
        - income -= taxableIncome
        - its possible that income becomes 0, therefore while looping over brackets, we also need to check income > 0

    time : o(numberOfBrackets)
    space: o(1)


*/
func calculateTax(brackets [][]int, income int) float64 {
    // when there is no income, there is nothing to tax on
    if income == 0 {return 0.0}
    taxes := 0.0
    for i := 0; i < len(brackets) && income > 0; i++ {
        taxableIncome := brackets[i][0]
        if i > 0 {taxableIncome = brackets[i][0] - brackets[i-1][0]}
        taxableIncome = min(taxableIncome, income)
        taxes += (float64(taxableIncome) * (float64(brackets[i][1])/100.0))
        income -= taxableIncome
    }
    return taxes
}

