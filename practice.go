package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type HighLevelTopic struct {
	name         string
	tutorialLink string
}

type IntermediateTopic struct {
	name         string
	tutorialLink string
}

func main() {

	var HighLevelTopic1 HighLevelTopic
	var HighLevelTopic2 HighLevelTopic
	var HighLevelTopic3 HighLevelTopic
	var HighLevelTopic4 HighLevelTopic
	var HighLevelTopic5 HighLevelTopic

	HighLevelTopic1.name = "Variables"
	HighLevelTopic1.tutorialLink = "https://www.tutorialspoint.com/go/go_variables.htm"

	HighLevelTopic2.name = "Data Types"
	HighLevelTopic2.tutorialLink = "https://www.tutorialspoint.com/go/go_data_types.htm"

	HighLevelTopic3.name = "Functions"
	HighLevelTopic3.tutorialLink = "https://www.tutorialspoint.com/go/go_functions.htm"

	HighLevelTopic4.name = "Arrays"
	HighLevelTopic4.tutorialLink = "https://www.tutorialspoint.com/go/go_arrays.htm"

	HighLevelTopic5.name = "Loops"
	HighLevelTopic5.tutorialLink = "https://www.tutorialspoint.com/go/go_loops.htm"

	highLevelTopicsArray := [5]HighLevelTopic{HighLevelTopic1, HighLevelTopic2, HighLevelTopic3, HighLevelTopic4, HighLevelTopic5}

	var ItermediateTopic1 IntermediateTopic
	var ItermediateTopic2 IntermediateTopic
	var ItermediateTopic3 IntermediateTopic
	var ItermediateTopic4 IntermediateTopic
	var ItermediateTopic5 IntermediateTopic

	ItermediateTopic1.name = "Pointers"
	ItermediateTopic1.tutorialLink = "https://www.tutorialspoint.com/go/go_pointers.htm"

	ItermediateTopic2.name = "Structs"
	ItermediateTopic2.tutorialLink = "https://www.tutorialspoint.com/go/go_structures.htm"

	ItermediateTopic3.name = "Maps"
	ItermediateTopic3.tutorialLink = "https://www.tutorialspoint.com/go/go_maps.htm"

	ItermediateTopic4.name = "Recursion"
	ItermediateTopic4.tutorialLink = "https://www.tutorialspoint.com/go/go_recursion.htm"

	ItermediateTopic5.name = "Iterfaces"
	ItermediateTopic5.tutorialLink = "https://www.tutorialspoint.com/go/go_interfaces.htm"

	intermediateTopicsArray := [5]IntermediateTopic{ItermediateTopic1, ItermediateTopic2, ItermediateTopic3, ItermediateTopic4, ItermediateTopic5}

	addSpace()
	greetUser(highLevelTopicsArray, intermediateTopicsArray)

}

func addSpace() {
	fmt.Println(" ")
}

func greetUser(highLevelTopicsArray [5]HighLevelTopic, intermediateTopicsArray [5]IntermediateTopic) {

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
		askAboutLastHighLevelTopic(highLevelTopicsArray, intermediateTopicsArray)
	case timeInput == 365:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for 1 year!")
		addSpace()
		askAboutLastHighLevelTopic(highLevelTopicsArray, intermediateTopicsArray)
	case timeInput%365 == 0 && timeInput != 365:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for %d years!!", timeInput/365)
		fmt.Println("Happy studying!!")
	case timeInput > 365 && timeInput < 730:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for 1 year and %d day(s)!!", timeInput%365)
		askAboutLastIntermediateTopic(intermediateTopicsArray)
	default:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for %d years and %d day(s)!!", timeInput/365, timeInput%365)
		fmt.Println("Happy studying!!")
	}

	addSpace()
}

func askAboutLastHighLevelTopic(highLevelTopicsArray [5]HighLevelTopic, intermediateTopicsArray [5]IntermediateTopic) {

	lastHigheLevelTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What was the last high level topic you studied?")
	fmt.Println("(if you would rather see a list of topics to choose from, or you haven't studied any high level topics, type: topics)")

	lastHigheLevelTopicScanner.Scan()

	lastHighLevelTopic := lastHigheLevelTopicScanner.Text()

	if lastHighLevelTopic == "topics" {
		addSpace()
		printHighLevelTopics(highLevelTopicsArray)
		addSpace()
		askAboutStudyingHighLevelTopics(highLevelTopicsArray, intermediateTopicsArray)
	} else {
		addSpace()
		askAboutStudyingHighLevelTopics(highLevelTopicsArray, intermediateTopicsArray)
	}

}

func askAboutLastIntermediateTopic(intermediateTopicsArray [5]IntermediateTopic) {

	lastIntermediateTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What was the last intermediate topic you studied?")
	fmt.Println("(if you would rather see a list of topics to choose from, or you haven't studied any intermediate topics, type: topics)")

	lastIntermediateTopicScanner.Scan()

	lastItermediateTopic := lastIntermediateTopicScanner.Text()

	if lastItermediateTopic == "topics" {
		addSpace()
		printIntermediateTopics(intermediateTopicsArray)
		addSpace()
		askAboutStudyingIntermediateTopics(intermediateTopicsArray)
	} else {
		addSpace()
		askAboutStudyingIntermediateTopics(intermediateTopicsArray)
	}

}

func askAboutStudyingHighLevelTopics(highLevelTopicsArray [5]HighLevelTopic, intermediateTopicsArray [5]IntermediateTopic) {
	studyHighLevelScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to study a high level topic today? (Y/N)")
	studyHighLevelScanner.Scan()

	studyHighScannerYN := studyHighLevelScanner.Text()

	if studyHighScannerYN == "Y" || studyHighScannerYN == "y" {
		addSpace()
		fmt.Println("Which high level topic would you like to study today?")
		printHighLevelTopics(highLevelTopicsArray)
		chooseHighLevelTopic(highLevelTopicsArray)
	} else if studyHighScannerYN == "N" || studyHighScannerYN == "n" {
		addSpace()
		askAboutStudyingIntermediateTopics(intermediateTopicsArray)
	} else {
		addSpace()
		fmt.Println("Please enter Y or N")
		askAboutStudyingHighLevelTopics(highLevelTopicsArray, intermediateTopicsArray)
	}
}

func askAboutStudyingIntermediateTopics(intermediateTopicsArray [5]IntermediateTopic) {
	studyIntermediateScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to study an intermediate topic today? (Y/N)")
	studyIntermediateScanner.Scan()

	studyIntermediateScannerYN := studyIntermediateScanner.Text()

	if studyIntermediateScannerYN == "Y" || studyIntermediateScannerYN == "y" {
		addSpace()
		fmt.Println("Which intermediate topic would you like to study today?")
		printIntermediateTopics(intermediateTopicsArray)
		chooseIntermediateTopic(intermediateTopicsArray)
	} else if studyIntermediateScannerYN == "N" || studyIntermediateScannerYN == "n" {
		addSpace()
		fmt.Println("Happy studying!!")
	} else {
		addSpace()
		fmt.Println("Please enter Y or N")
		askAboutStudyingIntermediateTopics(intermediateTopicsArray)
	}
}

func chooseHighLevelTopic(highLevelTopicsArray [5]HighLevelTopic) {

	chooseHighLevelTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please capitalize the first letter of your choice.")
	addSpace()

	chooseHighLevelTopicScanner.Scan()

	chosenHighLevelTopic := chooseHighLevelTopicScanner.Text()

	for i := 0; i < len(highLevelTopicsArray); i++ {
		if chosenHighLevelTopic == highLevelTopicsArray[i].name {
			addSpace()
			fmt.Println("Here is a link to a tutorial website for your selected topic:")
			fmt.Println(highLevelTopicsArray[i].tutorialLink)
			fmt.Println("Happy studying!!")
		}
	}
}

func chooseIntermediateTopic(intermediateTopicsArray [5]IntermediateTopic) {

	chooseIntermediateTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please capitalize the first letter of your choice.")
	addSpace()

	chooseIntermediateTopicScanner.Scan()

	chosenHighLevelTopic := chooseIntermediateTopicScanner.Text()

	for i := 0; i < len(intermediateTopicsArray); i++ {
		if chosenHighLevelTopic == intermediateTopicsArray[i].name {
			addSpace()
			fmt.Println("Here is a link to a tutorial website for your selected topic:")
			fmt.Println(intermediateTopicsArray[i].tutorialLink)
			fmt.Println("Happy studying!!")
		}
	}
}

func printHighLevelTopics(highLevelTopicsArray [5]HighLevelTopic) {

	for i := 0; i < len(highLevelTopicsArray); i++ {
		fmt.Println(highLevelTopicsArray[i].name)
	}
}

func printIntermediateTopics(intermediateTopicsArray [5]IntermediateTopic) {

	for i := 0; i < len(intermediateTopicsArray); i++ {
		fmt.Println(intermediateTopicsArray[i].name)
	}
}
