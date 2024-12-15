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
	// str := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	dat, err := os.ReadFile("dataP1.txt")
	check(err)

	stringifyDat := string(dat)

	// mul(44,46) 1-3 digit numbers
	r, _ := regexp.Compile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	r2, _ := regexp.Compile(`[0-9]{1,3}\,[0-9]{1,3}`)
	total := int64(0)

	for _, match := range r.FindAllString(stringifyDat, -1) {
		// fmt.Println(match, "found at index", i)
		match2 := r2.FindString(match)
		num := s.Split(match2, ",")
		val1, _ := strconv.ParseInt(num[0], 0, 64)
		val2, _ := strconv.ParseInt(num[1], 0, 64)
		product := val1 * val2
		total = total + product
	}

	fmt.Println(total)

}
