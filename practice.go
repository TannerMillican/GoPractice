package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var highLevelTopicMap map[string]string

	highLevelTopicMap = make(map[string]string)

	highLevelTopicMap["Variables"] = "https://www.tutorialspoint.com/go/go_variables.htm"
	highLevelTopicMap["Data Types"] = "https://www.tutorialspoint.com/go/go_data_types.htm"
	highLevelTopicMap["Functions"] = "https://www.tutorialspoint.com/go/go_functions.htm"
	highLevelTopicMap["Arrays"] = "https://www.tutorialspoint.com/go/go_arrays.htm"
	highLevelTopicMap["Loops"] = "https://www.tutorialspoint.com/go/go_loops.htm"

	var intermediateTopicMap map[string]string

	intermediateTopicMap = make(map[string]string)

	intermediateTopicMap["Pointers"] = "https://www.tutorialspoint.com/go/go_pointers.htm"
	intermediateTopicMap["Structs"] = "https://www.tutorialspoint.com/go/go_structures.htm"
	intermediateTopicMap["Maps"] = "https://www.tutorialspoint.com/go/go_maps.htm"
	intermediateTopicMap["Recursion"] = "https://www.tutorialspoint.com/go/go_recursion.htm"
	intermediateTopicMap["Iterfaces"] = "https://www.tutorialspoint.com/go/go_interfaces.htm"

	addSpace()
	greetUser(highLevelTopicMap, intermediateTopicMap)

}

func addSpace() {

	fmt.Println(" ")

}

func greetUser(highLevelTopicMap map[string]string, intermediateTopicMap map[string]string) {

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
		askAboutLastHighLevelTopic(highLevelTopicMap, intermediateTopicMap)
	case timeInput == 365:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for 1 year!")
		addSpace()
		askAboutLastHighLevelTopic(highLevelTopicMap, intermediateTopicMap)
	case timeInput%365 == 0 && timeInput != 365:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for %d years!!", timeInput/365)
		fmt.Println("Happy studying!!")
	case timeInput > 365 && timeInput < 730:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for 1 year and %d day(s)!!", timeInput%365)
		askAboutLastIntermediateTopic(intermediateTopicMap)
	default:
		addSpace()
		fmt.Printf("Hello world! I've worked with Go for %d years and %d day(s)!!", timeInput/365, timeInput%365)
		fmt.Println("Happy studying!!")
	}

	addSpace()

}

func askAboutLastHighLevelTopic(highLevelTopicMap map[string]string, intermediateTopicMap map[string]string) {

	lastHigheLevelTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What was the last high level topic you studied?")
	fmt.Println("(if you would rather see a list of topics to choose from, or you haven't studied any high level topics, type: topics)")

	lastHigheLevelTopicScanner.Scan()

	lastHighLevelTopic := lastHigheLevelTopicScanner.Text()

	if lastHighLevelTopic == "topics" {
		addSpace()
		printHighLevelTopics(highLevelTopicMap)
		addSpace()
		askAboutStudyingHighLevelTopics(highLevelTopicMap, intermediateTopicMap)
	} else {
		addSpace()
		askAboutStudyingHighLevelTopics(highLevelTopicMap, intermediateTopicMap)
	}

}

func askAboutLastIntermediateTopic(intermediateTopicMap map[string]string) {

	lastIntermediateTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What was the last intermediate topic you studied?")
	fmt.Println("(if you would rather see a list of topics to choose from, or you haven't studied any intermediate topics, type: topics)")

	lastIntermediateTopicScanner.Scan()

	lastItermediateTopic := lastIntermediateTopicScanner.Text()

	if lastItermediateTopic == "topics" {
		addSpace()
		printIntermediateTopics(intermediateTopicMap)
		addSpace()
		askAboutStudyingIntermediateTopics(intermediateTopicMap)
	} else {
		addSpace()
		askAboutStudyingIntermediateTopics(intermediateTopicMap)
	}

}

func askAboutStudyingHighLevelTopics(highLevelTopicMap map[string]string, intermediateTopicMap map[string]string) {

	studyHighLevelScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to study a high level topic today? (Y/N)")
	studyHighLevelScanner.Scan()

	studyHighScannerYN := studyHighLevelScanner.Text()

	if studyHighScannerYN == "Y" || studyHighScannerYN == "y" {
		addSpace()
		fmt.Println("Which high level topic would you like to study today?")
		addSpace()
		printHighLevelTopics(highLevelTopicMap)
		chooseHighLevelTopic(highLevelTopicMap)
	} else if studyHighScannerYN == "N" || studyHighScannerYN == "n" {
		addSpace()
		askAboutStudyingIntermediateTopics(intermediateTopicMap)
	} else {
		addSpace()
		fmt.Println("Please enter Y or N")
		askAboutStudyingHighLevelTopics(highLevelTopicMap, intermediateTopicMap)
	}

}

func askAboutStudyingIntermediateTopics(intermediateTopicMap map[string]string) {

	studyIntermediateScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to study an intermediate topic today? (Y/N)")
	studyIntermediateScanner.Scan()

	studyIntermediateScannerYN := studyIntermediateScanner.Text()

	if studyIntermediateScannerYN == "Y" || studyIntermediateScannerYN == "y" {
		addSpace()
		fmt.Println("Which intermediate topic would you like to study today?")
		addSpace()
		printIntermediateTopics(intermediateTopicMap)
		chooseIntermediateTopic(intermediateTopicMap)
	} else if studyIntermediateScannerYN == "N" || studyIntermediateScannerYN == "n" {
		addSpace()
		fmt.Println("Happy studying!!")
	} else {
		addSpace()
		fmt.Println("Please enter Y or N")
		askAboutStudyingIntermediateTopics(intermediateTopicMap)
	}

}

func chooseHighLevelTopic(highLevelTopicMap map[string]string) {

	chooseHighLevelTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please capitalize the first letter of your choice.")
	addSpace()

	chooseHighLevelTopicScanner.Scan()

	chosenHighLevelTopic := chooseHighLevelTopicScanner.Text()

	if topic, found := highLevelTopicMap[chosenHighLevelTopic]; found {
		addSpace()
		fmt.Println("Here is a link to a tutorial website for your selected topic:")
		fmt.Println(topic)
		fmt.Println("Happy studying!!")
	} else {
		addSpace()
		fmt.Println("Sorry, we couldn't find that topic. Please try entering it again or try entering a new topic.")
		addSpace()
		chooseHighLevelTopic(highLevelTopicMap)
	}

}

func chooseIntermediateTopic(intermediateTopicMap map[string]string) {

	chooseIntermediateTopicScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please capitalize the first letter of your choice.")
	addSpace()

	chooseIntermediateTopicScanner.Scan()

	chosenIntermediateTopic := chooseIntermediateTopicScanner.Text()

	if topic, found := intermediateTopicMap[chosenIntermediateTopic]; found {
		addSpace()
		fmt.Println("Here is a link to a tutorial website for your selected topic:")
		fmt.Println(topic)
		fmt.Println("Happy studying!!")
	} else {
		addSpace()
		fmt.Println("Sorry, we couldn't find that topic. Please try entering it again or try entering a new topic.")
		addSpace()
		chooseHighLevelTopic(intermediateTopicMap)
	}

}

func printHighLevelTopics(highLevelTopicMap map[string]string) {

	for topic := range highLevelTopicMap {
		fmt.Println(topic)
	}

}

func printIntermediateTopics(intermediateTopicMap map[string]string) {

	for topic := range intermediateTopicMap {
		fmt.Println(topic)
	}

}
