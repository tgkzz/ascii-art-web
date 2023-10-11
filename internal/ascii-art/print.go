package asciiart

func Print(split_text []string, fsChar map[rune][]string) string {
	result := ""
	for _, word := range split_text {
		if word == "\n" {
			result += "\n"
			continue
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					result += fsChar[char][i]
				}
				if i != 7 {
					result += "\n"
				}
			}
		}
	}
	return result
}
