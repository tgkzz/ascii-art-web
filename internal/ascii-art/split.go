package asciiart

import "strings"

func Split(text string) []string {
	var split_text []string
	text = strings.ReplaceAll(text, string(rune(13)), "\n")
	parts := strings.Split(text, "\n")

	for i, part := range parts {
		if part != "" {
			split_text = append(split_text, part)
		}
		if i < len(parts)-1 {
			split_text = append(split_text, "\n")
		}
	}
	return split_text
}
