// ServiceInterface
package web

/*
业务接口，不同业务必须实现该接口
*/
import (
	"net/http"
)

type ServiceInterface interface {
	Handle(response http.ResponseWriter, request *http.Request) IWebView
}
