func isValidSudoku(board [][]byte) bool {
    m := 9
    n := 9
    
    // {"0:1": {5,3},  }
    boxMap := map[string]map[byte]struct{}{}
    
    for i := 0; i < m; i++ {
        rowSet := map[byte]struct{}{}
        colSet := map[byte]struct{}{}
        
        for j := 0; j < n; j++ {
            
            rowVal := board[i][j]
            colVal := board[j][i]
            // fmt.Println(i,j)
            // handle row
            if rowVal != '.' {
                _, existsInRow := rowSet[rowVal]
                if existsInRow { return false }
                rowSet[rowVal] = struct{}{}
            }
            
            // handle col
            if colVal != '.' {
                _, existsInCol := colSet[colVal]
                if existsInCol {  return false }
                colSet[colVal] = struct{}{}               
            }
            
            // handle 3x3 box
            if rowVal != '.'{
                boxKey := fmt.Sprintf("%v:%v",i/3, j/3)
                _, boxExists := boxMap[boxKey]
                if !boxExists {
                    boxMap[boxKey] = map[byte]struct{}{}
                    boxMap[boxKey][rowVal] = struct{}{}
                } else {
                    _, ok := boxMap[boxKey][rowVal]
                    if ok {return false}
                    boxMap[boxKey][rowVal] = struct{}{}
                }
            }
        }
    }
        fmt.Println(boxMap)

    return true
    
}