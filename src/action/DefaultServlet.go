// DefaultServlet
package action

import (
	"fmt"
	"web"
)

type DefaultServlet struct {
}

func (defaultServlet *DefaultServlet) DefaultMethod(request *web.HttpRequest, response *web.HttpResponse) web.IWebView {
	fmt.Println("----------------ok----------------")
	fmt.Println(request.GetSession().GetAttribute("hello"))
	return &web.WebView{"success.html", nil, "success template"}
}
