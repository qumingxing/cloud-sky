package web

import (
	"net/http"
)

type HttpResponse struct {
	response http.ResponseWriter
}

func (response *HttpResponse) Write(responseStr string) (int, error) {

	return response.response.Write([]byte(responseStr))
}
func (response *HttpResponse) WriteHeader(headers map[string]string) {
	for key, value := range headers {
		response.response.Header().Set(key, value)
	}

}
