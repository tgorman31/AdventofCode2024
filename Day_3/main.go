package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	str := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	dat, err := os.ReadFile("dataP1.txt")
	check(err)
	input := ""
	stringifyDat := string(dat)

	test := false
	include := true

	if test {
		input = str
	} else {
		input = stringifyDat
	}

	r, _ := regexp.Compile(`(?:mul\([0-9]{1,3}\,[0-9]{1,3}\))|(?:do\(\))|(?:don\'t\(\))`)
	r2, _ := regexp.Compile(`[0-9]{1,3}\,[0-9]{1,3}`)
	total := int64(0)

	for _, match := range r.FindAllString(input, -1) {

		switch match {
		case "don't()":
			include = false
		case "do()":
			include = true
		default:
			if include {
				match2 := r2.FindString(match)
				num := s.Split(match2, ",")
				val1, _ := strconv.ParseInt(num[0], 0, 64)
				val2, _ := strconv.ParseInt(num[1], 0, 64)
				product := val1 * val2
				total = total + product
			}

		}
	}

	fmt.Println(total)

}
