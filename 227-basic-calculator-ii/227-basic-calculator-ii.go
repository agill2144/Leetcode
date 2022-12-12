/*
    The hard part in this is enforcing operator precedence.
    If we did not have x or / , then it would not be as difficult.
    
    We must do multiplcation or division first before any addition or subtraction
    
    approach: Stack
    - Maintain a current number 
    - Maintain a lastOp ( last operator ) - start with "+"
    - Loop the input string
    - If current char is a digit, then
    - append the digit to current number ( current * 10 + $currentChar )
    - Otherwise if the current char is not a digit and not a empty space then its likely a new operator
    - When we run into a new operator, we must process lastOp before changing to new operator
        - if lastOp == "+", push +currentNum to stack
        - if lastOp == "-", push -currentNum to stack
        - if lastOp == "*", pop the top, multpliy with currentNum and push the result to stack
        - if lastOp == "/", pop the top, divide top with currentNum and push the result to stack
        - finally, after processing, change the current number to 0 ( reset )
        - and change the lastOp to new operator
    - Why Stack tho???
        - firstly, if we need to enforce multiplication and or division first, we need to know the prev number
        - for example : 2x3 
            - when we get to 3, we need to multiply 2 with 3
            - we can easily get to 3 IF we did not have spaces in the middle by doing idxOf3-2 ( go 2 steps back )
                - sure we can santize if we do have spaces but that takes space to create a sanitized string
                - or we can go back until we get a valid digit
            - however, its not guranteed that the 2 will be a single digit, it could be N digit...
                - so just for loop from where you first got a valid digit and keep going back until you no longer have a valid digit?
                - sure, but this increases time.. 
                - why increase time when we can concatanate as we are traversing
                - we just need an easy way to get the last concatanated number while processing multiplication and division
            - if we concatnate them as we are traversing, we can store the concatnated value in the stack
            - then we need to process a multi/div operator, we can use the top of the stack ( last concatnated number )
            - and easily perform our operation... 
            - therefore the use of stack.
    time: o(n)
    space: o(n)
*/
// func calculate(s string) int {
//     stack := []int{}
//     lastOp := "+"
//     curr := 0
//     for i, char := range s {
//         stringChar := string(char)
//         n, _ := strconv.Atoi(stringChar)

//         if isDigit(stringChar) {
//             curr = curr * 10 + n
//         }

//         // there are 2 places where we process
//         // 1. When the current char is not a digit || space
//         //  - Process the last operator
//         //  - update to new operator
//         // 2. When we have reached the last character
//         //  - '2x2'
//         //  - current = 2 , operator = +
//         //  - current = 0, operator = x
//         //  - current = 2, operator = x
//         //  if we let our for loop exit on the last 2 without processing, we have missed processing the lastOperator (x)
//         //  Therefore we must also process whenever we are on the last character ( i.e i == len(s)-1 )
//         if (stringChar != " ")  || i == len(s)-1 {
//             // we either have a new Op or we have reached the last char in string
//             // process current op before changing to new op
//             if lastOp == "+" {
//                 stack = append(stack, curr)
//             } else if lastOp == "-" {
//                 stack = append(stack, -curr)
//             } else if lastOp == "*" {
//                 // pop top, multiply with curr and push the new result
//                 stack[len(stack)-1] = stack[len(stack)-1] * curr
//             } else if lastOp == "/" {
//                 // pop top, divide with curr and push the new result
//                 stack[len(stack)-1] = stack[len(stack)-1] / curr
//             }
//             // reset current
//             curr = 0
//             // change the lastOp to new op
//             lastOp = stringChar
//         }
//     }
    
   
    
//     result := 0
//     for len(stack) != 0 {
//         result += stack[len(stack)-1]
//         stack = stack[:len(stack)-1]
//     }
//     return result
// }



/*
    approach: maintain the mess(tail) of new calculated value
    - same thing as expression add operator problem
        - When multiplying : calc - tail + tail x current;  tail = tail * currentNum
        - When dividing : calc - tail + (tail / currentNum); tail = tail / currentNum
    - Maintain a current number ( running number until we hit a operator )
    - Maintain a lastOp ( last operator ) - start with "+"
    - Maintain a tail = 0
    - Maintain a calc = 0 ( final result )
    - Loop over the input string
    - if current char is a digit,
    - then append it to current running number ( currentNum *= 10 + currentChar )
    - if current char is not a digit and not a space , then this means we have ran into a new operator
    - When we run into a new operator, we must process the lastOp before updating to new operator
    - In stack, for addition and subtraction, we would just push to stack (+currentNum or -currentNum based on the lastOp )
    - In stack, for multiplication or division, we would pop the top and multiply/divide with currentNum and push the result back into stack
    - Now we wont need a stack, because our calc variable acts as our running stack ( but not a stack, i.e a running final result )
        - if lastOp == "+", calc += currentNum, tail = currentNum
        - if lastOp == "-", calc -= currentNum, tail = -currentNum
        - if lastOp == "*", calc = calc - tail + (tail * currentNum), tail = tail * currentNum
        - if lastOp == "/", calc = calc - tail + (tail / currentNum), tail = tail / currentNum
    
    time: o(n)
    space: o(1)
*/
func calculate(s string) int {
    calc := 0
    current := 0
    tail := 0
    lastOp := "+"
    for i, char := range s {
        stringChar := string(char)
        n, _ := strconv.Atoi(stringChar)
        if isDigit(stringChar) {
            current = current * 10 + n
        }
        if (stringChar == "+" || stringChar == "-" || stringChar == "*" || stringChar == "/") || i == len(s)-1 {
            if lastOp == "+" {
                calc = calc + current
                tail = current
            } else if lastOp == "-" {
                calc = calc - current
                tail = -current
            } else if lastOp == "*" {
                calc = calc - tail + (tail*current)
                tail = tail * current
            } else if lastOp == "/" {
                calc = calc - tail + (tail/current)
                tail = tail / current
            }
            // reset
            current = 0
            lastOp = stringChar
        }
    }
    return calc
}


func isDigit(s string) bool {
    return s >= "0" && s <= "9"
}