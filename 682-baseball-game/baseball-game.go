/*
    approach: simulate stack
    time: o(2n) or o(n)
    space: o(n) for the stack
*/
func calPoints(operations []string) int {
    stack := []int{}
    for i := 0; i < len(operations); i++ {
        op := operations[i]
        if op == "D" {
            stack = append(stack, stack[len(stack)-1]*2)
        } else if op == "C" {
            stack = stack[:len(stack)-1]
        } else if op == "+" {
            stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
        } else {
            opInt, _ := strconv.Atoi(op)
            stack = append(stack, opInt)
        }
    }
    total := 0
    for i := 0; i < len(stack); i++ {
        total += stack[i]
    }
    return total
}