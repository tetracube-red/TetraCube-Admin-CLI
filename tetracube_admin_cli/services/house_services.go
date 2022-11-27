package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateHouse(host string, houseName string) {
	fullApi := fmt.Sprintf("http://%s/houses", host)
	bodyMap := map[string]string{"name": houseName}
	jsonValue, _ := json.Marshal(bodyMap)

	resp, err := http.Post(fullApi, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("Error in house creation %s\n", err.Error())
	}
	if resp.StatusCode == 409 {
		fmt.Printf("House %s already exists \n", houseName)
		return
	}

	body, _ := io.ReadAll(resp.Body)
	var responseMap map[string]string
	_ = json.Unmarshal(body, &responseMap)
	fmt.Printf("House %s created with id %s", responseMap["name"], responseMap["id"])
}
