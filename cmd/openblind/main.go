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

	"github.com/jacoelho/openblind/interviews"
	"github.com/jacoelho/openblind/reviews"
)

const errorExitCode = 1

type config struct {
	targetURL string
	timeout   time.Duration
	section   string
}

const (
	sectionInterviews = "interviews"
	sectionReviews    = "reviews"
)

func main() {
	var c config
	flag.StringVar(&c.targetURL, "url", "", "url to parse")
	flag.DurationVar(&c.timeout, "timeout", 5*time.Second, "timeout duration")
	flag.StringVar(&c.section, "section", "interviews", "type of section, one of: interviews, reviews")
	flag.Parse()

	if c.targetURL == "" {
		flag.Usage()
		os.Exit(errorExitCode)
	}

	if c.section != sectionInterviews && c.section != sectionReviews {
		flag.Usage()
		os.Exit(errorExitCode)
	}

	if err := run(c); err != nil {
		log.Println(err)
		os.Exit(errorExitCode)
	}

}

func run(cfg config) error {
	timeout, cancel := context.WithTimeout(context.Background(), cfg.timeout)
	defer cancel()

	u, err := url.Parse(cfg.targetURL)
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

	var res interface{}
	if cfg.section == sectionInterviews {
		res, err = interviews.Parse(resp.Body)
	} else {
		res, err = reviews.Parse(resp.Body)
	}
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "\t")
	enc.Encode(&res)

	return nil
}
