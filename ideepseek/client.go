package ideepseek

const (
	ModelChat     = "deepseek-chat"
	ModelReasoner = "deepseek-reasoner"

	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
)

type Client struct {
	Url    string // 接口地址：https://api.deepseek.com
	Apikey string // apikey

	Model  string // 调用模型
	Stream bool   // 是否流式
	// 代码生成/数学解题	0.0
	// 数据抽取/分析	1.0
	// 通用对话	1.3
	// 翻译	1.3
	// 创意类写作/诗歌创作	1.5
	Temperature float64 // 温度，默认1.0
	MaxTokens   int64   // 最大输出长度，默认为4k
}

func NewClient(url, apikey string) *Client {
	return &Client{
		Url:    url,
		Apikey: apikey,
		Model:  ModelChat,
	}
}

func (c *Client) WithModel(model string) *Client {
	c.Model = model
	return c
}

func (c *Client) WithStream(stream bool) *Client {
	c.Stream = stream
	return c
}

func (c *Client) WithTemperature(temperature float64) *Client {
	c.Temperature = temperature
	return c
}

func (c *Client) WithMaxTokens(maxTokens int64) *Client {
	c.MaxTokens = maxTokens
	return c
}
