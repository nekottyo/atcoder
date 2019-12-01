package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var a, b, c int
	var s string

	if sc.Scan() {
		a, _ = strconv.Atoi(sc.Text())
	}

	if sc.Scan() {
		str := sc.Text()
		var strs []string
		strs = strings.Split(str, " ")

		b, _ = strconv.Atoi(strs[0])
		c, _ = strconv.Atoi(strs[1])
	}

	if sc.Scan() {
		s = sc.Text()
	}

	fmt.Printf("%d %s\n", a+b+c, s)
}
