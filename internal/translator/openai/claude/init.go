package claude

import (
	. "github.com/ikiwihome/evas-api/v6/internal/constant"
	"github.com/ikiwihome/evas-api/v6/internal/interfaces"
	"github.com/ikiwihome/evas-api/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		Claude,
		OpenAI,
		ConvertClaudeRequestToOpenAI,
		interfaces.TranslateResponse{
			Stream:     ConvertOpenAIResponseToClaude,
			NonStream:  ConvertOpenAIResponseToClaudeNonStream,
			TokenCount: ClaudeTokenCount,
		},
	)
}
