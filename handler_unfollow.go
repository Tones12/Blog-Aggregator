package main

import (
	"context"
	"fmt"

	"github.com/tones12/blog-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	url := cmd.Args[0]

	feedID, err := s.db.GetFeedID(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed ID: %w", err)
	}

	userIDfeedID := database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feedID,
	}

	err = s.db.UnfollowFeed(context.Background(), userIDfeedID)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	fmt.Printf("%s is unfollowed", url)

	return nil
}
