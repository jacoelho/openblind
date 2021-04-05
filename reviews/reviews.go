package reviews

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jacoelho/openblind"
	"golang.org/x/net/html"
)

// Sun Mar 28 2021 06:27:08 GMT+0100
const datetimeFormat = "Mon Jan 02 2006 15:04:05 MST-0700"

var (
	reviewRe               = regexp.MustCompile(`^empReview_(?P<ID>\d+)$`)
	matcherReviewContainer = openblind.WithIDRe(reviewRe)

	ErrParseID     = errors.New("failed to parse id")
	ErrParseDate   = errors.New("failed to parse date")
	ErrParseRating = errors.New("failed to parse rating")
	ErrParseTitle  = errors.New("failed to parse title")
	ErrParsePros   = errors.New("failed to parse pros")
	ErrParseCons   = errors.New("failed to parse cons")
	ErrParseAdvice = errors.New("failed to parse advice")
)

type Review struct {
	ID     string    `json:"id,omitempty"`
	Date   time.Time `json:"date,omitempty"`
	Rating float64   `json:"rating,omitempty"`
	Pros   []string  `json:"pros,omitempty"`
	Cons   []string  `json:"cons,omitempty"`
	Advice []string  `json:"advice,omitempty"`
}

func ParseID(node *html.Node) (string, error) {
	var value string

	_, found := openblind.Find(node, func(n *html.Node) bool {
		v, ok := openblind.WithAttr(n, func(s string) bool { return s == "id" })

		if ok && reviewRe.MatchString(v) {
			idx := reviewRe.SubexpIndex("ID")
			value = reviewRe.FindStringSubmatch(v)[idx]
		}

		return ok
	})
	if !found {
		return "", ErrParseID
	}

	return value, nil
}

func ParseDatetime(node *html.Node) (time.Time, error) {
	var value string

	rating, found := openblind.Find(node, func(n *html.Node) bool {
		v, ok := openblind.WithAttr(n, func(s string) bool { return s == "class" })
		return ok && v == "date subtle small"
	})
	if !found {
		return time.Time{}, ErrParseRating
	}

	_, found = openblind.Find(rating, func(n *html.Node) bool {
		v, ok := openblind.WithAttr(n, func(s string) bool { return s == "datetime" })
		value = v
		return ok
	})
	if !found {
		return time.Time{}, ErrParseRating
	}

	// Sun Mar 28 2021 06:27:08 GMT+0100 (British Summer Time)
	split := strings.Split(value, " (")
	if len(split) != 2 {
		return time.Time{}, ErrParseRating
	}

	parseTime, err := time.Parse(datetimeFormat, split[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("%s: %w", err.Error(), ErrParseDate)
	}

	return parseTime, nil
}

func ParseRating(node *html.Node) (float64, error) {
	var value string

	rating, found := openblind.Find(node, func(n *html.Node) bool {
		v, ok := openblind.WithAttr(n, func(s string) bool { return s == "class" })
		return ok && v == "rating"
	})
	if !found {
		return 0, ErrParseRating
	}

	_, found = openblind.Find(rating, func(n *html.Node) bool {
		v, ok := openblind.WithAttr(n, func(s string) bool { return s == "title" })
		value = v
		return ok
	})
	if !found {
		return 0, ErrParseRating
	}

	parsedRating, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", err.Error(), ErrParseRating)
	}

	return parsedRating, nil
}

func ParseReviewTitle(node *html.Node) ([]string, error) {
	titleNode, found := openblind.Find(node, func(n *html.Node) bool {
		_, found := openblind.WithAttr(n, func(s string) bool { return s == "mainText mb-0" })
		return found
	})
	if !found {
		return nil, ErrParseTitle
	}

	return openblind.Text(titleNode), nil
}

func ParsePros(node *html.Node) ([]string, error) {
	pros, found := openblind.Find(node, openblind.WithDataTest("pros"))
	if !found {
		return nil, ErrParsePros
	}

	return openblind.Text(pros), nil
}

func ParseCons(node *html.Node) ([]string, error) {
	cons, found := openblind.Find(node, openblind.WithDataTest("cons"))
	if !found {
		return nil, ErrParseCons
	}

	return openblind.Text(cons), nil
}

func ParseAdvice(node *html.Node) ([]string, error) {
	advice, found := openblind.Find(node, openblind.WithDataTest("advice-management"))
	if !found {
		return nil, ErrParseAdvice
	}

	return openblind.Text(advice), nil
}

func ParseReview(node *html.Node) (Review, error) {
	var result Review

	id, err := ParseID(node)
	if err != nil {
		return result, err
	}

	reviewTime, err := ParseDatetime(node)
	if err != nil {
		return result, err
	}

	rating, err := ParseRating(node)
	if err != nil {
		return result, err
	}

	pros, err := ParsePros(node)
	if err != nil {
		return result, err
	}

	cons, err := ParseCons(node)
	if err != nil {
		return result, err
	}

	// not all reviews have advice
	advice, _ := ParseAdvice(node)

	return Review{
		ID:     id,
		Date:   reviewTime.UTC(),
		Rating: rating,
		Pros:   openblind.FlattenByNewLine(pros),
		Cons:   openblind.FlattenByNewLine(cons),
		Advice: openblind.FlattenByNewLine(advice),
	}, nil
}

func Parse(r io.Reader) ([]Review, error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	list, ok := openblind.Find(root, openblind.WithID("ReviewsFeed"))
	if !ok {
		return nil, errors.New("failed to find reviews")
	}

	reviews := openblind.FindAll(list, matcherReviewContainer)

	result := make([]Review, 0, len(reviews))
	for _, review := range reviews {
		res, err := ParseReview(review)
		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}
