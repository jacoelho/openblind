package interviews

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
	interviewRe        = regexp.MustCompile(`^Interview(?P<ID>\d+)Container$`)
	matcherContainer   = openblind.WithDataTestRe(interviewRe)
	matcherTitle       = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+Title$`))
	matcherApplication = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+ApplicationDetails$`))
	matcherProcess     = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+Process$`))
	matcherQuestions   = openblind.WithDataTestRe(regexp.MustCompile(`^Interview\d+Questions$`))

	ErrNoDateTime       = errors.New("no date time")
	ErrParseID          = errors.New("failed to parse id")
	ErrParseDate        = errors.New("failed to parse date")
	ErrParseTitle       = errors.New("failed to parse title")
	ErrParseApplication = errors.New("failed to parse application")
	ErrParseProcess     = errors.New("failed to parse process")
	ErrParseQuestions   = errors.New("failed to parse questions")
)

type Interview struct {
	ID          string    `json:"id,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Title       string    `json:"title,omitempty"`
	Application []string  `json:"application,omitempty"`
	Process     []string  `json:"process,omitempty"`
	Questions   []string  `json:"questions,omitempty"`
}

func parseID(node *html.Node) (string, error) {
	var value string

	_, found := openblind.Find(node, func(n *html.Node) bool {
		v, ok := openblind.WithAttr(n, "data-test")
		if ok && interviewRe.MatchString(v) {
			idx := interviewRe.SubexpIndex("ID")
			value = interviewRe.FindStringSubmatch(v)[idx]
		}

		return ok
	})
	if !found {
		return "", ErrParseID
	}

	return value, nil
}

// <time dateTime="2021-3-25">25 Mar 2021</time>
func parseDateTime(node *html.Node) (time.Time, error) {
	value, found := openblind.AttrValue(node, "datetime")
	if !found {
		return time.Time{}, ErrParseDate
	}

	return time.Parse(datetimeFormat, value)
}

func parseTitle(node *html.Node) ([]string, error) {
	titleNode, found := openblind.Find(node, matcherTitle)
	if !found {
		return nil, ErrParseTitle
	}

	return openblind.ExtractText(titleNode), nil
}

func parseApplication(node *html.Node) ([]string, error) {
	applicationNode, found := openblind.Find(node, matcherApplication)
	if !found {
		return nil, ErrParseApplication
	}

	return openblind.ExtractText(applicationNode), nil
}

func parseProcess(node *html.Node) ([]string, error) {
	processNode, found := openblind.Find(node, matcherProcess)
	if !found {
		return nil, ErrParseProcess
	}

	return openblind.ExtractText(processNode), nil
}

func parseQuestions(node *html.Node) ([]string, error) {
	questionsNode, found := openblind.Find(node, matcherQuestions)
	if !found {
		return nil, ErrParseQuestions
	}

	return openblind.ExtractText(questionsNode), nil
}

func parseInterview(node *html.Node) (Interview, error) {
	var result Interview

	id, err := parseID(node)
	if err != nil {
		return result, err
	}

	datetime, err := parseDateTime(node)
	if err != nil {
		return result, ErrNoDateTime
	}

	title, err := parseTitle(node)
	if err != nil {
		return result, err
	}

	application, err := parseApplication(node)
	if err != nil {
		return result, err
	}

	process, err := parseProcess(node)
	if err != nil {
		return result, err
	}

	questions, err := parseQuestions(node)
	if err != nil {
		return result, err
	}

	return Interview{
		ID:          id,
		Date:        datetime,
		Title:       strings.Join(title, ","),
		Application: openblind.RemoveStrings("Application")(openblind.FlattenByNewLine(application)),
		Process:     openblind.FlattenByNewLine(process),
		Questions:   openblind.RemoveStrings("Answer Question", "1 Answer")(openblind.FlattenByNewLine(questions)),
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
