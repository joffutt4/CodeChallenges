package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Println(recurseBaum(val))
	}
}

func recurseBaum(val int64) (arr []int) {
	var i int64
	arr = append(arr, 1)
	for i = 1; i <= val; i++ {
		arr = append(arr, baumSweet(i))
	}
	return arr
}

func baumSweet(val int64) (baum int) {
	subs := strings.Split(strconv.FormatInt(val, 2), "1")
	for _, zeros := range subs {
		if len(zeros)%2 == 1 {
			return 0
		}
	}
	return 1
}
