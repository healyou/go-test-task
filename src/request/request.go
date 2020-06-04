package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../utils"
)

// выполнить поиск строки search по результатам всех запросов urlsLine с ограниченным кол-вом горутин
func SendUrlsAndGetSumm(urlsLine string, search string, maxGoroutineCnt int) int {
	var urls []string = strings.Fields(urlsLine)

	context := NewContext(maxGoroutineCnt)
	var urlsStack utils.Stack = utils.NewFromArray(urls)

	for !urlsStack.IsEmpty() {
		if !context.isLimitGoroutine() {
			var url string = urlsStack.Pop()
			context.runFuncAsGoroutine(searchForUrl, search, url)
		} else {
			context.releaseGoroutines()
		}
	}
	context.waitAndRelease()

	fmt.Println("Total: ", context.getSum())
	return context.getSum()
}

// отправляет get запрос по url и находит кол-во вхождений search
func searchForUrl(url string, search string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response string = string(body)
	var count int = strings.Count(response, search)
	fmt.Println("Count for ", url, ": ", count)
	return count
}
