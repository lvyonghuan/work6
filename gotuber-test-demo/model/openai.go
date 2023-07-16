package model

// Request 对话使用的Request body
type Request struct {
	Model            string            `json:"model"`
	Messages         []RequestMessages `json:"messages"` //message依靠传入信息获取
	MaxTokens        int               `json:"max_tokens"`
	Temperature      float64           `json:"temperature"`
	TopP             float64           `json:"top_p"`
	Stop             string            `json:"stop"`
	PresencePenalty  float64           `json:"presence_penalty"`
	FrequencyPenalty float64           `json:"frequency_penalty"`
	User             string            `json:"user"`
}

// OpenAiRcv Response
type OpenAiRcv struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Message      RequestMessages `json:"message"`
		FinishReason string          `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens    int `json:"prompt_tokens"`
		CompletionTokes int `json:"completion_tokens"`
		TotalTokens     int `json:"total_tokens"`
	}
}

type RequestMessages struct {
	Role    string `json:"role"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
