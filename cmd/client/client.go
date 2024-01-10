package main

import (
	"flag"

	"github.com/bersennaidoo/netops/presentation/tui/httpclient"
)

func main() {
	var url string
	var file string

	flag.StringVar(&file, "filename", "", "json config file")
	flag.StringVar(&url, "url", "", "server address")
	flag.Parse()

	httpclient.NewPost(file, url)
}
