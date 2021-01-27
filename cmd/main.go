package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var d int64 = 200

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	answer := input / d
	fmt.Printf("LKR To Dollars : %d", answer)
}
