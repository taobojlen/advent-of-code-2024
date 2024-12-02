package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
	lines := strings.Split(input, "\n")
	var reports [][]int
	for _, line := range lines {
		var report []int
		if line == "" {
			continue
		}
		numbers := strings.Fields(line)
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return reports
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func filter[T any](list []T, test func(T) bool) (ret []T) {
	for _, s := range list {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func pop[T any](list []T, index int) []T {
	if index < 0 || index >= len(list) {
		panic("Index out of range")
	}
	ret := make([]T, 0)
	ret = append(ret, list[:index]...)
	return append(ret, list[index+1:]...)
}

func isSafe(report []int, problemDampener bool) bool {
	if len(report) < 2 {
		return true
	}
	positiveSign := (report[1] - report[0]) > 0
	for i, num1 := range report {
		if i == len(report)-1 {
			break
		}
		num2 := report[i+1]
		difference := num2 - num1
		if positiveSign != (difference > 0) {
			if problemDampener {
				// try isSafe after removing each index of the report
				safeWithRemoval := make([]bool, 0)
				for i := range report {
					newSlice := pop(report, i)
					safeWithRemoval = append(safeWithRemoval, isSafe(newSlice, false))
				}
				return len(filter(safeWithRemoval, func(b bool) bool { return b })) > 0
			}
			return false
		}
		if abs(difference) < 1 || abs(difference) > 3 {
			if problemDampener {
				safeWithRemoval := make([]bool, 0)
				for i := range report {
					newSlice := pop(report, i)
					safeWithRemoval = append(safeWithRemoval, isSafe(newSlice, false))
				}
				return len(filter(safeWithRemoval, func(b bool) bool { return b })) > 0
			}

			return false
		}
	}
	return true
}

func solve1(input string) int {
	reports := parseInput(input)
	safeReports := filter(reports, func(report []int) bool {
		return isSafe(report, false)
	})
	return len(safeReports)
}

func solve2(input string) int {
	reports := parseInput(input)
	safeReports := filter(reports, func(report []int) bool {
		return isSafe(report, true)
	})
	return len(safeReports)
}

func main() {
	input, _ := os.ReadFile("./day02/input.txt")
	part1Result := solve1(string(input))
	fmt.Printf("Part 1: %d\n", part1Result)
	part2Result := solve2(string(input))
	fmt.Printf("Part 2: %d\n", part2Result)
}
