package pkg

import (
	"testing"
)

func TestOnlyNumberReg(t *testing.T) {
	t.Log(OnlyNumber.MatchString(""))            // false
	t.Log(OnlyNumber.MatchString("123123123"))   // true
	t.Log(OnlyNumber.MatchString("sdf2131sf"))   // false
	t.Log(OnlyNumber.MatchString("123sdl123fs")) // false
	t.Log(OnlyNumber.MatchString("-123123123"))  // false
}

func TestContainChineseReg(t *testing.T) {
	t.Log(ContainChinese.MatchString("你好，你叫什么"))       // true
	t.Log(ContainChinese.MatchString("hello, world!")) // false
	t.Log(ContainChinese.MatchString("hello,你好!"))     // true
	t.Log(ContainChinese.MatchString("你好，world!"))     // true
}

func TestEmail(t *testing.T) {
	t.Log(Email.MatchString("1847@qq.com")) // true
	t.Log(Email.MatchString("1847xqq.com")) // false
	t.Log(Email.MatchString("1847@qq.xxx")) // true
}

func TestPhone(t *testing.T) {
	t.Log(Phone.MatchString("18758387692")) // true
	t.Log(Phone.MatchString("11111111111")) // false
	t.Log(Phone.MatchString("12345678901")) // false
}
