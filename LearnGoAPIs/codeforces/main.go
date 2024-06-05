package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	tc := readInt(in)

	for t_ := 0; t_ < tc; t_++ {
		n := readInt(in)
		a := readArrInt(in)
		mx := -1

		for i := 0; i < n; i++ {
			mx = max(mx, a[i])
		}
		res := make([]int, n)
		if mx == -1 {

			for i := 0; i < n; i++ {
				if i%2 == 0 {
					res[i] = 1
				} else {
					res[i] = 2
				}
			}
			fmt.Println(res)
		} else {

		}
	}

}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}
func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")

	return numbs
}
func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}
