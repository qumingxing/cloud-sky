package web

import (
	"net/http"
)

type Cookie struct {
	Name     string
	Value    string
	Path     string
	Domain   string
	HttpOnly bool
	MaxAge   int
}

func (cookie Cookie) GetCookies(httpRequest *HttpRequest) []*Cookie {
	request := httpRequest.GetRequest()
	var cookiesArr []*Cookie = make([]*Cookie, 0)
	cookies := request.Cookies()
	for _, cookie := range cookies {
		coo := &Cookie{cookie.Name, cookie.Value, cookie.Path, cookie.Domain, cookie.HttpOnly, cookie.MaxAge}
		cookiesArr = append(cookiesArr, coo)
	}
	return cookiesArr
}
func (cookie Cookie) GetCookie(name string, request *HttpRequest) *Cookie {
	cookies := cookie.GetCookies(request)
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}
func (cookie Cookie) AddCookie(newCookie *Cookie, resp *HttpResponse) {
	httpCookie := &http.Cookie{Name: newCookie.Name, Value: newCookie.Value, Path: newCookie.Path, Domain: newCookie.Domain, HttpOnly: newCookie.HttpOnly, MaxAge: newCookie.MaxAge}
	http.SetCookie(resp.response, httpCookie)
}
func (cookie Cookie) DeleteCookie(newCookie *Cookie, resp HttpResponse) {
	httpCookie := &http.Cookie{Name: newCookie.Name, Value: newCookie.Value, Path: newCookie.Path, Domain: newCookie.Domain, HttpOnly: newCookie.HttpOnly, MaxAge: -1}
	http.SetCookie(resp.response, httpCookie)
}
