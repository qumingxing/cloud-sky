package common

import (
	"net/http"
	"net/url"
)

func HttpPost(urlStr string, params map[string]string) string {
	var valus url.Values = make(url.Values)
	for key, value := range params {
		valus.Add(key, value)
	}
	response, err := http.PostForm(urlStr, valus)
	if err == nil {
		if response.StatusCode == http.StatusOK {
			return ReadString(response.Body)
		}
	}
	return ""
}
func HttpGet(urlStr string, params map[string]string) string {
	var builder StringBuilder
	var i int = 0
	var sum int = len(params)
	for key, value := range params {
		if (i + 1) == sum {
			builder.Concat(key + "=" + value)
		} else {
			builder.Concat(key + "=" + value + "&")
		}
		i++
	}
	response, err := http.Get(urlStr + "?" + builder.ToString())
	if err == nil {
		if response.StatusCode == http.StatusOK {
			return ReadString(response.Body)
		}
	}
	return ""
}
