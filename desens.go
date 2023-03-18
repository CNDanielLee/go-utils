package utils

import "regexp"

// DesensitizeEmail 邮箱脱敏
func DesensitizeEmail(source string) string {
	reg, err := regexp.Compile("(\\w+)\\w{3}@(\\w+)")
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(source, "$1***@$2")
}

// DesensitizePhone 手机号脱敏
func DesensitizePhone(source string) string {
	reg, err := regexp.Compile("(\\d{3})\\d*(\\d{2})")
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(source, "$1******$2")
}
