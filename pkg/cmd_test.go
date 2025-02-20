package pkg

import (
	"context"
	"testing"
)

func TestClone(t *testing.T) {
	url := "https://github.com/lcsin/leetcode.git"
	path := "./leetcode"

	if err := GitClone(context.Background(), url, "", path); err != nil {
		panic(err)
	}
}

func TestPull(t *testing.T) {
	path := "C:\\Users\\Administrator\\.kratos\\repo\\github.com\\go-kratos\\kratos-layout@main"

	if err := GitPull(context.Background(), path); err != nil {
		panic(err)
	}
}
