package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do You want to add test user into database? (y/n): ")
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	response = strings.TrimSpace(response)

	if strings.ToLower(response) == "y" {
		apiUrl := "http://localhost:8080/addtestuser1"

		// Create POST request
		resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer([]byte{}))
		if err != nil {
			fmt.Println("Error until request sending: ", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("User hasn't been added. Error code: ", resp.StatusCode)
			return
		}

		fmt.Println("New user has been sucessfully added.")
	} else {
		fmt.Println("Operation aborted.")
	}
}
