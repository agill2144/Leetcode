/*
    simulate tax calculation
    - the way tax calculation works is
    - you are given tax brackets in a specific order
    - which states that the first $ amount is tax at some % percent.
    - Then you take the min($income and $taxBracket) * taxPercent and add that to a totalTax var
    - Then you go apply the remaining salary (income - result of min) to the next taxBracket.
    - The next tax bracket is calculated by the abs difference of current tax bracket and prev taxBracket
        - the only exception where this rule does not apply is when we are at the very first tax bracket - where we do not have previous taxBracket
    - But when we have a previous taxBracket and current taxBracket - the new taxBracket becomes abs(currentTaxBracket-previousTaxBracket)
    - then you take this new taxBracket and get a min(taxBracket, $income) - and taxRate will be whatever current taxBracket has ( idx 1 )
    
    - you keep doing this until your salary becomes 0
*/
func calculateTax(brackets [][]int, income int) float64 {
    tax := 0.0
    i := 0
    for income > 0  {
        taxRate := float64(brackets[i][1])/100.00
        taxBracket := brackets[i][0]
        if i != 0 {
            taxBracket = abs(brackets[i][0]-brackets[i-1][0])
        }
        min := int(math.Min(float64(income), float64(taxBracket)))
        tax += (float64(min) * float64(taxRate))
        i++
        income -= min
    }
    
    return tax
}

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}