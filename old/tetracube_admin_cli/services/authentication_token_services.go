package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateTokenForNewUser(host string, houseName string) map[string]string {
	bodyMap := map[string]string{"houseName": houseName}
	jsonValue, _ := json.Marshal(bodyMap)

	fullApi := fmt.Sprintf("http://%s:8080/authentication-token", host)
	resp, err := http.Post(fullApi, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("Error in authentication token creation %s\n", err.Error())
		return nil
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Cannot create authentication token, received status code %d\n", resp.StatusCode)
		return nil
	}
	body, _ := io.ReadAll(resp.Body)
	var responseMap map[string]string
	_ = json.Unmarshal(body, &responseMap)
	fmt.Printf("Authentication token %s and valid until %s", responseMap["token"], responseMap["validUntil"])

	return responseMap
}
