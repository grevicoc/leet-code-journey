import "math"

func minCostClimbingStairs(cost []int) int {
    memo := [1001]float64{}
    memo[0] = float64(cost[0])
    memo[1] = float64(cost[1])

    idx := 2
    for idx <= len(cost){
        currCost := 0.0
        if idx < len(cost){
            currCost = float64(cost[idx])
        }

        memo[idx] = math.Min(currCost+memo[idx-1], currCost+memo[idx-2])
        idx++ 
    }

    return int(memo[len(cost)])
}