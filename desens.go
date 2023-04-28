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

type PhoneDigit uint8

const (
	PhoneDigitFour PhoneDigit = iota
	PhoneDigitSix
)

// DesensitizePhone 手机号脱敏
func DesensitizePhone(source string, digit PhoneDigit) string {
	switch digit {
	case PhoneDigitFour:
		reg, err := regexp.Compile("(\\d{3})\\d*(\\d{4})")
		if err != nil {
			return ""
		}
		return reg.ReplaceAllString(source, "$1****$2")
	case PhoneDigitSix:
		reg, err := regexp.Compile("(\\d{3})\\d*(\\d{2})")
		if err != nil {
			return ""
		}
		return reg.ReplaceAllString(source, "$1******$2")
	}
	return ""
}
