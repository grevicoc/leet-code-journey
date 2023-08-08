// func solve(num int, memo *map[int]int) int{
//     if num == 0 {
//         return 0
//     }
//     if num == 1 || num == 2 {
//         return 1
//     }
//     if val, ok := (*memo)[num]; ok {
//         return val
//     }

//     calcVal := solve(num-1, memo) + solve(num-2, memo) + solve(num-3, memo)
//     (*memo)[num] = calcVal
//     return calcVal
// }

func tribonacci(n int) int {
    //top-down
    // memo := map[int]int{}
    // retval := solve(n, &memo)
    // return retval

    //bottom-up
    var memo [38]int
    memo[0] = 0
    memo[1] = 1
    memo[2] = 1

    if (n == 0 || n == 1 || n == 2){
        return memo[n]
    }

    idx := 3
    for idx <= n {
        memo[idx] = memo[idx-3] + memo[idx-2] + memo[idx-1]
        idx++
    }

    return memo[n]
}