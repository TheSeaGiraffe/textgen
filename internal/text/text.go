package text

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	dummyFileName         = "dummy"
	dummyFileNameZerosPad = 3
	dummyText             = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud
exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure
dolor in reprehenderit in voluptate velit esse nulla pariatur. Excepteur sint occaecat
cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`
)

func shuffleText(text string) string {
	words := strings.Fields(strings.ToLower(text))
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	return strings.Join(words, " ")
}

func genText(numParagraphs int) []string {
	var newText []string
	for range numParagraphs {
		newText = append(newText, shuffleText(dummyText))
	}
	return newText
}

func genAndWrite(numParagraphs int, outputPath string) error {
	// Create and write to file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("could not create file for writing dummy text: %w", err)
	}
	// Will need to figure out how to catch any errors on a deferred file `Close` call.
	// Leave unhandled for now.
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i, text := range genText(numParagraphs) {
		paragraphText := text + "\n"
		if i > 0 {
			paragraphText = "\n" + paragraphText
		}
		_, err = writer.WriteString(paragraphText)
		if err != nil {
			return fmt.Errorf("could not write dummy text to file: %w", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func padLeft(text string, length int, padChar string) string {
	if len(text) >= length {
		return text
	}
	return strings.Repeat(padChar, length-len(text)) + text
}

func GenDummyFiles(numFiles, numParagraphs int, outputPath string) error {
	// Ensure that the directory exists
	err := os.MkdirAll(outputPath, 0755)
	if err != nil {
		return fmt.Errorf("couldn't create directories: %w", err)
	}

	// Write dummy files
	for i := range numFiles {
		currFileNum := strconv.Itoa(i + 1)
		fileName := dummyFileName + padLeft(currFileNum, dummyFileNameZerosPad, "0") + ".txt"
		err := genAndWrite(numParagraphs, filepath.Join(outputPath, fileName))
		if err != nil {
			return err
		}
	}
	return nil
}
