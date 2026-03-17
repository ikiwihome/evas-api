package geminiCLI

import (
	. "github.com/ikiwihome/evas-api/v6/internal/constant"
	"github.com/ikiwihome/evas-api/v6/internal/interfaces"
	"github.com/ikiwihome/evas-api/v6/internal/translator/translator"
)

func init() {
	translator.Register(
		GeminiCLI,
		Codex,
		ConvertGeminiCLIRequestToCodex,
		interfaces.TranslateResponse{
			Stream:     ConvertCodexResponseToGeminiCLI,
			NonStream:  ConvertCodexResponseToGeminiCLINonStream,
			TokenCount: GeminiCLITokenCount,
		},
	)
}
