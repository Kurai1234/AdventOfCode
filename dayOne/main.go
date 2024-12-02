package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/kurai1234/advent-of-code/utils"
)

func main() {
	scanner, err := utils.ScanFile("./dayOne/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	var inputs [][]string
	var input []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inputs = append(inputs, input)
			input = []string{}
			continue
		}
		input = append(input, line)
	}
	if len(input) > 0 {
		inputs = append(inputs, input)
	}

	for index, firstEntry := range inputs {
		columnOne := make([]int, 0)
		columnTwo := make([]int, 0)
		for _, row := range firstEntry {
			var filteredRow []int
			for _, val := range strings.Split(row, " ") {
				if val == "" {
					continue
				}

				if val, err := strconv.Atoi(val); err != nil {
					log.Fatal("strcon error:", err.Error())
				} else {
					filteredRow = append(filteredRow, val)
				}
			}
			if len(filteredRow) == 0 {
				continue
			}
			columnOne = append(columnOne, filteredRow[0])
			columnTwo = append(columnTwo, filteredRow[len(filteredRow)-1])
		}

		sort.Ints(columnOne)
		sort.Ints(columnTwo)

		if len(columnOne) != len(columnTwo) {
			log.Fatal("column do not match")
		}
		sum := 0
		for i := 0; i < len(columnOne); i++ {
			sum += int(math.Max(float64(columnOne[i]), float64(columnTwo[i]))) - int(math.Min(float64(columnOne[i]), float64(columnTwo[i])))
		}
		fmt.Printf("Test: %d, distance: %d", index, sum)
	}
}
