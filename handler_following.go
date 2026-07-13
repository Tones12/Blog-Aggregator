package main

import (
	"context"
	"fmt"

	"github.com/tones12/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	userFeedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Printf("%s is following the following feeds:\n", user.Name)
	
	for _, follow := range userFeedFollows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}
