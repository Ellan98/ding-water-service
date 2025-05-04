package ports

type PostChatCompletionRequest struct {
	Prompt          string `form:"prompt" json:"prompt" binding:"required"`
	SearchEnabled   bool   `form:"searchEnabled" json:"searchEnabled"`
	ThinkingEnabled bool   `form:"thinkingEnabled" json:"thinkingEnabled"`
}
