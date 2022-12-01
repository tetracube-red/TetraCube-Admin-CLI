package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateTokenForNewUser(host string) map[string]string {
	fullApi := fmt.Sprintf("http://%s:8080/authentication-token", host)

	resp, err := http.Post(fullApi, "application/json", nil)
	if err != nil {
		fmt.Printf("Error in authentication token creation %s\n", err.Error())
	}

	body, _ := io.ReadAll(resp.Body)
	var responseMap map[string]string
	_ = json.Unmarshal(body, &responseMap)
	fmt.Printf("Authentication token %s and valid until %s", responseMap["token"], responseMap["validUntil"])

	return responseMap
}
