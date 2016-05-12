package web

import ()

type WebIntercepter interface {
	Intercepter(request *HttpRequest, response *HttpResponse) bool
}
