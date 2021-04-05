package openblind

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func RemoveStrings(remove ...string) func(s []string) []string {
	removeMap := make(map[string]struct{})

	for _, c := range remove {
		removeMap[c] = struct{}{}
	}

	return func(s []string) []string {
		var result []string

		for _, el := range s {
			trimmed := strings.TrimSpace(el)
			_, remove := removeMap[trimmed]

			if trimmed == "" || remove {
				continue
			}

			result = append(result, trimmed)
		}

		return result
	}
}

func Flatten(fn func(string) []string) func([]string) []string {
	return func(s []string) []string {
		var result []string

		for _, subs := range s {
			splitted := fn(subs)
			for _, v := range splitted {
				if v != "" {
					result = append(result, v)
				}
			}
		}
		return result
	}
}

func FlattenByNewLine(v []string) []string {
	return Flatten(func(s string) []string {
		return strings.Split(s, "\n")
	})(v)
}

func WithDataTest(value string) Matcher {
	return func(n *html.Node) bool {
		v, found := WithAttr(n, func(s string) bool { return s == "data-test" })
		return found && v == value
	}
}

func WithDataTestRe(re *regexp.Regexp) Matcher {
	return func(n *html.Node) bool {
		v, found := WithAttr(n, func(s string) bool { return s == "data-test" })
		return found && re.MatchString(v)
	}
}

func WithIDRe(re *regexp.Regexp) Matcher {
	return func(n *html.Node) bool {
		v, found := WithAttr(n, func(s string) bool { return s == "id" })
		return found && re.MatchString(v)
	}
}
