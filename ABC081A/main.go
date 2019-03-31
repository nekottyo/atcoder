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
	input := scanner.Text()
	var r int
	for _, a := range input {
		n, _ := strconv.Atoi(string(rune(a)))
		r += n
	}

	fmt.Println(r)
}
