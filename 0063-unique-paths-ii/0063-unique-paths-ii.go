import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if obstacleGrid[0][0] == 1 {
        return 0
    } else {
        obstacleGrid[0][0] = 1
    }

    row := len(obstacleGrid)
    column := len(obstacleGrid[0])
    fmt.Println(obstacleGrid)

    for i:=0;i<row;i++ {
        for j:=0;j<column;j++{
            if obstacleGrid[i][j] == 1 {
                if i==0 && j==0 {
                    continue
                }
                obstacleGrid[i][j] = -1
                continue
            }

            topVal := 0
            leftVal := 0
            if i > 0 && obstacleGrid[i-1][j] > 0 {
                topVal = obstacleGrid[i-1][j]
            }
            if j > 0 && obstacleGrid[i][j-1] > 0 {
                leftVal = obstacleGrid[i][j-1]
            }

            obstacleGrid[i][j] = topVal + leftVal
        }
    }

    fmt.Println(obstacleGrid)
    if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] < 0 {
        return 0
    }
    return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}