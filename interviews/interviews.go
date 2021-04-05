package openblind

import (
	"errors"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/jacoelho/openblind"
	"golang.org/x/net/html"
)

const datetimeFormat = "2006-1-2"

var (
	ErrNoDateTime      = errors.New("no date time")
	interviewRe        = regexp.MustCompile(`^Interview(?P<ID>\d+)Container$`)
	matcherContainer   = openblind.WithDataTestRe(interviewRe)
	matcherTitle       = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+Title$`))
	matcherRating      = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+Rating$`))
	matcherApplication = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+ApplicationDetails$`))
	matcherProcess     = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+Process$`))
	matcherQuestions   = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+Questions$`))
)

type Interview struct {
	ID          string    `json:"id,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Title       string    `json:"title,omitempty"`
	Application []string  `json:"application,omitempty"`
	Process     []string  `json:"process,omitempty"`
	Questions   []string  `json:"questions,omitempty"`
}

// <time dateTime="2021-3-25">25 Mar 2021</time>
func ParseDateTime(node *html.Node) (time.Time, error) {
	var value string

	_, found := openblind.Find(node, func(n *html.Node) bool {
		v, ok := openblind.WithAttr(n, func(s string) bool {
			return s == "datetime"
		})
		value = v
		return ok
	})
	if !found {
		return time.Time{}, errors.New("failed to find datetime")
	}

	return time.Parse(datetimeFormat, value)
}

func ParseTitle(node *html.Node) (string, error) {
	titleNode, found := openblind.Find(node, matcherTitle)
	if !found {
		return "", errors.New("failed to find title")
	}

	return strings.Join(openblind.Text(titleNode), ","), nil
}

func ParseApplication(node *html.Node) ([]string, error) {
	applicationNode, found := openblind.Find(node, matcherApplication)
	if !found {
		return nil, errors.New("failed to find application")
	}

	// it always starts with word Application
	return openblind.RemoveStrings("Application")(openblind.Text(applicationNode)), nil
}

func ParseProcess(node *html.Node) ([]string, error) {
	processNode, found := openblind.Find(node, matcherProcess)
	if !found {
		return nil, errors.New("failed to find process")
	}

	return openblind.FlattenByNewLine(openblind.Text(processNode)), nil
}

func ParseQuestions(node *html.Node) ([]string, error) {
	questionsNode, found := openblind.Find(node, matcherQuestions)
	if !found {
		return nil, errors.New("failed to find process")
	}

	return openblind.RemoveStrings("Answer Question", "1 Answer")(openblind.FlattenByNewLine(openblind.Text(questionsNode))), nil
}

func parseInterview(node *html.Node) (Interview, error) {
	var result Interview

	datetime, err := ParseDateTime(node)
	if err != nil {
		return result, ErrNoDateTime
	}

	title, err := ParseTitle(node)
	if err != nil {
		return result, err
	}

	application, err := ParseApplication(node)
	if err != nil {
		return result, err
	}

	process, err := ParseProcess(node)
	if err != nil {
		return result, err
	}

	questions, err := ParseQuestions(node)
	if err != nil {
		return result, err
	}

	return Interview{
		Date:        datetime,
		Title:       title,
		Application: application,
		Process:     process,
		Questions:   questions,
	}, nil
}

func Parse(r io.Reader) ([]Interview, error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	list, ok := openblind.Find(root, openblind.WithDataTest("InterviewList"))
	if !ok {
		return nil, errors.New("failed to find interview list")
	}

	interviews := openblind.FindAll(list, matcherContainer)

	result := make([]Interview, 0, len(interviews))
	for _, interview := range interviews {
		res, err := parseInterview(interview)
		if err != nil {
			// featured interviews don't have a datetime, ignore
			if errors.Is(err, ErrNoDateTime) {
				continue
			}
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}
