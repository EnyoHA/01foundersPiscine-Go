package main

import (
	"fmt"
	"strconv"
	"strings"

	// fmt allows the package to format all ouputs and has many built in functions depending on your use
	"io/ioutil"
	// allows you to interact with files and to read them.
	"os"
	// allows you to open and create files on your computer
	//"https://git.learn.01founders.co/EnyoHA/go-reloaded/right/simpleCode.git"
	// this is an internal package written to handle the logic split up into different file functions.
)

// all of the functions I have made for alternated answers actually exist already in the packages

func hexChange(str string) string {
	// strconv.ParseUint interprets a string from the given base
	// and the bit size. EG the 16 and the 64.
	toChange, _ := strconv.ParseUint(str, 16, 64)
	//
	return strconv.FormatUint(toChange, 10)
}

func binChange(str string) string {
	// same as hex but start from base 2
	toChange, _ := strconv.ParseUint(str, 2, 64)

	return strconv.FormatUint(toChange, 10)
}

func upperChange(str string) string {
	return strings.ToUpper(str)
}

func lowerChange(str string) string {
	return strings.ToLower(str)
}

func capitalChange(str string) string {
	return strings.Title(str)
}

func puncChange(input string, punctuation string) string {
	fields := strings.Split(input, punctuation)
	finalArray := make([]string, 0)
	for i, field := range fields {
		if len(field) == 0 {
			finalArray[len(finalArray)-1] = finalArray[len(finalArray)-1] + punctuation
			continue
		}
		fields[i] = strings.TrimSpace(field)
		finalArray = append(finalArray, fields[i])
	}

	return strings.Join(finalArray, punctuation+" ")
}

func startsWithVowel(word string) bool {
	if string(word[0]) == "a" ||
		string(word[0]) == "e" ||
		string(word[0]) == "i" ||
		string(word[0]) == "o" ||
		string(word[0]) == "u" {
		return true
	}

	return false
}

func vowelChange(input string) string {
	fields := strings.Fields(input)
	for i, field := range fields {
		if field == "a" {
			toCheck := string(fields[i+1])
			if startsWithVowel(toCheck) || string(toCheck[0]) == "h" {
				fields[i] = "an"
			}
		}
	}

	return strings.Join(fields, " ")
}

// we are calling the first file, and then outputting the right file needed from it.
func main() {
	inputFileName := os.Args[1]                        // passing an argument into input.txt, input.txt is the 1st argument
	outputFileName := os.Args[2]                       // passing the result.txt output after its run through the functions. output.txt 2nd argument
	fileContent, err := ioutil.ReadFile(inputFileName) // using ioutil to read the input file data and then running the functions through it
	if err != nil {                                    // If we run into a problem , throw an error. != means not equals
		panic(err) // if there is an error it will panic and throw an exception
	}
	// fmt.Println(string(fileContent))
	output := hexChange(string(fileContent))
	output = binChange(string(output))
	output = upperChange(string(output))
	output = lowerChange(string(output))
	output = capitalChange(string(output))
	output = puncChange(string(output), ".")
	output = puncChange(string(output), ",")
	output = puncChange(string(output), "!")
	output = puncChange(string(output), "?")
	output = puncChange(string(output), ":")
	output = puncChange(string(output), ";")
	output = vowelChange(string(output))
	// create the file
	outputFile, err := os.Create(outputFileName) // after we have computed everything we hve to now create a file passing through an output of OS(this is the argument)
	if err != nil {                              // ideally there wouldnt be an error, but hypothetically if you dont have space it may not work
		fmt.Println(err)
	}
	// close the file with defer
	defer outputFile.Close() // when you put defer before a function you are saying it must run it, and run it last.
	// write a string
	outputFile.WriteString(output) // Once everything is computed in the functions, it will then write the output to file as a string.
}
