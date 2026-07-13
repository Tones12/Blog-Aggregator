package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
	"context"

	"github.com/tones12/blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	
	url := cmd.Args[0]
	
	feedID, err := s.db.GetFeedID(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed ID: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		UserID:		user.ID,
		FeedID:		feedID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}
	
	fmt.Println("Feed successfully followed!")
	
	fmt.Printf("Feed: %v\nFollowed by: %v\n", feedFollow.FeedName, user.Name)

	return nil
}