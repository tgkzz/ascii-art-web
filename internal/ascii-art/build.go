package asciiart

import "log"

func Build(text, fs string) (string, error) {
	if err := AsciiChar(text); err != nil {
		log.Print("here1")
		return "", err
	}
	if err := Validate(); err != nil {
		log.Print("here2")
		return "", err
	}
	fsChar, err := LoadMap(fs)
	if err != nil {
		log.Print("here3")
		return "", err
	}

	split_text := Split(text)
	return Print(split_text, fsChar), nil
}
