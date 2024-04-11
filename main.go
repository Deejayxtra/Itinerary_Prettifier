package main

import (
	"flag"
	"fmt"
	"os"
)

type Match struct {
	Index int    // sort the matches
	Value string // matched text
	Type  string
}

func main() {
	helpFlag := flag.Bool("h", false, "Display help")
	flag.Parse()

	if *helpFlag {
		fmt.Println("itinerary usage:")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}

	inputFilePath := "./input.txt"
	outputFilePath := "./output.txt"
	csvFilePath := "./airport-lookup.csv"

	csvFile, err := openCSV(csvFilePath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer csvFile.Close()

	header, err := readCSVHeader(csvFile)
	if err != nil {
		fmt.Println("Airport lookup malformed.", err)
		return
	}

	iataIndex, icaoIndex, nameIndex := findColumnIndices(header)
	if iataIndex == -1 || icaoIndex == -1 || nameIndex == -1 {
		fmt.Println("Airport lookup malformed.")
		return
	}

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Input not found.")
		return
	}
	defer inputFile.Close()

	output, err := processInput(inputFile, csvFile, iataIndex, icaoIndex, nameIndex)
	if err != nil {
		fmt.Println("error:")
		return
	}

	if err := writeOutput(outputFilePath, output); err != nil {
		fmt.Println("error:", err)
	}
}

// write the output string to a file.
func writeOutput(filename, output string) error {
    if output == "" {
        return nil
    }
    
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    _, err = file.WriteString(output)
    if err != nil {
        return err
    }
    
    return nil
}
