package pkg

import "regexp"

var (
	// OnlyNumber 只包含数字的正则表达式
	OnlyNumber *regexp.Regexp

	// ContainChinese 包含中文字符的正则表达式
	ContainChinese *regexp.Regexp

	// Email 邮箱
	Email *regexp.Regexp

	// Phone 手机号（中国）
	Phone *regexp.Regexp
)

func init() {
	OnlyNumber = regexp.MustCompile("^\\d+$")

	ContainChinese = regexp.MustCompile("\\p{Han}")

	Email = regexp.MustCompile("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$")

	Phone = regexp.MustCompile("^1[3-9]\\d{9}$")
}
