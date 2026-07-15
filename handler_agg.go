package main

import (
	"fmt"
	"context"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	time_between_reqs := cmd.Args[0]
	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		fmt.Errorf("error in duration between requests: %s", err)
	}
	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("New feed request in progress")
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed fetched: %w", err)
	}

	feedRSS, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error collecting RSS information from feed: %w", err)
	}

	fmt.Printf("==========================================\n\nStarting feed...\n\n==========================================\n\n")
	fmt.Println(feedRSS.Channel.Title)
	fmt.Println(feedRSS.Channel.Description)

	for _, item := range feedRSS.Channel.Item {
		fmt.Println(item.Title)
		fmt.Println(item.Description)
	}

	return nil
}