package common

import (
	//"fmt"
	"strings"
)

type StringBuilder string

//var aaa common.StringBuilder
//aaa.Concat("456").Concat("789")
func (builder *StringBuilder) Concat(target string) *StringBuilder {
	*builder = *builder + StringBuilder(target)
	return builder
}
func (builder *StringBuilder) ToString() string {
	return string(*builder)
}
func Equals(s, t string) bool {
	return strings.EqualFold(s, t)
}
func Trim(str string) string {
	return strings.TrimSpace(str)
}
func IsNotEmpty(str string) bool {
	return len(Trim(str)) > 0
}
func IsEmpty(str string) bool {
	return len(Trim(str)) == 0
}
func IsBlank(obj interface{}) bool {
	return obj == nil
}
func IsNotBlank(obj interface{}) bool {
	return obj != nil
}
func LastIndexOf(str string, part string) int {
	return strings.LastIndex(str, part)
}
func IndexOf(str, part string) int {
	return strings.Index(str, part)
}
func Replace(str, oldStr, newStr string) string {
	return strings.Replace(str, oldStr, newStr, -1)
}
func Suffix(str, part string) bool {
	return strings.HasSuffix(str, part)
}
func Prefix(str, part string) bool {
	return strings.HasPrefix(str, part)
}
