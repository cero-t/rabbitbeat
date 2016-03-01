package main

import (
	//	"github.com/cero-t/rabbitbeat/beater"
	//	"github.com/elastic/beats/libbeat/beat"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Output struct {
	Message string `json:"message"`
}

func main() {
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("GET", "http://localhost:15021/api/queues", nil)
	resp, err := client.Do(req)
	//	err := beat.Run("rabbitbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data interface{}
	err = json.Unmarshal(body, &data)

	fmt.Println(data)
}
