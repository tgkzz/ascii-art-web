package asciiart

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
)

func Validate() error {
	shadow_hash := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	standard_hash := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	thinkertoy_hash := "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52"

	shadow_text, err := os.ReadFile("internal/ascii-art/fs/shadow.txt")
	standard_text, err2 := os.ReadFile("internal/ascii-art/fs/standard.txt")
	thinkertoy_text, err3 := os.ReadFile("internal/ascii-art/fs/thinkertoy.txt")
	if err != nil || err2 != nil || err3 != nil {
		log.Print("asd")
		errtext := fmt.Sprintf("there is a proble with reading template files: \n1) %v\n2) %v \n3) %v\n", err, err2, err3)
		return errors.New(errtext)
	}

	hash_shdw := sha256.Sum256(shadow_text)
	hash_shdw_string := fmt.Sprintf("%x", hash_shdw)

	hash_stnd := sha256.Sum256(standard_text)
	hash_stdn_string := fmt.Sprintf("%x", hash_stnd)

	hash_thnk := sha256.Sum256(thinkertoy_text)
	hash_thnk_string := fmt.Sprintf("%x", hash_thnk)

	if hash_shdw_string != shadow_hash || hash_stdn_string != standard_hash || hash_thnk_string != thinkertoy_hash {
		log.Print("asd2")
		return errors.New("the template files were changed")
	}
	return nil
}
