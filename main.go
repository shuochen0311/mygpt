package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chatgp/gpt3"
)

const (
	apiKey = "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"
)

func askGPT(cli *gpt3.Client, input string) string {
	uri := "/v1/chat/completions"
	params := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{"role": "user", "content": input},
		},
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v", err)
	}

	message := res.Get("choices.0.message.content").String()

	return message
}

func main() {

	apiKey := "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"

	// new gpt-3 client
	cli, _ := gpt3.NewClient(&gpt3.Options{
		ApiKey:  apiKey,
		Timeout: 30 * time.Second,
		Debug:   false,
	})

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 2 {
			result := line
			// result := askGPT(cli, line)
			fmt.Println(result)
		}
	}

}
