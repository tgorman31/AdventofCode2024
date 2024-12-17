package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func search_grid(data []string, r int, c int, rows int, cols int, word string) bool {
	// end := len(word)
	count := 0
	directions := [][2]int{
		{-1, -1}, // Up-Left Diagonal
		// {1, 1},   // Down-Right Diagonal
		{1, -1}, // Down-Left Diagonal
		// {-1, 1},  // Up-Right Diagonal
	}
	// for i := 1; i < end; i++ {
	for _, d := range directions {

		ar := r + 1*d[0]
		or := r - 1*d[0]

		ac := c + 1*d[1]
		oc := c - 1*d[1]

		if ar < 0 || ac < 0 || ar >= rows || ac >= cols || or < 0 || oc < 0 || or >= rows || oc >= cols {
			return false
		}
		// print(data[ar][ac], " ", word[0])
		// print(data[or][oc], " ", word[2])
		// print(data[or][oc], " ", word[0])
		// print(data[ar][ac], " ", word[2])
		if (data[ar][ac] == word[0] && data[or][oc] == word[2]) || (data[or][oc] == word[0] && data[ar][ac] == word[2]) {
			count = count + 1
		}
	}
	// println(count)
	return count == 2

}

func main() {
	// input := ""
	file := ""

	test := false

	if test {
		file = "test.txt"
	} else {
		file = "data.txt"
	}

	f, err := os.Open(file)
	check(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Print(lines)

	// directions := [][2]int{
	// 	{0, 1},   // Right
	// 	{1, 0},   // Down
	// 	{0, -1},  // Left
	// 	{-1, 0},  // Up
	// 	{1, 1},   // Down-Right Diagonal
	// 	{1, -1},  // Down-Left Diagonal
	// 	{-1, 1},  // Up-Right Diagonal
	// 	{-1, -1}, // Up-Left Diagonal
	// }

	// directions := [][2]int{
	// 	{-1, -1}, // Up-Left Diagonal
	// 	// {1, 1},   // Down-Right Diagonal
	// 	{1, -1}, // Down-Left Diagonal
	// 	// {-1, 1},  // Up-Right Diagonal
	// }

	// MMMSXXMASM
	// MSAMXMSMSA
	// AMXSXMAAMM
	// MSAMASMSMX
	// XMASAMXAMM
	// XXAMMXXAMA
	// SMSMSASXSS
	// SAXAMASAAA
	// MAMMMXMMMM
	// MXMXAXMASX

	word := "MAS"
	// word1 := "MS"
	// if M is found then word1
	// if S is found then word2
	// word2 := "SM"
	rows := len(lines)
	cols := len(lines[0])
	count := 0

	for r := range rows {
		for c := range cols {
			if string(lines[r][c]) == "A" {
				// Search
				// for _, d := range directions {
				if search_grid(lines, r, c, rows, cols, word) {
					count = count + 1
					// }
				}
			}
		}
	}

	fmt.Println(count)
}
