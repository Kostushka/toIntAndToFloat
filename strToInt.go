package main

import (
	"fmt"
	"errors"
)

func pov(base, x int) int {
	res := 1
	for x > 0 {
		res *= base
		x--
	}
	return res
}

func strToFloat(str string) (float32, error) {
	var result float32
	countDot := 0
	countTens := 0
	for _, v := range str {
		if v == '.' {
			countDot++
			if countDot > 1 {
				return -1, errors.New("Это не число")
			}
			continue
		}
		if v < '0' || v > '9' {
			return -1, errors.New("Это не число")
		}
		if countDot > 0 {
			countTens++			
		}
		result *= 10.0
		result += float32(v - '0')
	}
	result *= 1.0 / float32(pov(10, countTens))
	return result, nil
}

func strToInt(str string) (int, error) {
	var result int
	for _, v := range str {
		if v < '0' || v > '9' {
			return -1, errors.New("Это не число")
		}
		result *= 10
		result += int(v - '0')
	}
	return result, nil
}

func main() {
	base := "2147483647"
	res, err := strToInt(base)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s %T - %d %T\n", base, base, res, res)
	}
	
	base2 := "15.478111"
	res2, err := strToFloat(base2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s %T - %f %T\n", base2, base2, res2, res2)
	}
}
