package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner
var writer *bufio.Writer

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func qpow(x, y int) int {
	res := 1
	for y > 0 {
		if y%2 == 1 {
			res = res * x
		}
		x = x * x
	}
	return res
}
func solve() {
	a := nextInt()
	b := nextInt()
	c := nextInt()
	if a == 0 {
		if c != 1 {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, b)
		}
	} else {
		tot := 0
		cur := 0
		for tot+qpow(2, cur) <= a {
			tot += qpow(2, cur)
			cur++
		}
		used := tot - a
		if used < 0 {
			used *= -1
		}
		extra := qpow(2, cur) - used
		root := 2*used + extra
		if root != c {
			fmt.Fprintln(writer, -1)
		} else {
			if b <= extra {
				if used > 0 || extra > 0 {
					fmt.Fprintln(writer, cur+1)
				} else {
					fmt.Fprintln(writer, cur)
				}
			} else {
				cur += (b + root - 1) / root
				fmt.Fprintln(writer, cur)
			}
		}
	}

}

func main() {
	t := nextInt()
	for i := 0; i < t; i++ {
		solve()
	}
	writer.Flush()
}
