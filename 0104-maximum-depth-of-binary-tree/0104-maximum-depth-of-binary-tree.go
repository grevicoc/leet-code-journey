/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 import (
     "fmt"
 )
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    hist := make([]*TreeNode, 0)
    levelHist := make([]int, 0)

    maxLevel := 1
    currNode := (*root).Left
    currLevel := 1
    if (*root).Right != nil {
        hist = append(hist,(*root).Right)
        levelHist = append(levelHist, currLevel)
        maxLevel = currLevel+1
    }

    for len(hist) != 0 || currNode != nil {
        if currNode != nil {
            currLevel++
            if currNode.Right != nil {
                hist = append(hist,currNode.Right)
                levelHist = append(levelHist, currLevel)
            }
            
            if currLevel > maxLevel {
                maxLevel = currLevel
            }
            
            currNode = currNode.Left
            continue
        }
        currNode = hist[len(hist)-1]
        currLevel = levelHist[len(levelHist)-1]

        hist = hist[:len(hist)-1]
        levelHist = levelHist[:len(levelHist)-1]
    }

    return maxLevel
}