package examples

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func ExampleGetConversations() {
	cli := getClient()

	res, err := cli.GetConversations()

	if err != nil {
		log.Fatalf("get conversations fail: %v\n", err)
	}

	for _, v := range res.Array() {
		conversationId := v.Get("uuid").String()
		conversationName := v.Get("name").String()
		log.Printf("conversation: %s, %s\n", conversationId, conversationName)
	}

	fmt.Println(1)
	// Output: 1
}

func ExampleNewConversation() {
	conversationId := uuid.New().String()
	name := "my conversation"

	cli := getClient()
	params := map[string]interface{}{
		"name": name,
		"uuid": conversationId,
	}

	res, err := cli.NewConversation(params)
	if err != nil {
		log.Fatalf("new conversation fail: %v\n", err)
	}

	conversationName := res.Get("name").String()
	log.Printf("conversation name: %s\n", conversationName)

	fmt.Println(1)
	// Output: 1
}
