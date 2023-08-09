func uniquePaths(m int, n int) int {
    memoMatrix := [100][100]int{}
    memoMatrix[0][0] = 1
    memoMatrix[0][1] = 1
    memoMatrix[1][0] = 1

    for i:=0;i<m;i++ {
        for j:=0;j<n;j++ {
            if memoMatrix[i][j] != 0 {
                continue
            }

            botVal := 0
            if j > 0 {
                botVal = memoMatrix[i][j-1]
            }
            upVal := 0
            if i > 0 {
                upVal = memoMatrix[i-1][j]
            }

            memoMatrix[i][j] = upVal + botVal
        }
    }

    return memoMatrix[m-1][n-1]
}