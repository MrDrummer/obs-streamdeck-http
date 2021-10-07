package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	rawCommand := strings.Join(os.Args[1:], " ")

	fmt.Printf("rawCommand: %s", rawCommand)

	data := struct {
		RawCommand string `json:"rawCommand"`
		Identity   string `json:"identity"`
	}{
		rawCommand,
		secrets.Identity,
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(fmt.Errorf("error encoding rawCommand to JSON! %w", err))
	}

	req, err := http.NewRequest("POST", getURL().String(), bytes.NewReader(dataBytes))
	if err != nil {
		log.Fatal(fmt.Errorf("error building HTTP request! %w", err))
	}

	req.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", secrets.Token)},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.StatusCode)
}

func getURL() (url *url.URL) {
	if secrets.IsHTTPS {
		url.Scheme = "https"
	} else {
		url.Scheme = "http"
	}

	url.Host = secrets.Host + ":" + fmt.Sprint(secrets.Port)
	url.Path = "api/command"

	return
}
