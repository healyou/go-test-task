package main

import (
	"errors"
	"os"

	"./request"
	"./utils"
)

func main() {
	argsWithoutProg := os.Args[1:]
	var argsCount int = len(argsWithoutProg)
	if argsCount <= 0 || argsCount > 1 {
		errors.New("Неправильное количество параметров")

	} else {
		var urlsLine string = argsWithoutProg[0]
		request.SendUrlsAndGetSumm(urlsLine, utils.SERCH_STR, utils.MAX_GOROUTINE_COUNT)
	}
}
