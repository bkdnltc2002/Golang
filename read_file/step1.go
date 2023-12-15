package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFileSize(url string) (int64, error) {

	// sends an HTTP HEAD request
	// returns two values: a response (resp) and an error (err).
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	var size int64
	size = resp.ContentLength

	return size, nil
}

func createFile(url string, filename string) {
	// Send an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer response.Body.Close()

	name := filename + ".txt"
	// Create a new file to save the downloaded content
	outputFile, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Copy the content from the HTTP response to the output file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error copying content to the output file:", err)
	}
	modified := " " + filename + " modified"
	_, err = outputFile.WriteString(modified)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	fmt.Printf("Created files successfully")
}

func main() {
	var url string
	fmt.Print("Enter the URL of the file: ")
	fmt.Scanln(&url)

	size, err := getFileSize(url)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Printf("The file size of '%s' is %d bytes.\n", url, size)

	// List of filenames you want to create
	filenames := []string{"Thanh", "Thang"}

	for _, filename := range filenames {
		createFile(url, filename)
	}
}
