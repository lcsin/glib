package pkg

import (
	"testing"
)

func TestUtil(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	slice = DelSliceElement(slice, 3)
	t.Log(slice)
}

func TestRemoveDuplicates(t *testing.T) {
	slice := []int{1, 2, 3, 4, 4, 5, 6}
	slice = RemoveSliceDuplicates(slice)
	t.Log(slice)
}

func TestExtractTextBetweenWildcards(t *testing.T) {
	text := "这是一段包含[[姓名]]的[文本]"
	t.Log(ExtractTextBetweenWildcards(text, "[", "]")) // [[姓名] [[姓名]] [文本]]
}

func TestBalancedWildcards(t *testing.T) {
	dict := make(map[string]string, 0)
	dict["{"] = "}"
	dict["["] = "]"
	dict["<"] = ">"

	text := "a{nihao}b,c[nihao]d,<nihao>.aaa"
	t.Log(BalancedWildcards(text, dict)) // true

	text = "a{nihao},b{nihao},c[<sdfsf>]"
	t.Log(BalancedWildcards(text, dict)) // true

	text = "{<>[]<}}"
	t.Log(BalancedWildcards(text, dict)) // false
}
