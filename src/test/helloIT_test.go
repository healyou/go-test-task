package test

import (
	"testing"

	"../request"
	"../utils"
)

func TestSendUrls(t *testing.T) {
	var urls6Line = "https://golang.org/\nhttps://golang.org/\nhttps://golang.org/\nhttps://golang.org/\nhttps://golang.org/\nhttps://golang.org/"
	var urls1Line = "https://golang.org/"
	testCases := []struct {
		expectedSum int
		gorCnt      int
		urlsLine    string
	}{
		{120, 3, urls6Line},
		{120, 6, urls6Line},
		{120, 10, urls6Line},
		{20, 1, urls1Line},
		{20, 2, urls1Line},
	}
	for _, tc := range testCases {
		var sum int = request.SendUrlsAndGetSumm(tc.urlsLine, utils.SERCH_STR, tc.gorCnt)
		validateValue(tc.expectedSum, sum, t)
	}
}

func validateValue(expected int, actual int, t *testing.T) {
	if actual != expected {
		t.Error("Expected ", expected, " got ", actual)
	}
}
