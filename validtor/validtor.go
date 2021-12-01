package validtor

import (
	"fmt"
	"regexp"
	"unicode"
)

var (
	reMail     = regexp.MustCompile(PatternMail)
	reUserName = regexp.MustCompile(PatternUserName)
	reIpv4     = regexp.MustCompile(PatternIpv4)
	reIpv6     = regexp.MustCompile(PatternIpv6)
	reDomain   = regexp.MustCompile(PatternDomain)
)

func Mail(s string) bool {
	return reMail.MatchString(s)
}

//UserName 帐号是否合法(字母开头，允许5-16字节，允许字母数字下划线)
func UserName(s string) bool {
	return reUserName.MatchString(s)
}

func Ipv4(s string) bool {
	return reIpv4.MatchString(s)
}

func Ipv6(s string) bool {
	return reIpv6.MatchString(s)
}

func Domain(s string) bool {
	return reDomain.MatchString(s)
}

//Password 帐号是否合法(字母开头，允许5-16字节，允许字母数字下划线)
func Password(str string, minLen, maxLen int) error {
	var (
		isUpper   = false
		isLower   = false
		isNumber  = false
		isSpecial = false
	)

	if len(str) < minLen || len(str) > maxLen {
		return fmt.Errorf("The password must contain uppercase and lowercase letters, numbers or punctuation, and must be %d-%d digits long. ", minLen, maxLen)
	}

	for _, s := range str {
		switch {
		case unicode.IsUpper(s):
			isUpper = true
		case unicode.IsLower(s):
			isLower = true
		case unicode.IsNumber(s):
			isNumber = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			isSpecial = true
		default:
		}
	}

	if (isUpper && isLower) && (isNumber || isSpecial) {
		return nil
	}
	return fmt.Errorf("The password must contain uppercase and lowercase letters, numbers or punctuation, and must be %d-%d digits long. ", minLen, maxLen)
}
