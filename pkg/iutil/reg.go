package iutil

import "regexp"

var (
	// OnlyNumberRegexp 只包含数字的正则表达式
	OnlyNumberRegexp *regexp.Regexp

	// ContainChineseRegexp 包含中文字符的正则表达式
	ContainChineseRegexp *regexp.Regexp

	// EmailRegexp 邮箱
	EmailRegexp *regexp.Regexp

	// PhoneRegexp 手机号（中国）
	PhoneRegexp *regexp.Regexp
)

func init() {
	OnlyNumberRegexp = regexp.MustCompile("^\\d+$")

	ContainChineseRegexp = regexp.MustCompile("\\p{Han}")

	EmailRegexp = regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")

	PhoneRegexp = regexp.MustCompile("^1[3-9]\\d{9}$")
}
