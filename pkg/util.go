package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5 md5
func MD5(text string) string {
	md := md5.New()
	md.Write([]byte(text))
	return strings.ToUpper(hex.EncodeToString(md.Sum(nil)))
}

// LogicalPaginate 逻辑分页，返回分页后的数组
func LogicalPaginate[T any](list []T, pageNo, pageSize int64) []T {
	if pageNo <= 0 || pageSize <= 0 {
		return []T{}
	}

	total := int64(len(list))
	totalPage := total / pageSize
	if totalPage%pageSize != 0 {
		totalPage++
	}

	start := (pageNo - 1) * pageSize
	if start >= total {
		return []T{}
	}

	end := start + pageSize
	if end > total {
		return list[start:]
	}
	return list[start:end]
}

// DelSliceElement 删除切片中某个下标的元素
func DelSliceElement[T any](list []T, idx int) []T {
	return append(list[:idx], list[idx+1:]...)
}

// RemoveSliceDuplicates slice去重
func RemoveSliceDuplicates[E comparable](slice []E) []E {
	set := make(map[E]struct{})
	var res []E
	for _, v := range slice {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// ExtractTextBetweenWildcards 提取被通配符包含的内容包含通配符本身
func ExtractTextBetweenWildcards(text, start, end string) []string {
	var result []string
	var stack []int

	for i := 0; i < len(text); i++ {
		if i+len(start) <= len(text) && text[i:i+len(start)] == start {
			stack = append(stack, i)
			i += len(start) - 1
		} else if i+len(end) <= len(text) && text[i:i+len(end)] == end && len(stack) > 0 {
			startIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			extracted := text[startIndex : i+len(end)]
			result = append(result, extracted)
			i += len(end) - 1
		}
	}
	return result
}

// BalancedWildcards 判断文本中的通配符对是否配对
func BalancedWildcards(text string, dict map[string]string) bool {
	if len(text) == 0 {
		return false
	}

	stack := make([]string, 0, len(text))
	left := make(map[string]bool, len(dict)/2)
	right := make(map[string]bool, len(dict)/2)
	reverse := make(map[string]string, len(dict))
	for k, v := range dict {
		reverse[v] = k
	}

	for k, v := range dict {
		left[k] = true
		right[v] = true
	}

	for _, v := range text {
		str := string(v)
		switch true {
		case left[str]: // 左边直接入栈
			stack = append(stack, str)
		case right[str]: // 右边判断是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != reverse[str] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
