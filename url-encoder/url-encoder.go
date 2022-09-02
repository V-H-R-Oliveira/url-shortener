package urlencoder

import (
	"log"
	goUrl "net/url"
	"strings"

	"github.com/jxskiss/base62"
)

func GetEncodedUrlPath(url string) string {
	parsedUrl, err := goUrl.ParseRequestURI(url)

	if err != nil {
		log.Printf("Failed to parse url %s due error %s\n", url, err.Error())
		return ""
	}

	before, _, _ := strings.Cut(strings.Replace(parsedUrl.EscapedPath(), "/", "", 1), "/")
	return before
}

func EncodeUrlPath(url string) string {
	return base62.EncodeToString([]byte(url))
}
