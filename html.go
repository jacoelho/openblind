package openblind

import (
	"regexp"

	"golang.org/x/net/html"
)

type Matcher func(*html.Node) bool

func WithAttrFn(n *html.Node, fn func(string) bool) (string, bool) {
	for _, a := range n.Attr {
		if fn(a.Key) {
			return a.Val, true
		}
	}
	return "", false
}

func WithAttr(n *html.Node, s string) (string, bool) {
	return WithAttrFn(n, func(v string) bool {
		return s == v
	})
}

func WithID(id string) Matcher {
	return func(n *html.Node) bool {
		v, found := WithAttr(n, "id")
		return found && v == id
	}
}

func WithClass(class string) Matcher {
	return func(n *html.Node) bool {
		v, found := WithAttr(n, "class")
		return found && v == class
	}
}

func WithDataTest(value string) Matcher {
	return func(n *html.Node) bool {
		v, found := WithAttr(n, "data-test")
		return found && v == value
	}
}

func WithDataTestRe(re *regexp.Regexp) Matcher {
	return func(n *html.Node) bool {
		v, found := WithAttr(n, "data-test")
		return found && re.MatchString(v)
	}
}

func WithIDRe(re *regexp.Regexp) Matcher {
	return func(n *html.Node) bool {
		v, found := WithAttr(n, "id")
		return found && re.MatchString(v)
	}
}

// Find returns first node that matches Matcher
func Find(node *html.Node, m Matcher) (*html.Node, bool) {
	if m(node) {
		return node, true
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if n, ok := Find(c, m); ok {
			return n, true
		}
	}
	return nil, false
}

func FindAll(node *html.Node, m Matcher) []*html.Node {
	var result []*html.Node

	if m(node) {
		result = append(result, node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		found := FindAll(c, m)

		if len(found) > 0 {
			result = append(result, found...)
		}
	}

	return result
}

func ExtractText(node *html.Node) []string {
	nodes := FindAll(node, func(n *html.Node) bool {
		return n.Type == html.TextNode
	})

	result := make([]string, len(nodes))
	for i, n := range nodes {
		result[i] = n.Data
	}

	return result
}

func AttrValue(node *html.Node, attr string) (string, bool) {
	var value string

	_, found := Find(node, func(n *html.Node) bool {
		v, ok := WithAttr(n, attr)
		if ok {
			value = v
		}
		return ok
	})

	return value, found
}
