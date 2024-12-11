package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var entryNumber int
var questions []Info
var file os.File
var filePath string

type Info struct {
	QuestionID int
	EasyWork   string
	EasyClue   string
	HardWork   string
	HardClue   string
	Author     string
	Answers    []string
}

func printInfoSlice() {
	for i := range questions {
		fmt.Printf("-------------------------------------------\n")
		fmt.Printf("Question ID: %d\n", questions[i].QuestionID)
		fmt.Printf("EasyWork: %s\n", questions[i].EasyWork)
		fmt.Printf("EasyClue: %s\n", questions[i].EasyClue)
		fmt.Printf("HardWork: %s\n", questions[i].HardWork)
		fmt.Printf("HardClue: %s\n", questions[i].HardClue)
		fmt.Printf("Author: %s\n", questions[i].Author)
		fmt.Printf("Answers: ")
		for j := range questions[i].Answers {
			fmt.Printf("%s, ", questions[i].Answers[j])
		}
		fmt.Printf("\n")
		fmt.Printf("-------------------------------------------\n")
	}
}

func addAnswers(answers []string) []string {
	var newAnswer string
	looping := "y"
	for looping == "y" {
		fmt.Printf(">> Add a new answer:\n")
		fmt.Scanln(&newAnswer)
		answers = append(answers, newAnswer)
		fmt.Printf(">> Add another? (y/n)\n")
		fmt.Scanln(&looping)
	}
	return answers

}

func addNewEntry() {
	var answer string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">> Add new entry? (y/n)")
	fmt.Scanln(&answer)
	if answer != "y" {
		return
	}
	var newEntry Info
	newEntry.QuestionID = entryNumber + 1
	entryNumber++
	fmt.Printf(">> Enter the easy work question:\n")
	newEntry.EasyWork, _ = reader.ReadString('\n')
	newEntry.EasyWork = strings.TrimSpace(newEntry.EasyWork)
	fmt.Printf(">> Enter an easy clue:\n")
	newEntry.EasyClue, _ = reader.ReadString('\n')
	newEntry.EasyClue = strings.TrimSpace(newEntry.EasyClue)
	fmt.Printf(">> Enter a hard work question:\n")
	newEntry.HardWork, _ = reader.ReadString('\n')
	newEntry.HardWork = strings.TrimSpace(newEntry.HardWork)
	fmt.Printf(">> Enter a hard clue:\n")
	newEntry.HardClue, _ = reader.ReadString('\n')
	newEntry.HardClue = strings.TrimSpace(newEntry.HardClue)
	fmt.Printf(">> Enter the author's name in full:\n")
	newEntry.Author, _ = reader.ReadString('\n')
	newEntry.Author = strings.TrimSpace(newEntry.Author)
	fmt.Printf(">> Enter any possible answer choices:\n")
	newEntry.Answers = addAnswers(newEntry.Answers)
	questions = append(questions, newEntry)
	printInfoSlice()
	saveFile(filePath)
	addNewEntry()
}

func saveFile(filePath string) {
	finalOutput, err := json.Marshal(questions)
	if err != nil {
		fmt.Printf("> Error marshalling JSON file into string, %s", err)
		return
	}
	os.WriteFile(filePath, finalOutput, 0777)
}

func main() {
	fmt.Printf("-----------------------\n")
	fmt.Printf("JFW Quizwriter v2 \n")
	fmt.Printf("  now with autosaving!\n")
	fmt.Printf("-----------------------\n")
	fmt.Printf("> Insert file path: \n")
	fmt.Scanln(&filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("> Error opening file path: %s", err)
		return
	}
	defer file.Close()
	fileByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("> Error converting to byte data: %s", err)
		return
	}
	err = json.Unmarshal(fileByte, &questions)
	if err != nil {
		fmt.Printf("> Error unmarshalling JSON file into slice: %s", err)
		return
	}
	entryNumber = len(questions) - 1
	printInfoSlice()
	addNewEntry()
	finalOutput, err := json.Marshal(questions)
	if err != nil {
		fmt.Printf("> Error marshalling JSON file into string: %s", err)
		return
	}
	file.Close()
	var answer string
	fmt.Printf("> Save to file? (y/n)\n")
	fmt.Scanln(&answer)
	if answer == "y" {
		os.WriteFile(filePath, finalOutput, 0777)
		file.Close()
		return
	}
	fmt.Printf("> Are you sure? Save to file? (y/n)\n")
	fmt.Scanln(&answer)
	if answer == "y" {
		os.WriteFile(filePath, finalOutput, 0777)
		file.Close()
		return
	}
}
