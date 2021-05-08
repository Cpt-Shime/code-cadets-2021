package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type applicatants struct {
	Name   string
	Age    int
	Passed bool
	Skills []string
}

func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

const mocky = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func main() {
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(mocky)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards m API"),
		)
	}
	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of "),
		)
	}

	var applications []applicatants
	err = json.Unmarshal(bodyContent, &applications)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}
	log.Printf("Response: %v", applications)

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}
	defer f.Close()

	for _, application := range applications {
		if !application.Passed {
			continue
		}
		for _, skill := range application.Skills {
			if skill == "Go" || skill == "Java" {
				fmt.Println("Pisem u file")
				f.WriteString(fmt.Sprint(application))
			}
		}
	}
}
