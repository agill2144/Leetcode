// we have to compute how much water can be trapped on top of each building
// we need some sort of gurantee that there exists a maxLeft and maxRight without extra space
// we can use 2 vars to keep track of maxLeft (leftWall) and maxRight (rightWall)
// and use 2 ptrs (left and right) to process either left building or right building
// if we have a leftWall <= rightWall; meaning rightWall is greater than or equal to leftWall
// meaning we have a maxRight, then can probably process left building
// because 1. we have a leftWall(maxLeft) and 2. rightWall(rightMax) and 3. we have left ptr( building to process)
// time = o(n)
// space = o(1)
func trap(height []int) int {
    n := len(height)
    left := 0
    right := n-1
    leftWall := height[left]
    rightWall := height[right]
    total := 0
    for left < right {
        // if leftWall <= rightWall
        // then we know for-sure that there is a rightWall >= leftWall
        // this means, we can for sure figure out how much can be trapped on left building
        // therefore process left building
        // (this is like taking the leftMax from ```min(leftMax,rightMax)-building[i]``` )
        if leftWall <= rightWall {
            // we can calcuate how much we can trap on top of left building
            // we can only trap water on left building IF left leftWall's height >= left building
            // even if leftWall == left building; we can process left building ; i.e we can trap 0 unit of water
            if leftWall >= height[left] {
                total += (leftWall-height[left])
                // we have computed and saved how much we can trap on top of left building while keeping leftWall as-is
                // therefore we can evaluate the next left building
                // and therefore we can move away from left building
                left++
            } else {
                // we have a new leftWall since height[left] > leftWall
                leftWall = height[left]
            }
        } else {
            // we can calcuate how much we can trap on top of right building
            // we can only trap water on right building IF rightWall's height >= right building
            // even if rightWall == right building; we can process right building ; i.e we can trap 0 unit of water
            // another reason of why ==;
            // if rightWall == height[right]; and we promote rightWall to be the value of right building
            // then we never move away from right ptr;
            // rightPtr will only ever move if we were able to trap some water and it can only trap some water if 
            // height[right] > rigthWall
            // BUT does it make sense to stay on right building if rightWall == rightBuilding; no
            // it means, we cannot trap any water on top of right
            // therefore using rightWall >= rightBuilding
            if rightWall >= height[right] {
                // we have computed and saved how much we can trap on top of right building while keeping rightWall as-is
                // therefore we can evaluate the next right building
                // and therefore we can move away from right building
                total += (rightWall-height[right])
                right--
            }  else {
                // we have a new leftWall since height[right] > rightWall
                rightWall = height[right]
            }
        }
    }
    return total
}

/*

    - in brute force, we were looking for maxLeft and maxRight for each building
    - we can pre-compute this in prefix and suffix array
    - and then processing the heights array in linear time

    approach: prefix / suffix
    - for each building;
    - find how much can be trapped above a building
    - we can trap water above a building if;
        - there exists a left building > curr building height
        - there exists a right buidling > curr building height
            __          __
            ||    __    ||
            || __ || __ ||
        - we can trap water above building whose height is 0
            - can trap 2 units of water ; min(leftBuilding, rightBuilding)-currHeight
        - we can trap water above building whose height is 1
            - can trap 1 unit of water ; min(leftBuilding, rightBuilding)-currHeight
        - we can trap water above building whose height is 0
            - can trap 2 units of water ; min(leftBuilding, rightBuilding)-currHeight
        - leftBuilding and rightBuilding are the max heights ( max vals ) excluding current building
        - i.e max val from left and max val from right
        - i.e max prefix from left and max suffix from right
    - once left and right prefixes
    - we can figure out how much can be trapped above each ith building
        - total += min(left[i], right[i]) - height[i]

    time: o(n) + o(n)
    spae = o(n)
*/
// func trap(height []int) int {
//     n := len(height)
//     left := make([]int, n)
//     right := make([]int, n)
//     for i := 0; i < n; i++ {
//         if i == 0 || i == n-1 {continue}
//         maxLeftSoFar := max(left[i-1], height[i-1])
//         maxRightSoFar := max(right[n-1-i+1], height[n-1-i+1])
//         left[i] = maxLeftSoFar
//         right[n-1-i] = maxRightSoFar
//     }
//     total := 0
//     for i := 0; i < n; i++ {
//         leftH, rightH := left[i],right[i]
//         canBeTrappedAboveIthBuilding := min(leftH, rightH)-height[i]
//         if canBeTrappedAboveIthBuilding > 0 {
//             total += canBeTrappedAboveIthBuilding
//         }
//     }
//     return total
// }


// brute force
// time: o(n^2)
// space = o(1)
// func trap(height []int) int {
//     n := len(height)
//     total := 0
//     for i := 1; i < n-1; i++ {
//         curr := height[i]
//         left := 0
//         for k := 0; k < i; k++ {
//             left = max(left, height[k])
//         }
//         right := 0
//         for k := i+1; k < n; k++ {
//             right = max(right, height[k])
//         }
//         wall := min(left, right)
//         if wall > curr {
//             total += wall-curr            
//         }
//     }
//     return total
// }