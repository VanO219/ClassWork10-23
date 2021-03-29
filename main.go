package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(col(23536234523.60))
}

func col(num float64) (numS string) {
	temp := (int64)(num)

	a := (int64)(num * 100)
	point := a % 100

	// str := strconv.FormatInt(temp, 10)

	// b1 := []byte(str)


	var arr []string

	i := 0
	for temp > 0 {
		arr = append(arr, strconv.FormatInt(temp % 10, 10))
		temp /= 10
	}

	for i = 0; i < len(arr); i++ {
		numS += string(arr[i])
		if (i+1)%3 == 0 && (i+1) < len(arr) {
			numS += "__"
		}
	}

	b := []byte(numS)

	c := []byte{}
	for i = (len(b) - 1); i == 0; i-- {
		c = append(c, b[i])
	}
	numS = string(c)
	numS += fmt.Sprintf(".%d", point)

	return
}
