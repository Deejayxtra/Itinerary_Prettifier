package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// process the input file and return the output string.
func processInput(inputFile *os.File, csvFile *os.File, iataIndex, icaoIndex, nameIndex int) (string, error) {
	scanner := bufio.NewScanner(inputFile)
	// use regexp to recognize patterns for codes and dates
	iataRegex, _ := regexp.Compile(`#[A-Z]{3}`)
	icaoRegex, _ := regexp.Compile(`##[A-Z]{4}`)
	dateRegex, _ := regexp.Compile(`([DT])(\d{2})?\((\d{4}-\d{2}-\d{2}T\d{2}:\d{2}(Z|[\+\-]\d{2}:\d{2}))\)`)
	var output strings.Builder

	prevLineEmpty := false // flag to track previous line was empty
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// If the line is empty and previous line was empty, skip writing to output
		if line == "" && prevLineEmpty {
			continue
		}

		// Add newline if previous line was not empty
		if prevLineEmpty {
			output.WriteString("\n")
		}

		// trim white space and replace consecutive white space with a single blank line
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "\v", "\n")
		line = strings.ReplaceAll(line, "\f", "\n")
		line = strings.ReplaceAll(line, "\r", "\n")

		re := regexp.MustCompile(`\n{3,}`)
		line = re.ReplaceAllString(line, "\n\n")

		matches := findMatches(line, iataRegex, icaoRegex, dateRegex)
		output.WriteString(processAllMatches(matches, line, csvFile, iataIndex, icaoIndex, nameIndex))

		// Update flag
		prevLineEmpty = line == ""
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading from input file: %v", err)
	}
	return output.String(), nil
}

// find all IATA and ICAO code and date matches.
func findMatches(line string, iataRegex, icaoRegex, dateRegex *regexp.Regexp) []Match {
	var matches []Match

	iataMatches := iataRegex.FindAllStringIndex(line, -1)
	for _, match := range iataMatches {
		matches = append(matches, Match{Index: match[0], Value: line[match[0]:match[1]], Type: "iata"})
	}

	icaoMatches := icaoRegex.FindAllStringIndex(line, -1)
	for _, match := range icaoMatches {
		matches = append(matches, Match{Index: match[0], Value: line[match[0]:match[1]], Type: "icao"})
	}

	dateMatches := dateRegex.FindAllStringIndex(line, -1)
	for _, match := range dateMatches {
		matches = append(matches, Match{Index: match[0], Value: line[match[0]:match[1]], Type: "date"})
	}

	return matches
}

// process matches and return the result string.
func processAllMatches(matches []Match, line string, csvFile *os.File, iataIndex, icaoIndex, nameIndex int) string {
	// sort matches by index
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Index < matches[j].Index
	})

	for _, match := range matches {
		var replacement string
		switch match.Type {
		case "iata", "icao":
			replacement = codeLookup(match.Value, csvFile, iataIndex, icaoIndex, nameIndex)
		case "date":
			adjustedTime, err := processLine(match.Value)
			if err == nil {
				replacement = adjustedTime
			}
		}
		line = strings.Replace(line, match.Value, replacement, 1)
	}

	return line + "\n"
}

