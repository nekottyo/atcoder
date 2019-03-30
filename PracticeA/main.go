package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	bc := scanner.Text()

	var b, c int
	scanner.Scan()
	fmt.Sscanf(bc, "%d %d", &b, &c)
	s := scanner.Text()
	fmt.Printf("%d %s\n", a+b+c, s)
}
