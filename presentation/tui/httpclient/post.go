package httpclient

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewPost(file, url string) {

	config, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	bufferBody := bytes.NewBuffer(config)

	resp, err := http.Post(url, "application/json", bufferBody)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)

	bodyAnswer := bufio.NewScanner(resp.Body)

	for bodyAnswer.Scan() {
		fmt.Println(bodyAnswer.Text())
	}
}
