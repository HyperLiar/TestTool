package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	solutions := combinationSum2(candidates, 5)
	fmt.Println(solutions)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var solutions [][]int
	var single []int

	DFS(&solutions, single, candidates, target, 0)
	return solutions
}

func DFS(solutions *[][]int, single, candidates []int, target, sum int) {
	prev := -1
	for i, candidate := range candidates {
		if prev != -1 && candidate == prev {
			continue
		}
		sumNow := candidate + sum
		if sumNow == target {
			singleRes := make([]int, len(single)+1)
			copy(singleRes, append(single, candidate))
			*solutions = append(*solutions, singleRes)
			break
		} else if sumNow < target && i+1 < len(candidates) {
			DFS(solutions, append(single, candidate), candidates[i+1:], target, sumNow)
			prev = candidate
		} else {
			break
		}
	}
}
