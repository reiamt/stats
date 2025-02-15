package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiUrl := "https://www-genesis.destatis.de/genesisWS/rest/2020/find/find"	
	// apiUrl := "https://www-genesis.destatis.de/genesisWS/rest/2020/helloworld/logincheck"
	// create post request
	payload := map[string]string{
		"username": "GAST",//"tim.maier1@web.de",
		"password": "7iYiCVJF&@qN",
		// "name": "KREISE",
		// "selection": "12*",
		// "area": "all",
		// "searchcriterion": "code",
		// "sortcriterion": "code",
		// "pagelength": "20",
		"term": "Abfall",
		"category": "all",
		"pagelength": "20",
		"language": "de",
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(respBody))

}