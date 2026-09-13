package mee6

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	// "log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func GetGuildInfo(id int, page int) (Response, error) {
	var jsonr Response
	endpoint := fmt.Sprintf("https://mee6.xyz/api/plugins/levels/leaderboard/%d?page=%d", id, page)
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Println("Error while GET:", err)
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response body:", body)
		return Response{}, err
	}

	err = json.Unmarshal(body, &jsonr)
	if err != nil {
		log.Println("Error while JSON decoding:", err)
		// Print the raw response body for debugging
		fmt.Println("Raw Response Body:", string(body))
		return Response{}, err
	}
	return jsonr, nil
}

func MockGetInfo(id int, page int) (Response, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory:", err)
		return Response{}, err
	}

	// Construct the full path to the JSON file
	dir := filepath.Join(wd, "mock", fmt.Sprintf("%d.json", page))

	file, err := os.Open(dir)
	if err != nil {
		return Response{}, err
	}
	defer file.Close()

	// Read the content of the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return Response{}, err
	}

	// Unmarshal the JSON content into a struct
	var person Response
	err = json.Unmarshal(content, &person)
	if err != nil {
		return Response{}, err
	}
	return person, nil

}

func CrawlGuild(id int) ([]Response, error) {
	var responses []Response
	for i := 0; i >= 0; i++ {
		data, err := GetGuildInfo(id, i)
		if err != nil {
			return responses, err
		}
		if len(data.Players) == 0 {
			break
		}
		responses = append(responses, data)
		time.Sleep(2 * time.Second)
	}
	return responses, nil
}
