func calculate(s string) int {
	n := len(s)
	curr := 0
	var op byte = '+'
	total := 0
	lastContr := 0
	for i := 0; i < n; i++ {
		char := s[i]
		if char >= '0' && char <= '9' {
			curr = curr * 10 + int(char-'0')
		}
		if i == n-1 || char == '+' || char == '-' || char == '/' || char == '*' {
			if op == '+' {
				total += curr
				lastContr = +curr
			} else if op == '-' {
				total -= curr
				lastContr = -curr
			} else if op == '*' {
				multiRes := lastContr * curr
				total = total - lastContr + multiRes
				lastContr = multiRes
			} else if op == '/' {
				divRes := lastContr / curr
				total = total - lastContr + divRes
				lastContr = divRes
			}
			curr = 0
			op = char
		}
	}
	return total
}