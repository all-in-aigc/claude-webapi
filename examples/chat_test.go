package examples

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func ExampleChatWithQuestion() {
	text := "say hello to me 3 times"
	conversationId := "625e4059-1e68-41c7-8262-6442881f2cee"

	cli := getClient()
	params := map[string]interface{}{
		"attachments": []map[string]interface{}{},
		"completion": map[string]interface{}{
			"incremental": true,
			"model":       cli.GetModel(),
			"prompt":      "",
			"timezone":    "Asia/Shanghai",
		},
		"organization_uuid": cli.GetOrgid(),
		"conversation_uuid": conversationId,
		"text":              text,
	}

	res, err := cli.GetChatStream(params)
	if err != nil {
		log.Fatalf("chat with prompt failed: %v", err)
	}

	reply := ""
	for v := range res.Stream {
		log.Printf("chat stream: %s\n", v.String())
		if v.Get("stop_reason").String() == "stop_sequence" {
			break
		}
		reply += v.Get("completion").String()
	}

	log.Println("reply:", reply)

	fmt.Println(1)
	// Output: 1
}

func ExampleChatWithAttachment() {
	text := `summary room chat messages in file, what're they talk about?`
	conversationId := "625e4059-1e68-41c7-8262-6442881f2cee"
	attachment, _ := getAttachment()

	cli := getClient()
	params := map[string]interface{}{
		"attachments": attachment,
		"completion": map[string]interface{}{
			"incremental": true,
			"model":       cli.GetModel(),
			"prompt":      text,
			"timezone":    "Asia/Shanghai",
		},
		"organization_uuid": cli.GetOrgid(),
		"conversation_uuid": conversationId,
		"text":              text,
	}

	res, err := cli.GetChatStream(params)

	if err != nil {
		log.Printf("get chat stream failed: %v\n", err)
	}

	reply := ""
	for v := range res.Stream {
		log.Printf("chat stream: %s\n", v.String())
		if v.Get("stop_reason").String() == "stop_sequence" {
			break
		}
		reply += v.Get("completion").String()
	}

	log.Println("reply:", reply)

	fmt.Println(1)
	// Output: 1
}

type ChatMessage struct {
	CreatedAt int64
	UserName  string
	Content   string
}

func getAttachment() ([]map[string]interface{}, error) {
	attachments := []map[string]interface{}{}

	msgs := []*ChatMessage{
		{1690041600, "sam", "hello"},
		{1690041605, "gpt-4", "hello my friend"},
		{1690041610, "claude", "hello boy"},
		{1690041618, "sam", "let's talk about something about llm"},
		{1690041622, "gpt-4", "sure, I'am intertesd in it"},
		{1690041625, "claude", "sounds funny"},
		{1690041631, "sam", "let's go head"},
	}

	content := "room chat messages with format: [time] user: content\n\n"
	for _, msg := range msgs {
		if msg.Content == "" {
			continue
		}

		t := time.Unix(msg.CreatedAt, 0)
		content += fmt.Sprintf("%s %s: %q\n", t.Format("2006-01-02 15:04:05"), msg.UserName, msg.Content)
	}

	if content == "" {
		return nil, errors.New("no content")
	}

	fileName := "chat_messages.txt"
	attachments = append(attachments, map[string]interface{}{
		"extracted_content": content,
		"file_name":         fileName,
		"file_size":         len(content),
		"file_type":         "text/plain",
	})

	return attachments, nil
}
