package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func tailFile(filename string, startPos int64) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Seek to the specified start position
	_, err = file.Seek(startPos, 0)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Process the new line
		println(line)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Return the current offset position
	pos, err := file.Seek(0, 1)
	if err != nil {
		return 0, err
	}

	return pos, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the filename as a command-line argument")
	}

	filename := os.Args[1]
	var startPos int64 = 0

	for {
		newPos, err := tailFile(filename, startPos)
		if err != nil {
			log.Fatal(err)
		}

		// Check if the file was truncated
		fileInfo, err := os.Stat(filename)
		if err != nil {
			log.Fatal(err)
		}
		if fileInfo.Size() < newPos {
			startPos = 0
		} else {
			startPos = newPos
		}

		time.Sleep(1 * time.Second) // Adjust the sleep duration as per your needs
	}
}
