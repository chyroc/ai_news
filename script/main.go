package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chyroc/requests"
)

const output = "./docs"

func main() {
	fmt.Println("start")
	assert(
		rssHubTwitter("OpenAI"),
		rssHubTwitter("GoogleAI"),
		rssHubTwitter("LangChainAI"),

		rssHubHuggingfacePaper(),

		reddit("LangChain"),
	)
}

func rssHubTwitter(username string) error {
	output := output + "/twitter/" + username + ".xml"
	url := fmt.Sprintf("https://rsshub.app/twitter/user/%s", username)
	return pureWebPage("twitter", url, output)
}

func rssHubHuggingfacePaper() error {
	output := output + "/huggingface/daily-papers.xml"
	url := fmt.Sprintf("https://rsshub.app/huggingface/daily-papers")
	return pureWebPage("twitter", url, output)
}

func reddit(username string) error {
	output := output + "/reddit/" + username + ".xml"
	url := fmt.Sprintf("https://www.reddit.com/r/%s.rss", username)
	return pureWebPage("reddit", url, output)
}

func pureWebPage(title, url, output string) error {
	if err := os.MkdirAll(filepath.Dir(output), os.ModePerm); err != nil {
		return fmt.Errorf("%s mkdir failed: %s", title, err)
	}
	text := requests.Get(url).WithTimeout(time.Second * 10).Text()
	if text.IsErr() {
		return fmt.Errorf("%s get failed: %s", title, text.Err())
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
