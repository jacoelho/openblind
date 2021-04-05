package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jacoelho/openblind/interviews"
	"github.com/jacoelho/openblind/reviews"
)

const (
	exitCodeOK    = 0
	exitCodeError = 1
)

type config struct {
	targetURL string
	timeout   time.Duration
	section   string
	userAgent string
}

const (
	sectionInterviews = "interviews"
	sectionReviews    = "reviews"
)

var version string = "development"

func main() {
	var (
		c           config
		showVersion bool
	)

	flag.StringVar(&c.targetURL, "url", "", "url to parse")
	flag.DurationVar(&c.timeout, "timeout", 5*time.Second, "timeout duration")
	flag.StringVar(&c.section, "section", "interviews", "type of section, one of: interviews, reviews")
	flag.StringVar(&c.userAgent, "user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36", "user agent to use")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Println(version)
		os.Exit(exitCodeOK)
	}

	if c.targetURL == "" {
		flag.Usage()
		os.Exit(exitCodeError)
	}

	if c.section != sectionInterviews && c.section != sectionReviews {
		flag.Usage()
		os.Exit(exitCodeError)
	}

	if err := run(c); err != nil {
		log.Println(err)
		os.Exit(exitCodeError)
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

	req.Header.Set("User-Agent", cfg.userAgent)

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
