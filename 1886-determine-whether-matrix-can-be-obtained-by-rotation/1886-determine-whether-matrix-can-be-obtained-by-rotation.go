// there are 4 max rotations we can do before origanl state is restored
// (transpose and reverse -> check if equal) is 1 iteraion, do it 3 times because original state is restored on the 4th time
func findRotation(mat [][]int, target [][]int) bool {
    for  i := 0; i < 4; i++ {
        if areEqual(mat, target) {return true}
        rotate(mat)
    }
    return false
}

// transpose and reverse
func rotate(matrix [][]int) {
    n := len(matrix)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }
}


func areEqual(src, target [][]int)bool {
    n := len(src)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if src[i][j] != target[i][j] {return false}
        }
    }
    return true
}