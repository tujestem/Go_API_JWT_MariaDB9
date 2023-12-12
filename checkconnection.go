package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// LoginResponse which is representing API answer, with JWT included.
type LoginResponse struct {
	Token string `json:"token"`
}

// QueryResponse which represents answer from SQL QUery.
type QueryResponse []map[string]interface{}

// Log in and download the JWT token.
func getToken() (string, error) {
	loginURL := "http://localhost:8080/login"
	credentials := []byte(`{"username":"jankowalsky","password":"987656789aA!"}`)

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(credentials))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return "", err
	}

	return loginResp.Token, nil
}

// Make QUERY with JWT token included.
func makeQuery(token string) error {
	queryURL := "http://localhost:8080/query"
	query := `{"token":"` + token + `","query":"SELECT * FROM tab1"}`

	resp, err := http.Post(queryURL, "application/json", bytes.NewBufferString(query))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var queryResp QueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&queryResp); err != nil {
		return err
	}

	for _, row := range queryResp {
		for col, val := range row {
			if s, ok := val.(string); ok {
				decodedVal, err := base64.StdEncoding.DecodeString(s)
				if err == nil {
					fmt.Printf("%s: %s\n", col, string(decodedVal))
				} else {
					fmt.Printf("%s: %s\n", col, s)
				}
			}
		}
		fmt.Println("-----")
	}

	return nil
}

func main() {
	token, err := getToken()
	if err != nil {
		fmt.Println("Error during logging: ", err)
		return
	}

	if err := makeQuery(token); err != nil {
		fmt.Println("Error during QUERY execution: ", err)
	}
}
