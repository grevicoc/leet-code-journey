import (
    "fmt"
)

func update(dp *[][]bool, i int, j int, s *string) {
    if i==j {
        (*dp)[i][j] = true
        return
    }

    if j-i==1 {
        if (*s)[i] == (*s)[j] {
            (*dp)[i][j] = true
        } else {
            (*dp)[i][j] = false
        }
        return
    }

    if (*s)[i] == (*s)[j] && (*dp)[i+1][j-1] {
        (*dp)[i][j] = true
    }
}

func longestPalindrome(s string) string {
    memoString := make([][]bool, len(s))
    for idx, _ := range memoString {
        memoString[idx] = make([]bool, len(s))
    }

    retval := ""
    maxLength := 0
    for i:=0;i<len(s);i++ {
        j := 0
        k := i
        for k<len(s) {
            update(&memoString, j, k, &s)
            if memoString[j][k] && k-j+1 > maxLength {
                retval = s[j:k+1]
                maxLength = k-j+1
            }
            j++
            k++
        }
    }
    return retval
}