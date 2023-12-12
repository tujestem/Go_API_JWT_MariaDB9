package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// checking if user exists in DB.
func checkUserExists(nazwisko string) (bool, error) {
	apiUrl := fmt.Sprintf("http://localhost:8080/checkuser/%s", nazwisko)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result map[string]bool
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result["exists"], nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("surname to deleting in DB: ")
		nazwisko, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error until user surename reading: ", err)
			continue
		}
		nazwisko = strings.TrimSpace(nazwisko)

		// Checking that user exists or no, in DB.
		exists, err := checkUserExists(nazwisko)
		if err != nil {
			fmt.Println("Error until user checking: ", err)
			continue
		}
		if !exists {
			fmt.Println("User doesn't exists. ")
			continue
		}

		fmt.Print("Do You really want to remove user from DB? (y/n): ")
		confirm, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error until answer reading: ", err)
			continue
		}
		confirm = strings.TrimSpace(confirm)

		if strings.ToLower(confirm) == "y" {
			// DELETE opration request.
			apiUrl := fmt.Sprintf("http://localhost:8080/deleteuser/%s", nazwisko)
			req, err := http.NewRequest(http.MethodDelete, apiUrl, nil)
			if err != nil {
				fmt.Println("Error until demand creation: ", err)
				continue
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error until sending query: ", err)
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println("User cannot be deleted from DB! Error code: ", resp.StatusCode)
				continue
			}

			fmt.Println("User has been sucessfully deleted from DB.")
			break
		} else {
			fmt.Println("Operation aborted.")
			break
		}
	}
}
