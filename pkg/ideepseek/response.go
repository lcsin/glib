package ideepseek

type Choices struct {
	Index        int         `json:"index"`
	Message      Messages    `json:"message"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens          int                 `json:"prompt_tokens"`
	CompletionTokens      int                 `json:"completion_tokens"`
	TotalTokens           int                 `json:"total_tokens"`
	PromptTokensDetails   PromptTokensDetails `json:"prompt_tokens_details"`
	PromptCacheHitTokens  int                 `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens int                 `json:"prompt_cache_miss_tokens"`
}

type PromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type ResponseBody struct {
	Id                string    `json:"id"`
	Object            string    `json:"object"`
	Created           int       `json:"created"`
	Model             string    `json:"model"`
	Choices           []Choices `json:"choices"`
	Usage             Usage     `json:"usage"`
	SystemFingerprint string    `json:"system_fingerprint"`
}
