package ideepseek

import (
	"os"
	"testing"
)

func TestChat(t *testing.T) {
	url := "https://api.deepseek.com"
	apikey := os.Getenv("DeepSeek_APIKEY")
	client := NewClient(url, apikey)

	messages := []Messages{
		{
			Role:    RoleSystem,
			Content: "现在你是一个资深的翻译人员，我输入中文，请润色后合理的翻译成英文并输出",
		}, {
			Role:    RoleUser,
			Content: "还有一天就放春节了，我很期待。",
		},
	}
	response, err := client.Chat(messages...)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	t.Log(response.Choices[0].Message)
}
