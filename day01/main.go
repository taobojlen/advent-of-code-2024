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

func solve(input string) int {
    // Your solution here
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

func main() {
    // Read input from input.txt
    input, _ := os.ReadFile("./day01/input.txt")
    result := solve(string(input))
    fmt.Printf("Result: %d\n", result)
}