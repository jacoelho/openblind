package openblind

import (
	"strings"
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
