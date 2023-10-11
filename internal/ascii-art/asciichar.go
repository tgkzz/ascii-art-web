package asciiart

import "errors"

func AsciiChar(text string) error {
	if text == "" {
		return errors.New("empty response")
	}
	for _, char := range text {
		if char == rune(13) {
			char = rune(10)
		}
		if !(char >= rune(32) && char <= rune(126) || char == rune(10)) {
			return errors.New("an invalid character was received in the response")
		}
	}
	return nil
}
