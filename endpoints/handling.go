package endpoints

import (
	"net/http"
	"strings"
)

func HandleEndPoints(r *http.Request) string {
	request := r.RequestURI[1:]
	if request == "areyoualive" {
		return "yeah"
	} else if strings.HasPrefix(request, "get/") {
		return handleGetrequest(request[4:])
	} else if strings.HasPrefix(request, "set/") {
		return handleSetrequest(request[4:])
	}
	return "404"
}

func handleGetrequest(url string) string {
	return url
}

func handleSetrequest(url string) string {
	return url
}
