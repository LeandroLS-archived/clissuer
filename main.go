package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Issue struct {
	Title string `json:"title"`
}

func main() {

	var userF string
	var personalTokenF string
	var title string
	var action string

	flag.StringVar(&userF, "user", "", "Your Github User")
	flag.StringVar(&personalTokenF, "password", "", "Your Github Personal Token")
	flag.StringVar(&title, "title", "", "Issue Title")
	flag.StringVar(&action, "action", "", "What you wanna do ? search, create, delete, update")

	flag.Parse()
	if action == "create" {
		if userF == "" {
			log.Fatalf("-user flag should be informed")
		}
		if personalTokenF == "" {
			log.Fatalf("-password flag should be informed")
		}

		if title == "" {
			log.Fatalf("-title flag should be informed")
		}

		if action == "" {
			log.Fatalf("-action flag should be informed")
		}
		issue := Issue{Title: title}
		fmt.Println(issue.Title)
		jsonIssue, _ := json.Marshal(issue)
		payload := strings.NewReader(string(jsonIssue))
		url := "https://api.github.com/repos/LeandroLS/clissuer/issues"

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/vnd.github.v3+json")
		req.SetBasicAuth(userF, personalTokenF)
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))

	}

}
