package gemini

import (
	. "github.com/ikiwihome/evas-api/v6/internal/constant"
	"github.com/ikiwihome/evas-api/v6/internal/interfaces"
	"github.com/ikiwihome/evas-api/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		Gemini,
		OpenAI,
		ConvertGeminiRequestToOpenAI,
		interfaces.TranslateResponse{
			Stream:     ConvertOpenAIResponseToGemini,
			NonStream:  ConvertOpenAIResponseToGeminiNonStream,
			TokenCount: GeminiTokenCount,
		},
	)
}
