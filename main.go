package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "myText.txt"

	_, err := os.Stat(filename)
	if err == nil {
		os.Remove(filename)
	}

	writeFile("myText.txt", strings.Join(os.Args[1:], " "))

	text := readFromFile("myText.txt")
	s := strings.Fields(text)

	characterCount := countCharacters(text)
	wordCount := countWords(text)
	lineCount := countLines(text)
	m := countFrequency(s)

	printingFrequnecy(m, s)

	fmt.Printf(" \nCharacter count: %d\n", characterCount)
	fmt.Printf("Word count: %d\n", wordCount)
	fmt.Printf("Line count: %d\n", lineCount)
}

func writeFile(filename string, s string) error {
	err := os.WriteFile(filename, []byte(s), os.ModeDir)
	if err != nil {
		fmt.Printf("Some thing went wrong !! %v", err)
		os.Exit(1)
	}
	return err
}

func countCharacters(text string) int {
	return len(text)
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func readFromFile(filename string) string {
	resp, err := os.ReadFile(filename)

	if err != nil {
		os.Exit(1)
	}

	s := strings.Split(string(resp), "\n")

	return strings.Join(s, "\n")

}

func countLines(text string) int {
	lines := strings.Split(text, "\n")
	return len(lines)
}

func countFrequency(s []string) map[string]int {
	m := make(map[string]int)

	for i := range s {
		count := 0
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] {
				count++
			}
		}
		m[s[i]] = count
	}

	return m
}

func printingFrequnecy(m map[string]int, s []string) {
	visited := make(map[string]bool)

	fmt.Println("Printing frequency of each word : ")
	for i := range s {

		if !visited[s[i]] {
			fmt.Print(s[i], " => ", m[s[i]], " time ", "  ,  ")
		}

		visited[s[i]] = true
	}
}
