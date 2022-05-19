package news

import (
	hackernews "L30/hackernews"
	"context"
	"encoding/json"
	"log"
)

type Server struct {
	// Story hackernews.StoryService
}

func (s *Server) SayHello(ctx context.Context, message *MessageClient) (*Message, error) {
	ss := hackernews.NewStoryService(message.Title)
	res := ss.GetStories(10)

	result, _ := json.MarshalIndent(res, "", "")

	log.Printf("Received url bod from client:%s", message.Title)
	return &Message{Title: result}, nil
}
