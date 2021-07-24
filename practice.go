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

	switch {
	case time_input < 365:
		fmt.Printf("Hello world! I've worked with Go for %d days", time_input)
	case time_input%365 == 0:
		year_number := time_input%365 == 0
		if year_number {
			fmt.Printf("Hello world! I've worked with Go for 1 year!")
		} else {
			fmt.Printf("Hello world! I've worked with Go for %d years!!", time_input/365)
		}
	case time_input > 365 && time_input < 730:
		fmt.Printf("Hello world! I've worked with Go for 1 year and %d day(s)!!", time_input%365)
	default:
		fmt.Printf("Hello world! I've worked with Go for %d years and %d day(s)!!", time_input/365, time_input%365)
	}

}
