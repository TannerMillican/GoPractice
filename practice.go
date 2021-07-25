package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	addSpace()
	greetUser()

}

func addSpace() {
	fmt.Println(" ")
}

func greetUser() {

	greetingsScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, how long have you been working with Go?")
	fmt.Println("(please answer in number of days):")

	greetingsScanner.Scan()
	timeInput, _ := strconv.ParseInt(greetingsScanner.Text(), 10, 64)

	switch {
	case timeInput < 365:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for %d days", timeInput)
		addSpace()
		askAboutLastHighLevelTopic()
	case timeInput == 365:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for 1 year!")
		askAboutLastHighLevelTopic()
	case timeInput%365 == 0 && timeInput != 365:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for %d years!!", timeInput/365)
		fmt.Println("Happy studying!!")
	case timeInput > 365 && timeInput < 730:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for 1 year and %d day(s)!!", timeInput%365)
		askAboutLastIntermediateTopic()
	default:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for %d years and %d day(s)!!", timeInput/365, timeInput%365)
		fmt.Println("Happy studying!!")
	}

	addSpace()
}

func askAboutLastHighLevelTopic() {

	lastHigheLevelTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What was the last high level topic you studied?")
	fmt.Println("(if you would rather see a list of topics to choose from, or you haven't studied any high level topics, type: topics)")

	lastHigheLevelTopicScanner.Scan()

	lastHighLevelTopic := lastHigheLevelTopicScanner.Text()

	if lastHighLevelTopic == "topics" {
		addSpace()
		printHighLevelTopics()
		addSpace()
		askAboutStudyingHighLevelTopics()
	} else {
		addSpace()
		askAboutStudyingHighLevelTopics()
	}

}

func askAboutLastIntermediateTopic() {

	lastIntermediateTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What was the last intermediate topic you studied?")
	fmt.Println("(if you would rather see a list of topics to choose from, or you haven't studied any intermediate topics, type: topics)")

	lastIntermediateTopicScanner.Scan()

	lastItermediateTopic := lastIntermediateTopicScanner.Text()

	if lastItermediateTopic == "topics" {
		addSpace()
		printIntermediateTopics()
		addSpace()
		askAboutStudyingIntermediateTopics()
	} else {
		addSpace()
		askAboutStudyingIntermediateTopics()
	}

}

func askAboutStudyingHighLevelTopics() {
	studyHighLevelScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to study a high level topic today? (Y/N)")
	studyHighLevelScanner.Scan()

	studyHighScannerYN := studyHighLevelScanner.Text()

	if studyHighScannerYN == "Y" || studyHighScannerYN == "y" {
		addSpace()
		fmt.Println("Which high level topic would you like to study today?")
		printHighLevelTopics()
		chooseHighLevelTopic()
	} else if studyHighScannerYN == "N" || studyHighScannerYN == "n" {
		addSpace()
		askAboutStudyingIntermediateTopics()
	} else {
		addSpace()
		fmt.Println("Please enter Y or N")
		askAboutStudyingHighLevelTopics()
	}
}

func askAboutStudyingIntermediateTopics() {
	studyIntermediateScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to study an intermediate topic today? (Y/N)")
	studyIntermediateScanner.Scan()

	studyIntermediateScannerYN := studyIntermediateScanner.Text()

	if studyIntermediateScannerYN == "Y" || studyIntermediateScannerYN == "y" {
		addSpace()
		fmt.Println("Which intermediate topic would you like to study today?")
		printIntermediateTopics()
		chooseIntermediateTopic()
	} else if studyIntermediateScannerYN == "N" || studyIntermediateScannerYN == "n" {
		addSpace()
		fmt.Println("Happy studying!!")
	} else {
		addSpace()
		fmt.Println("Please enter Y or N")
		askAboutStudyingIntermediateTopics()
	}
}

func chooseHighLevelTopic() {
	highLevelTopicLinks := [5]string{"https://www.tutorialspoint.com/go/go_variables.htm", "https://www.tutorialspoint.com/go/go_data_types.htm", "https://www.tutorialspoint.com/go/go_functions.htm", "https://www.tutorialspoint.com/go/go_arrays.htm", "https://www.tutorialspoint.com/go/go_loops.htm"}

	highLevelTopics := [5]string{"variables", "data types", "functions", "arrays", "loops"}

	chooseHighLevelTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your choice without any capital letters:")

	chooseHighLevelTopicScanner.Scan()

	chosenHighLevelTopic := chooseHighLevelTopicScanner.Text()

	for i := 0; i < len(highLevelTopics); i++ {
		if chosenHighLevelTopic == highLevelTopics[i] {
			fmt.Println("Here is a link to a tutorial website for your selected topic:")
			fmt.Println(highLevelTopicLinks[i])
			fmt.Println("Happy studying!!")
		}
	}
}

func chooseIntermediateTopic() {
	intermediateTopicLinks := [5]string{"https://www.tutorialspoint.com/go/go_pointers.htm", "https://www.tutorialspoint.com/go/go_structures.htm", "https://www.tutorialspoint.com/go/go_maps.htm", "https://www.tutorialspoint.com/go/go_recursion.htm", "https://www.tutorialspoint.com/go/go_interfaces.htm"}

	intermediateTopics := [5]string{"Pointers", "Structs", "Maps", "Recursion", "Iterfaces"}

	chooseIntermediateTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your choice without any capital letters:")

	chooseIntermediateTopicScanner.Scan()

	chosenHighLevelTopic := chooseIntermediateTopicScanner.Text()

	for i := 0; i < len(intermediateTopics); i++ {
		if chosenHighLevelTopic == intermediateTopics[i] {
			fmt.Println("Here is a link to a tutorial website for your selected topic:")
			fmt.Println(intermediateTopicLinks[i])
			fmt.Println("Happy studying!!")
		}
	}
}

func printHighLevelTopics() {
	highLevelTopics := [5]string{"Variables", "Data Types", "Functions", "Arrays", "Loops"}

	for i := 0; i < len(highLevelTopics); i++ {
		fmt.Println(highLevelTopics[i])
	}
}

func printIntermediateTopics() {
	intermediateTopics := [5]string{"Pointers", "Structs", "Maps", "Recursion", "Iterfaces"}

	for i := 0; i < len(intermediateTopics); i++ {
		fmt.Println(intermediateTopics[i])
	}
}
