import (
    "fmt"
    "math"
)

func minPathSum(grid [][]int) int {
    row := len(grid)
    column := len(grid[0])

    for i:=0;i<row;i++ {
        for j:=0;j<column;j++ {
            if i==0 && j==0 {
                continue
            }
            topVal := 300
            leftVal := 300
            
            if i>0 {
                topVal = grid[i-1][j]
            }
            if j>0 {
                leftVal = grid[i][j-1]
            }

            grid[i][j] = int(math.Min(float64(grid[i][j]+topVal), float64(grid[i][j]+leftVal)))
        }
    }
    fmt.Println(grid)

    return grid[len(grid)-1][len(grid[0])-1]
}