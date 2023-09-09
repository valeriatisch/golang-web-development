package utils

import "log"

func CheckError(text string, err error) {
	if err != nil {
		log.Fatal(text, err)
	}
}
