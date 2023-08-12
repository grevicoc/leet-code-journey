import (
    "fmt"
    "math"
)

func minimumTotal(triangle [][]int) int {
    row := len(triangle)

    for i:=1;i<row;i++ {
        column := len(triangle[i])

        for j:=0;j<column;j++ {
            sameVal := 0.0
            beforeVal := 0.0
            if j==0 {
                sameVal = float64(triangle[i-1][j])
                beforeVal = 999999999
            }
            if j == column-1 {
                sameVal = 999999999
                beforeVal = float64(triangle[i-1][j-1])
            }
            if j>0 && j<column-1 {
                sameVal = float64(triangle[i-1][j])
                beforeVal = float64(triangle[i-1][j-1])
            }


            triangle[i][j] = int(math.Min(float64(triangle[i][j])+sameVal, float64(triangle[i][j])+beforeVal))
        }
    }

    min := triangle[len(triangle)-1][0]
    for _, num := range triangle[len(triangle)-1] {
        if num < min {
            min = num
        }
    } 
    fmt.Println(triangle[len(triangle)-1])
    return min
}