package main

import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}
		
	fmt.Println("Listing Feeds:")

	for _, feed := range feeds {
		name, err := s.db.GetUsername(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting username: %w", err)
		}
		
		fmt.Printf("\nName: %s\nURL: %s\nUser: %s\n", feed.Name, feed.Url, name)
	}

	return nil
}