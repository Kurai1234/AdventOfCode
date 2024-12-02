package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/kurai1234/advent-of-code/utils"
)

func ReadColumns(scanner *bufio.Scanner) ([]int, []int, error) {
	columnOne, columnTwo := make([]int, 0), make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		filteredLine := make([]string, 0)
		for _, val := range strings.Split(line, " ") {
			if val == "" {
				continue
			}
			filteredLine = append(filteredLine, val)
		}

		if num, err := strconv.Atoi(filteredLine[0]); err != nil {
			return nil, nil, err
		} else {
			columnOne = append(columnOne, num)
		}

		if num, err := strconv.Atoi(filteredLine[len(filteredLine)-1]); err != nil {
			return nil, nil, err
		} else {
			columnTwo = append(columnTwo, num)
		}
	}
	return columnOne, columnTwo, nil
}

func main() {
	scanner, err := utils.ScanFile("./dayOne/input.txt")

	if err != nil {
		panic(err)
	}

	columnOne, columnTwo, err := ReadColumns(scanner)

	if err != nil {
		panic(err)
	}

	if len(columnOne) != len(columnTwo) {
		panic(fmt.Errorf("column does not match"))
	}

	// Question 1
	sum := 0

	sort.Ints(columnOne)
	sort.Ints(columnTwo)

	for idx := range columnOne {
		sum += int(math.Abs(float64(columnOne[idx]) - float64(columnTwo[idx])))
	}
	fmt.Printf("Test One: %d \n", sum)

	sum = 0

	for _, colOne := range columnOne {
		occur := 0
		for _, colTwo := range columnTwo {
			if colOne == colTwo {
				occur += 1
			}
		}
		sum += occur * colOne
	}

	fmt.Printf("Test Two: %d \n", sum)

}
