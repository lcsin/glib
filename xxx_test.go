package main

import (
	"regexp"
	"testing"
)

func init() {
	onlyQuestionMark = regexp.MustCompile(`^\[[?]]$`)

	onlyDoubleAsterisk = regexp.MustCompile(`^\[\*\*]$`)

	onlyChinese = regexp.MustCompile(`^\[\p{Han}+\|\p{Han}+]$`)
}

func TestReg(t *testing.T) {
	v := "[联系方式]"
	t.Log(onlyChinese.MatchString(v))
	t.Log(onlyQuestionMark.MatchString(v))
	t.Log(onlyDoubleAsterisk.MatchString(v))
}
