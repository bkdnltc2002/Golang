package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func createFilev2(content []byte, filename string) {
	name := filename + ".txt"
	// Create a new file to save the downloaded content
	outputFile, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write the content to the output file
	_, err = outputFile.Write(content)
	if err != nil {
		fmt.Println("Error writing content to the output file:", err)
		return
	}

	modified := " " + filename + " modified"
	_, err = outputFile.WriteString(modified)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Created file %s successfully\n", name)
}

func main() {
	var url string
	fmt.Print("Enter the URL of the file: ")
	fmt.Scanln(&url)

	// Send an HTTP GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body into a variable called content
	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// List of filenames you want to create
	filenames := []string{"Thanh", "Thang"}

	for _, filename := range filenames {
		createFilev2(content, filename)
	}
}
