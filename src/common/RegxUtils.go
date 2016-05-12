package common

import (
	"regexp"
)

const DIGIT = "^[0-9]\\d*$"
const INTEGER = "^[1-9]\\d*$"
const EMAIL = "\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*"
const DATE = "(19|20)\\d\\d-(0[1-9]|1[0-2])-([0][1-9]|[1,2][0-9]|3[0-1])"
const AGE = "120|((1[0-1]|\\d)?\\d)"
const MONEY = "^(([1-9]\\d*)|0)(\\.\\d{1,2})?$"
const MOBILE = "^((\\d{3}))?1[3,5,8][0-9]\\d{8}"
const IP = "([1-9]|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])(\\.(\\d|[1-9]\\d|1\\d{2}|2[0-4]\\d|25[0-5])){3}"
const FULL_DATE = "((^((1[8-9]\\d{2})|([2-9]\\d{3}))(-)(10|12|0?[13578])(-)(3[01]|[12][0-9]|0?[1-9])$)" +
	"|(^((1[8-9]\\d{2})|([2-9]\\d{3}))(-)(11|0?[469])(-)(30|[12][0-9]|0?[1-9])$)|(^((1[8-9]" +
	"\\d{2})|([2-9]\\d{3}))(-)(0?2)(-)(2[0-8]|1[0-9]|0?[1-9])$)|(^([2468][048]00)(-)(0?2)(-)" +
	"(29)$)|(^([3579][26]00)(-)(0?2)(-)(29)$)|(^([1][89][0][48])(-)(0?2)(-)(29)$)|(^([2-9][0-9]" +
	"[0][48])(-)(0?2)(-)(29)$)|(^([1][89][2468][048])(-)(0?2)(-)(29)$)|(^([2-9][0-9][2468][048])" +
	"(-)(0?2)(-)(29)$)|(^([1][89][13579][26])(-)(0?2)(-)(29)$)|(^([2-9][0-9][13579][26])(-)(0?2)" +
	"(-)(29)$))"
const TELEPHONE = "0\\d{2,3}-\\d{7,8}|0\\d{4}-\\d{7,8}"
const CERT = "[\\d]{6}(19)?[\\d]{2}((0[1-9])|(10|11|12))([012][\\d]|(30|31))[\\d]{3}[xX\\d]*"
const ALL_DIGIT = "^\\d*(.\\d+)?$"

func IsDigit(str string) bool {
	reg, _ := regexp.Compile(DIGIT)
	res := reg.MatchString(str)
	return res
}
func IsInteger(str string) bool {
	reg, _ := regexp.Compile(INTEGER)
	res := reg.MatchString(str)
	return res
}
func IsEmail(str string) bool {
	reg, _ := regexp.Compile(EMAIL)
	res := reg.MatchString(str)
	return res
}
func IsDate(str string) bool {
	reg, _ := regexp.Compile(DATE)
	res := reg.MatchString(str)
	return res
}
func IsAge(str string) bool {
	reg, _ := regexp.Compile(AGE)
	res := reg.MatchString(str)
	return res
}
func IsMoney(str string) bool {
	reg, _ := regexp.Compile(MONEY)
	res := reg.MatchString(str)
	return res
}
func IsMobile(str string) bool {
	reg, _ := regexp.Compile(MOBILE)
	res := reg.MatchString(str)
	return res
}
func IsIp(str string) bool {
	reg, _ := regexp.Compile(IP)
	res := reg.MatchString(str)
	return res
}
func IsFullDate(str string) bool {
	reg, _ := regexp.Compile(FULL_DATE)
	res := reg.MatchString(str)
	return res
}
func IsTelephone(str string) bool {
	reg, _ := regexp.Compile(TELEPHONE)
	res := reg.MatchString(str)
	return res
}
func IsCert(str string) bool {
	reg, _ := regexp.Compile(CERT)
	res := reg.MatchString(str)
	return res
}
func IsAllDigit(str string) bool {
	reg, _ := regexp.Compile(ALL_DIGIT)
	res := reg.MatchString(str)
	return res
}
func IsMatch(rexg, str string) bool {
	reg, _ := regexp.Compile(rexg)
	res := reg.MatchString(str)
	return res
}
func IsURIMatch(requestURI, path string) bool {
	regExp := ""
	// /a/b/*.xxx
	// /a/b/*
	// /a/*/c/*.xxx
	// /a/*/c/*
	// /a/b/c.html or /a/b/c
	if LastIndexOf(path, "/*.") != -1 && LastIndexOf(path, "/*/") == -1 {
		regExp = Replace(path, "*.", ".*\\.")
		regExp += "$"
	} else if LastIndexOf(path, "/*") != -1 && LastIndexOf(path, "/*/") == -1 {
		regExp = Replace(path, "/*", "/.*")
	} else if LastIndexOf(path, "/*/") != -1 && LastIndexOf(path, "/*.") != -1 {
		str1 := Replace(path, "/*/", "/.*/")
		regExp = Replace(str1, "*.", ".*\\.")
		regExp += "$"
	} else if LastIndexOf(path, "/*/") != -1 && LastIndexOf(path, "/*") != -1 {
		str1 := Replace(path, "/*/", "/.*/")
		regExp = Replace(str1, "/*", "/.*")
	} else {
		regExp = path
	}
	if IsNotEmpty(regExp) && IsMatch(regExp, requestURI) {
		return true
	}
	return false
}
