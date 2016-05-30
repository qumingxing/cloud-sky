package web

import (
	"bufio"
	"common"
	"fmt"
	"io"
	"net/http"
	"os"
	"logs"
)

type HttpRequest struct {
	session      Session
	cookie       Cookie
	httpRequest  *http.Request
	httpResponse *HttpResponse
}

var defaultUploadPath = os.TempDir()

func (httpRequest *HttpRequest) GetSession() (newSession *Session) {
	var cookie Cookie
	resCookie := cookie.GetCookie("SESSIONID", httpRequest)
	if resCookie == nil {
		newSession = newSessionTask(httpRequest)
	} else {
		newSession = sessionMap[resCookie.Value]
		if newSession == nil {
			newSession = newSessionTask(httpRequest)
		}
	}
	return
}
func (httpRequest *HttpRequest) GetCookieByName(name string) string {
	cookie, err := httpRequest.httpRequest.Cookie(name)
	if err == nil {
		return cookie.Value
	}
	logs.Error(name, "获取cookie失败", err)
	return ""
}
func (httpRequest *HttpRequest) GetRequest() *http.Request {
	return httpRequest.httpRequest
}
func (httpRequest *HttpRequest) SetUploadPath(path string) {
	defaultUploadPath = path
}
func (httpRequest *HttpRequest) FormValue(name string) string {
	return httpRequest.httpRequest.FormValue(name)
}
func (httpRequest *HttpRequest) RequestURI() string {
	return httpRequest.httpRequest.RequestURI
}
func (httpRequest *HttpRequest) FormFile(fileName string) {
	file, fileHeader, err := httpRequest.GetRequest().FormFile(fileName)
	if err != nil {
		return
	}
	var bytes [1024]byte
	reader := bufio.NewReader(file)
	os.MkdirAll(defaultUploadPath, os.ModeDir)
	storePath := fmt.Sprint(defaultUploadPath, "/", fileHeader.Filename)
	writeFile, createErr := os.Create(storePath)
	defer func() {
		file.Close()
		writeFile.Close()
	}()
	if createErr != nil {
		return
	}
	writer := bufio.NewWriter(writeFile)
	for {
		offset, err := reader.Read(bytes[0:])
		if err == io.EOF {
			break
		} else {
			writer.Write(bytes[0:offset])
		}
	}
}
func newSessionTask(httpRequest *HttpRequest) (newSession *Session) {
	var cookie Cookie
	guid := common.GetGuid()
	var newCookie *Cookie = &Cookie{Name: "SESSIONID", Value: guid}
	cookie.AddCookie(newCookie, httpRequest.httpResponse)
	newSession = &Session{guid}
	sessionMap[guid] = newSession
	return
}
func (httpRequest *HttpRequest) Redirect(url string) {
	http.Redirect(httpRequest.httpResponse.response, httpRequest.httpRequest, url, http.StatusMovedPermanently)
}
