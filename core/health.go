package core

import (
	"net/http"
)

func getStatus(url string) int {
	resp, err := http.Head(url)
	if err != nil {
		return 500
	}
	return resp.StatusCode
}

func healthCheck(urls []string) map[string]int {
	statusMap := make(map[string]int)
	for x := 0; x < len(urls); x++ {
		statusMap[urls[x]] = getStatus(urls[x])
	}
	return statusMap
}

func Check() map[string]int {
	urls := GetUrls()
	return healthCheck(urls)
}
