/*
    approach: simulate the process ONLY once
    - We will follow all the instructions only once
    - If at the end, our robot is not looking north - that means it will for sure return to origin
    - If at the end, our robot is facing north;
        - if the current position is already 0,0; then return true
        - otherwise return false
    
    - So we need to keep track of 2 things after applying each instruction.
        1. The direction in which robot is facing
        2. The x,y cordinates in case the robot at the end is facing north
    
    - We will take a dirs array [ $x, $y ]
        - This dirs array will be clock wise format or anti-clock wise.
        - Not some random order.
        - Because this dirs array will help us accomplish the 2 things we need to keep track of.
        
    - dirs array clockwise ( U, R, D, L and then repeat ... )
        - like how a clock works.
            U.     R      D        L
        - { {0,1}, {1,0}, {0,-1}, {-1,0} }
    - We will have a idx ptr pointing to 0th dir ( facing up, thats how the robot starts)
    
    - If our instruction is G
        - Then take the x and y values of the direction in which robot is facing (idx ptr pointing to a dir in dirs array )
        - Simply add our robot's cordinates
    - If our instruction is R
        - Because our dirs arrays is clockwise, inorder to move right, it simply means idx++
        - but if idx goes out bounds?
        - Thats why we will use mod operator to return a idx relative to array size
        - R = (idx+1)%4
    - If our instruction is to turn Left
        - Because our dirs array is clockwise
        - We will do (idx+1)%3
    
    Finally check if our robot is not facing north (idx != 0 in dirs array)
    - if it is, return true if current robot cordinates is 0,0
    
    time: o(n)
    space: o(1)
*/
func isRobotBounded(instructions string) bool {
    x := 0
    y := 0
    dirs := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    dirPtr := 0
    
    for i := 0; i < len(instructions); i++ {
        inst := instructions[i]
        if inst == 'G' {
            x += dirs[dirPtr][0]
            y += dirs[dirPtr][1]
        } else if inst == 'R' {
            dirPtr = (dirPtr+1)%4
        } else {
            dirPtr = (dirPtr+3)%4
        }
    }
    return (x == 0 && y == 0) || dirPtr != 0
}


/*
    initial naive code after understanding the explanation
    time: o(n)
    space: o(1)
*/

// func isRobotBounded(instructions string) bool {
//     x := 0
//     y := 0
//     dir := "u"
//     idx := 0
    
//     for idx < len(instructions) {
//         inst := string(instructions[idx])
//         if inst == "G" {
//             if dir == "u" {
//                 y++
//             } else if dir == "d" {
//                 y--
//             } else if dir == "l" {
//                 x--
//             } else if dir == "r" {
//                 x++
//             }
//         } else if inst == "L" {
//             if dir == "u" {
//                 dir = "l"
//             } else if dir == "l" {
//                 dir = "d"
//             } else if dir == "d" {
//                 dir = "r"
//             } else if dir == "r" {
//                 dir = "u"
//             }
//         } else if inst == "R" {
//             if dir == "u" {
//                 dir = "r"
//             } else if dir == "l" {
//                 dir = "u"
//             } else if dir == "d" {
//                 dir = "l"
//             } else if dir == "r" {
//                 dir = "d"
//             }
//         }
//         idx++
//     }
//     return dir != "u" || (x == 0 && y == 0)
// }