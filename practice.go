package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	practice_scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("How many days have you been working with Go?")
	practice_scanner.Scan()
	time_input, _ := strconv.ParseInt(practice_scanner.Text(), 10, 64)
	fmt.Printf("Hello world! I've worked with Go for %d days", time_input)

}
