// WebView
package web

type IWebView interface {
	GetHtml() (content []byte, err error)
	GetResultPath(rootPath string) string
	GetResultObject() interface{}
	GetTemplateName() string
}
type WebView struct {
	ReturnURI    string
	ResultObj    interface{}
	TemplateName string
}

func (view *WebView) GetHtml() (content []byte, err error) {
	htmlParse := &ParserHtml{view.ReturnURI}
	content, err = htmlParse.GetHtml()
	return
}
func (view *WebView) GetResultPath(rootPath string) string {
	return rootPath + view.ReturnURI
}
func (view *WebView) GetResultObject() interface{} {
	return view.ResultObj
}
func (view *WebView) GetTemplateName() string {
	return view.TemplateName
}
