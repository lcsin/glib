package pkg

import (
	"testing"
)

func TestOnlyNumberReg(t *testing.T) {
	t.Log(OnlyNumberRegexp.MatchString(""))            // false
	t.Log(OnlyNumberRegexp.MatchString("123123123"))   // true
	t.Log(OnlyNumberRegexp.MatchString("sdf2131sf"))   // false
	t.Log(OnlyNumberRegexp.MatchString("123sdl123fs")) // false
	t.Log(OnlyNumberRegexp.MatchString("-123123123"))  // false
}

func TestContainChineseReg(t *testing.T) {
	t.Log(ContainChineseRegexp.MatchString("你好，你叫什么"))       // true
	t.Log(ContainChineseRegexp.MatchString("hello, world!")) // false
	t.Log(ContainChineseRegexp.MatchString("hello,你好!"))     // true
	t.Log(ContainChineseRegexp.MatchString("你好，world!"))     // true
}

func TestEmail(t *testing.T) {
	t.Log(EmailRegexp.MatchString("1847@qq.com")) // true
	t.Log(EmailRegexp.MatchString("1847xqq.com")) // false
	t.Log(EmailRegexp.MatchString("1847@qq.xxx")) // true
}

func TestPhone(t *testing.T) {
	t.Log(PhoneRegexp.MatchString("18758387692")) // true
	t.Log(PhoneRegexp.MatchString("11111111111")) // false
	t.Log(PhoneRegexp.MatchString("12345678901")) // false
}
