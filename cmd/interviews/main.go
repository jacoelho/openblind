package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jacoelho/openblind"
)

const errorExitCode = 1

func main() {
	var targetURL string

	flag.StringVar(&targetURL, "url", "", "url with interviews")
	flag.Parse()

	if targetURL == "" {
		flag.Usage()
		os.Exit(errorExitCode)
	}

	if err := run(targetURL); err != nil {
		log.Println(err)
		os.Exit(errorExitCode)
	}

}

func run(targetURL string) error {
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u, err := url.Parse(targetURL)
	if err != nil {
		return err
	}

	u.Query().Set("sort.sortType", "RD")
	u.Query().Set("sort.ascending", "false")

	req, err := http.NewRequestWithContext(timeout, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	res, err := openblind.ParseInterviews(resp.Body)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	enc.Encode(&res)

	return nil
}
