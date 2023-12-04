package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/chyroc/requests"
)

const output = "./docs"

func main() {
	fmt.Println("start")
	assert(
		twitter("OpenAI"),
		twitter("GoogleAI"),
		twitter("LangChainAI"),

		reddit("LangChain"),
	)
}

func twitter(username string) error {
	output := output + "/twitter/" + username + ".xml"
	url := fmt.Sprintf("https://rsshub.app/twitter/user/%s", username)
	if err := os.MkdirAll(filepath.Dir(output), os.ModePerm); err != nil {
		return fmt.Errorf("twitter mkdir failed: %s", err)
	}
	text := requests.Get(url).Text()
	if text.IsErr() {
		return fmt.Errorf("twitter get failed: %s", text.Err())
	}
	return ioutil.WriteFile(output, []byte(text.Value()), 0644)
}

func reddit(username string) error {
	output := output + "/reddit/" + username + ".xml"
	url := fmt.Sprintf("https://www.reddit.com/r/%s.rss", username)
	if err := os.MkdirAll(filepath.Dir(output), os.ModePerm); err != nil {
		return fmt.Errorf("reddit mkdir failed: %s", err)
	}
	text := requests.Get(url).Text()
	if text.IsErr() {
		return fmt.Errorf("reddit get failed: %s", text.Err())
	}
	return ioutil.WriteFile(output, []byte(text.Value()), 0644)
}

func assert(err ...error) {
	for _, v := range err {
		if v != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}
