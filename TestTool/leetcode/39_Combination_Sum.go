package main

import (
	"fmt"
	"sort"
)

func main() {
	candiates := []int{2,3,5}
	target := 8
	solutions := combinationSum(candiates, target)
	fmt.Println(solutions)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var solutions [][]int
	var single []int

	DFS(&solutions, single, candidates, target, 0)
	return solutions
}

func DFS(solutions *[][]int, single, candidates []int, target, sum int) {
	for i, candidate := range candidates {
		sumNow := sum + candidate
		if sumNow == target {
			singleRes := make([]int, len(single) + 1)
			copy(singleRes, append(single, candidate))
			*solutions = append(*solutions, singleRes)
			break
		} else if sumNow < target {
			DFS(solutions, append(single, candidate), candidates[i:], target, sumNow)
		} else {
			break
		}
	}
}