import (
    "fmt"
    "math"
)

func deleteAndEarn(nums []int) int {
    // var mapNums = map[int]int{}
    // for _, num := range nums {
    //     if _, ok := mapNums[num]; ok {
    //         mapNums[num] = mapNums[num] + num
    //         continue
    //     }
    //     mapNums[num] = num
    // }

    // retval := 0
    // for len(mapNums) > 0 {
    //     maxKey := 0
    //     maxDiff := -999

    //     for key, _ := range mapNums {
    //         diff := 0
    //         if _, ok := mapNums[key+1]; ok {
    //             diff = mapNums[key+1]
    //         }
    //         if _, ok := mapNums[key-1]; ok {
    //             diff += mapNums[key-1]
    //         }

    //         if mapNums[key]-diff > maxDiff {
    //             maxKey = key
    //             maxDiff = mapNums[key]-diff
    //         }
    //     }

    //     retval += mapNums[maxKey]
    //     delete(mapNums, maxKey)
    //     delete(mapNums, maxKey+1)
    //     delete(mapNums, maxKey-1)
    // }

    // return retval

    memo := [10001]float64{}
    for _, num := range nums {
        memo[num] += float64(num)
    }

    max := 0.0
    prev2 := 0.0
    prev1 := 0.0
    for idx, _ := range memo {
        max = math.Max(prev2 + memo[idx], prev1)
        prev2 = prev1
        prev1 = max
    }

    return int(max)


    // var mapNums = map[int]int{}

    // currIdx := 1
    // mapNums[nums[0]] = nums[0]
    // for nums[currIdx] != nums[currIdx-1] {
    //     mapNums[nums[currIdx]] += nums[currIdx]
    //     currIdx++
    // }

    // nextIdx := currIdx+1
    // mapNums[nums[nextIdx-1]] = nums[nextIdx-1]
    // for nums[nextIdx] != nums[nextIdx-1] {
    //     nextIdx++
    // }

    // fmt.Println(nums)
    // return 0
}