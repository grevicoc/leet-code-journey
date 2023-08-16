import (
    "fmt"
    "math"
)

func maximalSquare(matrix [][]byte) int {
    memoMatrix := make([][]int, len(matrix)+1)
    for idx, _ := range memoMatrix {
        memoMatrix[idx] = make([]int, len(matrix[0])+1)
    }

    max := 0
    for i:=0;i<len(matrix);i++ {
        for j:=0;j<len(matrix[0]);j++ {
            if matrix[i][j] == '1' {
                tempMin := int(math.Min(float64(memoMatrix[i][j]), float64(memoMatrix[i+1][j])))
                tempMin = int(math.Min(float64(tempMin), float64(memoMatrix[i][j+1])))

                memoMatrix[i+1][j+1] = tempMin + 1
                max = int(math.Max(float64(max), float64(memoMatrix[i+1][j+1])))
            }
        }
    }

    fmt.Println(memoMatrix)

    return max*max
}