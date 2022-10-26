package goslack

import (
	"go-log-slack/consts"
	"go-log-slack/methods"
	"time"
)

// addSingleBlock will add a single column block
func addSingleBlock(typ string, text *Text) Blocks {
	block := Blocks{
		Type: typ,
		Text: text,
	}
	return block
}

// addSectionBlock will add a multi-column block
func addSectionBlock(fields []*Fields) Blocks {
	block := Blocks{
		Type:   "section",
		Fields: fields,
	}
	return block
}

// addText will generate a text of type plain_text/mrkdwn
func addText(typ string, txt string, emoji *bool) *Text {
	text := Text{
		Type:  typ,
		Text:  txt,
		Emoji: emoji,
	}
	return &text
}

// addField will add a field of type mrkdwn
func addField(typ string, txt string) *Fields {
	field := Fields{
		Type: typ,
		Text: txt,
	}
	return &field
}

func getHeader(status int) string {
	var headerTitle string

	switch status {
	case Success:
		headerTitle = "Success"
	case Warning:
		headerTitle = "Warning"
	case Alert:
		headerTitle = "Alert"
	default:
		headerTitle = "Warning"
	}
	return headerTitle
}

// PrepareAttachmentBody will prepare whole Attachment body
func PrepareAttachmentBody(req ClientRequest) []Attachments {

	serviceName := req.ServiceName
	summary := req.Summary
	metadata := req.Metadata
	details := req.Details
	status := req.Status
	mentions := req.Mentions

	headerTitle := getHeader(status)

	if req.Header != "" {
		headerTitle = req.Header
	}

	if len(mentions) > 0 {
		mentionString := ""
		for _, m := range mentions {
			mentionString = mentionString + " " + m
		}

		summary = summary + mentionString
	}

	color := StatusMap[Warning]

	if v, ok := StatusMap[status]; ok {
		color = v
	}

	currentTime := time.Now()

	currentTimeStr := currentTime.Format("2006-01-02 15:04:05")

	emoji := true

	headerText := addText("plain_text", headerTitle, &emoji)

	headerBlock := addSingleBlock("header", headerText)

	serviceNameField := addField("mrkdwn", "*Service:*\n"+serviceName)

	serviceLogTimeField := addField("mrkdwn", "*Created At:*\n"+currentTimeStr)

	serviceInfoBlock := addSectionBlock([]*Fields{serviceNameField, serviceLogTimeField})

	summaryField := addField("mrkdwn", "*Summary:*\n"+summary)

	summaryBlock := addSectionBlock([]*Fields{summaryField})

	var detailsBlocks []Blocks
	detailsArr := methods.Chunks(details, consts.ChunkSize)

	for ind, detail := range detailsArr {
		if ind == 0 {
			detailsField := addField("mrkdwn", "*Details:*\n")
			detailsBlock := addSectionBlock([]*Fields{detailsField})
			detailsBlocks = append(detailsBlocks, detailsBlock)
		}
		detailsText := addText("mrkdwn", "```"+detail+"```", nil)
		detailsBlock := addSingleBlock("section", detailsText)
		detailsBlocks = append(detailsBlocks, detailsBlock)
	}

	metadataText := addText("mrkdwn", "*Metadata:*\n"+metadata, nil)
	metadataBlock := addSingleBlock("section", metadataText)

	blocks := []Blocks{headerBlock, serviceInfoBlock, summaryBlock, metadataBlock}

	for ind, block := range detailsBlocks {
		if ind <= consts.MaxBlocks {
			blocks = append(blocks, block)
		}
	}

	attachment := Attachments{
		Color:  color,
		Blocks: blocks,
	}

	return []Attachments{attachment}
}
