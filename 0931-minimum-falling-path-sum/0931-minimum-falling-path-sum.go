import (
    "fmt"
    // "math"
)

func minFallingPathSum(matrix [][]int) int {
    memoMatrix := make([][]int, len(matrix))

    memoMatrix[0] = make([]int, len(matrix))
    for idx, num := range matrix[0] {
        memoMatrix[0][idx] = num
    }
    
    for i:=1;i<len(matrix);i++ {
        memoMatrix[i] = make([]int, len(matrix))
        for j:=0;j<len(matrix);j++ {
            leftVal := 9999999999
            middleVal := 9999999999
            rightVal := 9999999999

            middleVal = memoMatrix[i-1][j] + matrix[i][j]
            if j==0 {
                rightVal = memoMatrix[i-1][j+1] + matrix[i][j]
            } else if j==len(matrix)-1 {
                leftVal = memoMatrix[i-1][j-1] + matrix[i][j]
            } else {
                rightVal = memoMatrix[i-1][j+1] + matrix[i][j]
                leftVal = memoMatrix[i-1][j-1] + matrix[i][j]
            }
            
            if leftVal < middleVal {
                if leftVal < rightVal {
                    memoMatrix[i][j] = leftVal
                } else {
                    memoMatrix[i][j] = rightVal
                }
            } else {
                if rightVal < middleVal {
                    memoMatrix[i][j] = rightVal
                } else {
                    memoMatrix[i][j] = middleVal
                }
            }
        }
    }
    // fmt.Println(memoMatrix)
    min := 9999999999
    for _, num := range memoMatrix[len(matrix)-1] {
        if num < min {
            min = num
        }
    }

    return min
}