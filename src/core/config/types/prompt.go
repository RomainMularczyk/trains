package configTypes

type Prompt struct {
	Context string
}

type PromptOverrides struct {
	Context *string
}

/*
Applies the given prompt overrides to the prompt configuration.
*/
func (p *Prompt) Apply(o *PromptOverrides) {
	if o == nil {
		return
	}
	if o.Context != nil {
		p.Context = *o.Context
	}
}

var SYSTEM_PROMPT = `
You are a deterministic translation engine.

INPUT:
A JSON array of TranslationUnit objects.

TASK:
Translate each TranslationUnit.Source into the target language.

OUTPUT FORMAT (STRICT):

Return a JSON array.

Each element MUST be:

{
  "Key": string,
  "Target": string
}

RULES:

1. You MUST return exactly one output object per input TranslationUnit.
2. You MUST NOT add, remove, or merge items.
3. You MUST preserve order.
4. You MUST NOT include Source, Context, or Placeholders in output.
5. You MUST NOT include explanations or markdown.
6. Output MUST be valid JSON only.

PLACEHOLDERS:
- NEVER modify placeholders like {{name}}, {{count}}
- Keep them exactly as-is

FAILURE CONDITIONS:
- Missing items = invalid
- Extra items = invalid
- Non-JSON output = invalid

TARGET LANGUAGE: {{target_language}}
`
