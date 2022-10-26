package goslack

type SlackRequestBody struct {
	Attachments []Attachments `json:"attachments"`
}
type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji *bool   `json:"emoji,omitempty"`
}
type Fields struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}
type Blocks struct {
	Type   string   `json:"type,omitempty"`
	Text   *Text     `json:"text,omitempty"`
	Fields []*Fields `json:"fields,omitempty"`
}
type Attachments struct {
	Color  string   `json:"color,omitempty"`
	Blocks []Blocks `json:"blocks,omitempty"`
}
