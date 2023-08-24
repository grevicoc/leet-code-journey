import (
    "fmt"
)

func isExist(s string, wordDict []string) bool {
    for _, word := range wordDict {
        if s==word {
            return true
        }
    }

    return false
}

func solve(memo *map[string]bool, s string, wordDict []string) bool {
    if len(s) == 0 {
        return true
    }

    if exist, ok := (*memo)[s]; ok {
        return exist
    }

    for i:=0;i<len(s);i++ {
        if isExist(s[:i+1], wordDict) {
            if solve(memo, s[i+1:], wordDict) {
                (*memo)[s[:i+1]] = true
                return true
            }
        }
    }

    (*memo)[s] = false
    return false
}

func wordBreak(s string, wordDict []string) bool {
    memoWords := make(map[string]bool)
    
    retval := solve(&memoWords, s, wordDict)
    fmt.Println(memoWords)
    return retval
}