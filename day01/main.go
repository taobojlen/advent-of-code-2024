package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "slices"
)

func zip(a, b []int) ([][2]int, error) {
    if len(a) != len(b) {
        return nil, fmt.Errorf("slices must be of equal length")
    }
    zip := make([][2]int, len(a))
    for i := range a {
        zip[i] = [2]int{a[i], b[i]}
    }
    return zip, nil
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func countOccurrences(list []int) map[int]int {
    occurrences := make(map[int]int)
    for _, num := range list {
        occurrences[num]++
    }
    return occurrences
}

func parseInput(input string) ([]int, []int) {
    lines := strings.Split(input, "\n")
    var list1, list2 []int
    for _, line := range lines {
        if line == "" {
            continue
        }
        numbers := strings.Fields(line)
        if len(numbers) != 2 {
            panic("Invalid input")
        }
        num1, err := strconv.Atoi(numbers[0])
        if err != nil {
            panic(err)
        }
        num2, err := strconv.Atoi(numbers[1])
        if err != nil {
            panic(err)
        }
        list1 = append(list1, num1)
        list2 = append(list2, num2)
    }
    return list1, list2
}

func solve1(input string) int {
    list1, list2 := parseInput(input)
    slices.Sort(list1)
    slices.Sort(list2)
    tuples, err := zip(list1, list2)
    if err != nil {
        panic(err)
    }
    difference := 0
    for _, tuple := range tuples {
        difference += abs(tuple[0] - tuple[1])
    }
    return difference
}

func solve2(input string) int {
    list1, list2 := parseInput(input)
    list2Counts := countOccurrences(list2)
    similarity := 0
    for _, num := range list1 {
        list2Count := list2Counts[num]
        similarity += (num * list2Count)
    }
    return similarity
}

func main() {
    // Read input from input.txt
    input, _ := os.ReadFile("./day01/input.txt")
    part1Result := solve1(string(input))
    part2Result := solve2(string(input))
    fmt.Printf("Result (part 1): %d\n", part1Result)
    fmt.Printf("Result (part 2): %d\n", part2Result)
}