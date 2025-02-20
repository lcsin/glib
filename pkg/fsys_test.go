package pkg

import (
	"testing"
)

func TestExists(t *testing.T) {
	t.Log(PathExists("fsys.go"))
}

func TestCopyFile(t *testing.T) {
	if err := CopyFile("fsys.go", "fsys1.go"); err != nil {
		panic(err)
	}
	t.Log("copy file ok")
}

func TestCopyDir(t *testing.T) {
	if err := CopyDir("lua", "./lua1", []string{"lua.go", "lua_test.go"}); err != nil {
		panic(err)
	}
	t.Log("copy dir ok")
}

func TestFileMD5(t *testing.T) {
	t.Log(FileMD5("fsys.go"))
}

func TestZipCompress(t *testing.T) {
	files := []string{"lua/lua.go", "fsys.go"}
	if err := ZipCompress(files, "./zip/pkg.zip", true); err != nil {
		panic(err)
	}
	t.Log("zip ok")
}
