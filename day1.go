package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getLists() ([]int, []int) {
	file, err := os.Open("inputs/day1input.txt")
	if err != nil {
		panic("Input file for day 1 was not found")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1, list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")

		if len(numbers) != 2 {
			panic("Invalid input, each line must have exactly 2 numbers")
		}

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			panic("Invalid input, only numeric inputs are allowed")
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)

	}

	return list1, list2
}

func Day1() {

	list1, list2 := getLists()

	slices.Sort(list1)
	slices.Sort(list2)

	diff := 0
	for i := 0; i < len(list1); i++ {
		diff += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println(diff)

	counter1 := make(map[int]int)

	for _, num := range list1 {
		counter1[num] += 1
	}


	counter2 := make(map[int]int)

	for _, num := range list2 {
		counter2[num] += 1
	}

	similarityScore := 0
	for key, value := range counter1 {
		value2 := counter2[key]
		similarityScore += key * value * value2
	}
	
	fmt.Println("Similarity score:", similarityScore)
}
