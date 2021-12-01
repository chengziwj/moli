package validtor

import (
	"regexp"
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
